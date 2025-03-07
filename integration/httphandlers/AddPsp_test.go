package httphandlers

import (
	"fmt"

	"github.com/stretchr/testify/assert"

	"github.com/andriykusevol/aktemplategorm/internal/adapters/driven/gorm"
	authDriven "github.com/andriykusevol/aktemplategorm/internal/adapters/driven/gorm/auth"
	feeSetDriven "github.com/andriykusevol/aktemplategorm/internal/adapters/driven/gorm/feeset"
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driving/restapi/handler"
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driving/restapi/middleware"
	"github.com/gin-gonic/gin"

	//"github.com/andriykusevol/aktemplategorm/internal/adapters/driving/restapi/handler"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/andriykusevol/aktemplategorm/internal/application/app"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Function to make HTTP request
func makeRequest(method, url, apiKey string, payload io.Reader) (int, string, error) {
	client := &http.Client{}
	//req, err := http.NewRequest(method, url, bytes.NewBufferString(payload))
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return 0, "", err
	}

	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return 0, "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, "", err
	}

	return res.StatusCode, string(body), nil
}

func Test_AddPsp(t *testing.T) {

	//appMode := os.Getenv("APP_MODE")
	//appPort := os.Getenv("APP_PORT")
	databaseDsn := os.Getenv("DATABASE_DSN")
	mysqlMaxOpenConns := os.Getenv("MYSQL_MAX_OPENCONNS")
	mysqlMaxIDLEConns := os.Getenv("MYSQL_MAX_IDLECONS")
	//appComponent := os.Getenv("COMPONENT")
	//appVersion := os.Getenv("API_VERSION")
	//appEnv := os.Getenv("ENV")

	db, err := orm.BuildGormDb(databaseDsn, mysqlMaxOpenConns, mysqlMaxIDLEConns)
	if err != nil {
		panic(err)
	}
	sqldb, _ := db.DB()
	defer sqldb.Close()

	authRepo := authDriven.NewRepository(db)
	authApp := app.NewAuthApp(authRepo)
	//authHandler := handler.NewAuthSimple(authApp)
	//authHandler.Login()

	ctx := context.TODO()
	authDomainObject, _ := authApp.GenerateToken(ctx, "User")
	apiKey := "Bearer " + authDomainObject.AccessToken

	//=================================================

	feeSetRepo := feeSetDriven.NewRepository(db)
	feeSetApp := app.NewFeeSetApp(feeSetRepo)
	feeSetHandler := handler.NewPSPFeeHandler(feeSetApp, nil)

	//=================================================

	gin.SetMode(gin.TestMode)
	// Mock server
	router := gin.New()

	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.RequestLogger())
	router.Use(middleware.AuthMiddleware(authRepo))
	router.POST("/pspfee", feeSetHandler.CreatePSP)

	server := httptest.NewServer(router)

	defer server.Close()

	//=================================================

	tests := []struct {
		name       string
		method     string
		path       string
		apiKey     string
		payload    *strings.Reader
		wantStatus int
	}{
		{
			name:   "Valid request",
			method: "POST",
			path:   "/pspfee",
			apiKey: apiKey,
			payload: strings.NewReader(`{
  				"PspCode": "CG-MTN-MTN",
  				"PspCountryCode": "CM",
  				"PspShortName": "CG-MTN-MTN_1 Short name"			
			}`),
			wantStatus: http.StatusCreated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status, body, err := makeRequest(tt.method, server.URL+tt.path, tt.apiKey, tt.payload)
			_ = body
			assert.NoError(t, err)
			assert.Equal(t, tt.wantStatus, status)
			// if tt.wantBody != "" {
			// 	assert.JSONEq(t, tt.wantBody, body)
			// }
		})
	}

}
