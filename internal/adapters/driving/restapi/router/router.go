// Package router ...
package router

import (
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driving/restapi/handler"
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driving/restapi/middleware"
	"github.com/andriykusevol/aktemplategorm/internal/domain/sport"

	"github.com/gin-gonic/gin"
)

// Router ...
type Router interface {
	Register(app *gin.Engine) error
}

type router struct {
	authRepo                        sport.AuthRepository
	authHandler                     handler.AuthHandler
	administrativeregionHandler     handler.AdminRegionHandler
	mobileMoneyWithdrawalFeeHandler handler.PSPFeeHandler
}

// NewRouter ...
func NewRouter(
	authRepo sport.AuthRepository,
	authHandler handler.AuthHandler,
	administrativeregionHandler handler.AdminRegionHandler,
	mobileMoneyWithdrawalFeeHandler handler.PSPFeeHandler,
) Router {
	return &router{
		authRepo:                        authRepo,
		authHandler:                     authHandler,
		administrativeregionHandler:     administrativeregionHandler,
		mobileMoneyWithdrawalFeeHandler: mobileMoneyWithdrawalFeeHandler,
	}
}

func (r *router) Register(app *gin.Engine) error {
	//app.Use(middleware.InternalServerErrorMiddleware())
	//app.Use(middleware.RequestHeadersLogger())
	//app.Use(middleware.ResponseHeadersLogger())
	//app.Use(middleware.LogResponseMiddleware())
	app.Use(middleware.DateHeaderMiddleware())
	app.Use(middleware.XCorrelationIDMiddleware())
	app.Use(middleware.ContentLengthMiddleware())
	r.RegisterAPI(app)
	return nil
}
