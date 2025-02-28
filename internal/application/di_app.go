// Package app ...
package application

import (
	"fmt"
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driven/gorm"
	adminRegionsDriven "github.com/andriykusevol/aktemplategorm/internal/adapters/driven/gorm/administrativeregion"
	authDriven "github.com/andriykusevol/aktemplategorm/internal/adapters/driven/gorm/auth"
	feeSetDriven "github.com/andriykusevol/aktemplategorm/internal/adapters/driven/gorm/feeset"
	api_logger "github.com/andriykusevol/aktemplategorm/internal/adapters/driving/logger"
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driving/restapi/handler"
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driving/restapi/router"
	"github.com/andriykusevol/aktemplategorm/internal/application/app"
	"github.com/andriykusevol/aktemplategorm/pkg/logger"

	"context"
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driving/restapi/middleware"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// Run the application
func Run(
	appMode string,
	appPort string,
	databaseDsn string,
	mysqlMaxOpenConns string,
	mysqlMaxIDLEConns string,
	appComponent string,
	appVersion string,
	appEnv string,
) {

	fmt.Println("1=====")
	fmt.Println(databaseDsn)
	fmt.Println("2=====")

	// Initialize the logger
	logger := logger.Logger()

	//=========================================================================

	// Initialize the database.
	logger.Info("Attempting to connect to the database")

	db, err := orm.BuildGormDb(databaseDsn, mysqlMaxOpenConns, mysqlMaxIDLEConns)
	if err != nil {
		panic(err)
	}

	logger.Info("Connected to the database")

	//=========================================================================
	// Dependency injection for Infrastructure layer.

	authRepo := authDriven.NewRepository(db)
	_ = authRepo
	adminRegionRepo := adminRegionsDriven.NewRepository(db)
	feeSetRepo := feeSetDriven.NewRepository(db)

	//=========================================================================
	//Dependency injections for Application layer.

	// Here we usually inject infrastructure layer. For example the dataabse. The mock object?
	// On the other hand, the application defines it's own interface (primary/driving port - left side of a hexagon),
	// to contract with the presentation layer (http adapter).
	// We're allowed to inject any application that implements corresponding interface. mock object?
	// It is also called the "Dependency Inversion Principle"

	authApp := app.NewAuthApp(authRepo)
	_ = authApp
	administrativeregionApp := app.NewAdminRegion(adminRegionRepo)
	feeSetApp := app.NewFeeSetApp(feeSetRepo)

	//=========================================================================
	//Dependency injections for Presentation layer (primary/driving adapters)

	apiLogGenerator := api_logger.NewApiLogGenerator(appComponent, appEnv, appVersion, logger)
	_ = apiLogGenerator

	// We're allowed to inject any application that implements corresponding interface. mock object?
	authHandler := handler.NewAuthSimple(authApp) // we can return a structure that implements interface
	//authHandler := handler.NewAuth(authApp)		// or interface itself.

	administrativeregionHandler := handler.NewAdminRegionHandler(administrativeregionApp)
	feeSetHandler := handler.NewPSPFeeHandler(feeSetApp, apiLogGenerator)

	// Now just inject handlers to router.
	routerRouter := router.NewRouter(
		authRepo,
		authHandler,
		administrativeregionHandler,
		feeSetHandler,
	)

	//=========================================================================
	// Init Gin Engine
	//=========================================================================

	// Create a channel to listen for OS signals
	sigChan := make(chan os.Signal, 1)
	// Notify the channel when an interrupt or termination signal is received
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	//----------------------------------------------------------------------

	gin.SetMode(appMode)
	g := gin.Default()

	g.Use(middleware.CORSMiddleware())
	g.Use(middleware.RequestLogger())

	//----------------------------------------------------------------------

	err = routerRouter.Register(g)
	if err != nil {
		logger.Fatal("ERROR: routerRouter.Register: " + err.Error())
		os.Exit(1)
	}

	ginPort, ok := os.LookupEnv("GINPORT")
	if !ok {
		ginPort = "8080"
	}

	srv := &http.Server{
		Addr:         "0.0.0.0:" + ginPort,
		Handler:      g,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		logger.Info("Starting HTTP server on port " + ginPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("ERROR starting server: " + err.Error())
			os.Exit(1)
		}
	}()

	sig := <-sigChan
	logger.Info("Received signal: " + sig.String())

	// Create a context with a timeout for graceful shutdown
	// You can specify a timeout to forcefully shut down the server
	// after a certain duration if it takes too long to stop.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shutdown the server
	logger.Info("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		logger.Info("Server forced to shutdown: " + err.Error())
	} else {
		logger.Info("Server shutdown gracefully.")
	}

	//------------------------------------------------
	// Disconnect from the database?
	//------------------------------------------------
	// Implementation.

	//------------------------------------------------

	logger.Sync()

}
