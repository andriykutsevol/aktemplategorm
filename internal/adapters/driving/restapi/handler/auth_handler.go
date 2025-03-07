package handler

import (
	"errors"
	"fmt"
	"net/http"

	apierror "github.com/andriykusevol/aktemplategorm/internal/adapters/driving/errors"
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driving/restapi/internal"
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driving/restapi/request"
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driving/restapi/response"
	"github.com/andriykusevol/aktemplategorm/internal/application/pport"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login(c *gin.Context)
}

type appsAuthHandler struct {
	auth pport.AuthApp
}

// We could do it like this, and just substitute it
// in the di_app.go, but in thise case wi'll lose
// the AuthHandler interface rule.
func NewAuthSimple(authApp pport.AuthApp) *appsAuthHandler {
	return &appsAuthHandler{auth: authApp}
}

func NewAuth(authApp pport.AuthApp) AuthHandler {
	return &appsAuthHandler{auth: authApp}
}

func (apps *appsAuthHandler) Login(c *gin.Context) {

	ctx := c.Request.Context()

	if !internal.IsValidBodyJson(c) {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_INVALID_DATA_FORMAT,
			"Invalid request body json format",
			apierror.SEVERITY_LOW, errors.New(""),
		)
		c.JSON(http.StatusBadRequest, errObj)
		return
	}

	var requestItem request.Login
	if err := c.ShouldBindJSON(&requestItem); err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_MISSING_REQUIRED_PARAMETER,
			"Missing Required Parameter: "+err.Error(),
			apierror.SEVERITY_LOW, err,
		)
		c.JSON(http.StatusBadRequest, errObj)
		return
	}

	user, err := apps.auth.Verify(ctx, requestItem.UserName, requestItem.Password)
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_INVALID_AUTHENTICATION_CREDENTIALS,
			err.Error(),
			apierror.SEVERITY_LOW, err,
		)
		c.JSON(http.StatusUnauthorized, errObj)
		return
	}

	fmt.Println("!!! Handler: user.UUID ", user.UUID)

	c.Set("UserID", user.UUID)

	authEntity, err := apps.auth.GenerateToken(ctx, user.UUID.String())

	authResponse, _ := response.FromDomain_Auth(*authEntity)

	c.JSON(http.StatusOK, authResponse)
}
