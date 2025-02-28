package application

import (
	"errors"
	"testing"

	mockrouter "github.com/andriykusevol/aktemplategorm/mocks/internal_/adapters/driving/restapi/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRouterRegister_Success(t *testing.T) {
	// Initialize Gin in test mode
	gin.SetMode(gin.TestMode)
	g := gin.Default()

	// Create a mock instance of your Router
	mockRouter := new(mockrouter.Router)

	// Define expected behavior for the mock
	mockRouter.On("Register", g).Return(nil)

	// Use the mock in place of the real routerRouter
	err := mockRouter.Register(g)

	// Assertions
	assert.NoError(t, err)
	mockRouter.AssertExpectations(t)
}

func TestRouterRegister_Error(t *testing.T) {
	// Initialize Gin in test mode
	gin.SetMode(gin.TestMode)
	g := gin.Default()

	// Create a mock instance of your Router
	mockRouter := new(mockrouter.Router)

	// Define expected behavior for the mock (returning an error)
	expectedError := errors.New("mock register error")
	mockRouter.On("Register", g).Return(expectedError)

	// Use the mock in place of the real routerRouter
	err := mockRouter.Register(g)

	// Assertions
	assert.EqualError(t, err, "mock register error")
	mockRouter.AssertExpectations(t)
}
