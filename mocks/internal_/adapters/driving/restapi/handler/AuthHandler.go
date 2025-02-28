// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// AuthHandler is an autogenerated mock type for the AuthHandler type
type AuthHandler struct {
	mock.Mock
}

// Login provides a mock function with given fields: c
func (_m *AuthHandler) Login(c *gin.Context) {
	_m.Called(c)
}

// NewAuthHandler creates a new instance of AuthHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthHandler {
	mock := &AuthHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
