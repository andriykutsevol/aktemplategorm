// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// PSPFeeHandler is an autogenerated mock type for the PSPFeeHandler type
type PSPFeeHandler struct {
	mock.Mock
}

// AddMobileProviderFeeRange provides a mock function with given fields: c
func (_m *PSPFeeHandler) AddMobileProviderFeeRange(c *gin.Context) {
	_m.Called(c)
}

// AddMobileProviderFeeSet provides a mock function with given fields: c
func (_m *PSPFeeHandler) AddMobileProviderFeeSet(c *gin.Context) {
	_m.Called(c)
}

// CalculateFeeForAmount provides a mock function with given fields: c
func (_m *PSPFeeHandler) CalculateFeeForAmount(c *gin.Context) {
	_m.Called(c)
}

// CalculateListFeeForAmount provides a mock function with given fields: c
func (_m *PSPFeeHandler) CalculateListFeeForAmount(c *gin.Context) {
	_m.Called(c)
}

// CreatePSP provides a mock function with given fields: c
func (_m *PSPFeeHandler) CreatePSP(c *gin.Context) {
	_m.Called(c)
}

// DeleteMobileProviderFeeRange provides a mock function with given fields: c
func (_m *PSPFeeHandler) DeleteMobileProviderFeeRange(c *gin.Context) {
	_m.Called(c)
}

// DeleteMobileProviderFeeSet provides a mock function with given fields: c
func (_m *PSPFeeHandler) DeleteMobileProviderFeeSet(c *gin.Context) {
	_m.Called(c)
}

// DeletePSP provides a mock function with given fields: c
func (_m *PSPFeeHandler) DeletePSP(c *gin.Context) {
	_m.Called(c)
}

// GetFeeSetRange provides a mock function with given fields: c
func (_m *PSPFeeHandler) GetFeeSetRange(c *gin.Context) {
	_m.Called(c)
}

// GetMobileProviderFeeSet provides a mock function with given fields: c
func (_m *PSPFeeHandler) GetMobileProviderFeeSet(c *gin.Context) {
	_m.Called(c)
}

// GetPSP provides a mock function with given fields: c
func (_m *PSPFeeHandler) GetPSP(c *gin.Context) {
	_m.Called(c)
}

// ListFeeSetRange provides a mock function with given fields: c
func (_m *PSPFeeHandler) ListFeeSetRange(c *gin.Context) {
	_m.Called(c)
}

// ListMobileProviderFeeSet provides a mock function with given fields: c
func (_m *PSPFeeHandler) ListMobileProviderFeeSet(c *gin.Context) {
	_m.Called(c)
}

// ListPSP provides a mock function with given fields: c
func (_m *PSPFeeHandler) ListPSP(c *gin.Context) {
	_m.Called(c)
}

// PatchFeeRange provides a mock function with given fields: c
func (_m *PSPFeeHandler) PatchFeeRange(c *gin.Context) {
	_m.Called(c)
}

// PspPatchByID provides a mock function with given fields: c
func (_m *PSPFeeHandler) PspPatchByID(c *gin.Context) {
	_m.Called(c)
}

// PspPatchByQuery provides a mock function with given fields: c
func (_m *PSPFeeHandler) PspPatchByQuery(c *gin.Context) {
	_m.Called(c)
}

// PspPatchByrray provides a mock function with given fields: c
func (_m *PSPFeeHandler) PspPatchByrray(c *gin.Context) {
	_m.Called(c)
}

// PspQueryByJson provides a mock function with given fields: c
func (_m *PSPFeeHandler) PspQueryByJson(c *gin.Context) {
	_m.Called(c)
}

// PspQueryByMap provides a mock function with given fields: c
func (_m *PSPFeeHandler) PspQueryByMap(c *gin.Context) {
	_m.Called(c)
}

// PspQueryByRequest provides a mock function with given fields: c
func (_m *PSPFeeHandler) PspQueryByRequest(c *gin.Context) {
	_m.Called(c)
}

// NewPSPFeeHandler creates a new instance of PSPFeeHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPSPFeeHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *PSPFeeHandler {
	mock := &PSPFeeHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
