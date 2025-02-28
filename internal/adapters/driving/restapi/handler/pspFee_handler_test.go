// Package handler to handle http requests.
package handler

import (
	"bytes"
	"errors"
	"fmt"

	//"fmt"

	//"context"
	"encoding/json"
	//"fmt"
	api_logger "github.com/andriykusevol/aktemplategorm/internal/adapters/driving/logger"
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driving/restapi/request"
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driving/restapi/response"

	//"github.com/andriykusevol/aktemplategorm/internal/adapters/driving/restapi/response"
	"github.com/andriykusevol/aktemplategorm/internal/application/pport"
	psp "github.com/andriykusevol/aktemplategorm/internal/domain/entity/psp"
	"net/http"
	"net/http/httptest"
	"testing"

	loggerMock "github.com/andriykusevol/aktemplategorm/mocks/internal_/adapters/driving/logger"
	pspResponseMock "github.com/andriykusevol/aktemplategorm/mocks/internal_/adapters/driving/restapi/response"
	feesetAppMock "github.com/andriykusevol/aktemplategorm/mocks/internal_/application/pport"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_appsPSPFeeHandler_CreatePSP_Gem(t *testing.T) {
	type fields struct {
		feeSet          pport.FeeSetApp
		apiLogGenerator api_logger.ApiLoggerIface
	}
	type args struct {
		c *gin.Context
	}

	mockFeeSetApp := feesetAppMock.NewFeeSetApp(t)
	_ = mockFeeSetApp
	mockApiLogGenerator := loggerMock.NewApiLoggerIface(t) // The mockery just generated NewApiLoggerIface name based on structure name
	_ = mockApiLogGenerator
	mockResponseConverter := pspResponseMock.NewPspConverterFromDomain(t)
	_ = mockResponseConverter

	handler := &appsPSPFeeHandler{
		feeSet:              mockFeeSetApp,
		apiLogGenerator:     mockApiLogGenerator,
		converterFromDomain: mockResponseConverter,
	}
	_ = handler

	// Test Cases
	testCases := []struct {
		name                                 string
		requestBody                          request.PSPObject
		invalidJSON                          []byte
		userID                               string
		correlationID                        string
		expectedStatus                       int
		expectation_CreateApiLogData         func()
		expectation_apps_feeSet_PspAdd_Error func()
		expectation_apps_feeSet_PspAdd_OK    func()
		expectation_FromDomain_Psp_Ok        func()
		expectation_FromDomain_Psp_Error     func()
	}{
		{
			name: "Unauthorized",
			requestBody: request.PSPObject{
				PspCode: "some string",
			},
			userID:         "",
			correlationID:  "",
			expectedStatus: http.StatusUnauthorized,
			expectation_CreateApiLogData: func() {
				// You need the second set of curly braces {} in map[ApiLogDataFieldName]interface{}{} because you are creating an empty map literal.
				// You might wonder why not just return nil.  While returning nil is sometimes acceptable,
				// returning an empty map is often preferred for these reasons:
				// mockApiLogGenerator.On("CreateApiLogData", mock.Anything, mock.Anything).Return(nil)
				mockApiLogGenerator.On("CreateApiLogData", mock.Anything, mock.Anything).Return(map[api_logger.ApiLogDataFieldName]interface{}{})
			},
		},
		{
			name:           "InvalidJson",
			invalidJSON:    []byte(`this-is-not-json`), // Not valid JSON format
			userID:         "someuser",
			correlationID:  "",
			expectedStatus: http.StatusBadRequest,
			expectation_CreateApiLogData: func() {
				mockApiLogGenerator.On("CreateApiLogData", mock.Anything, mock.Anything).Return(map[api_logger.ApiLogDataFieldName]interface{}{})
			},
		},
		{
			name:           "Missing Required Parameter",
			requestBody:    request.PSPObject{},
			userID:         "someuser",
			correlationID:  "",
			expectedStatus: http.StatusBadRequest,
			expectation_CreateApiLogData: func() {
				mockApiLogGenerator.On("CreateApiLogData", mock.Anything, mock.Anything).Return(map[api_logger.ApiLogDataFieldName]interface{}{})
			},
		},
		{
			name: "apps.feeSet.PspAdd error",
			requestBody: request.PSPObject{
				PspCode: "fail",
			},
			userID:         "someuser",
			correlationID:  "",
			expectedStatus: http.StatusInternalServerError,
			expectation_CreateApiLogData: func() {
				mockApiLogGenerator.On("CreateApiLogData", mock.Anything, mock.Anything).Return(map[api_logger.ApiLogDataFieldName]interface{}{})
			},

			//MatchedBy can be used to match a mock call based on only certain properties
			// from a complex struct or some calculation.
			// It takes a function that will be evaluated with the called argument
			// and will return true when there's a match and false otherwise.
			expectation_apps_feeSet_PspAdd_Error: func() {
				expectedError := errors.New("Application failed to add PSP")
				mockFeeSetApp.On("PspAdd", mock.Anything, mock.MatchedBy(func(pspArg *psp.PSP) bool {
					return pspArg.PspCode == "fail"
				})).Return(nil, expectedError)
			},
		},
		{
			name: "Created",
			requestBody: request.PSPObject{
				PspCode: "sucess",
			},
			userID:         "someuser",
			correlationID:  "",
			expectedStatus: http.StatusCreated,
			expectation_CreateApiLogData: func() {
				mockApiLogGenerator.On("CreateApiLogData", mock.Anything, mock.Anything).Return(map[api_logger.ApiLogDataFieldName]interface{}{})
			},

			expectation_apps_feeSet_PspAdd_OK: func() {
				domainItem := &psp.PSP{
					ID:      func(s uint) *uint { return &s }(uint(10)),
					PspCode: "PspCode",
				}
				mockFeeSetApp.On("PspAdd", mock.Anything, mock.MatchedBy(func(pspArg *psp.PSP) bool {
					return pspArg.PspCode == "sucess"
				})).Return(domainItem, nil)

			},

			expectation_FromDomain_Psp_Ok: func() {
				//expectedError := errors.New("Failed: FromDomain_Psp(*domainItem)")

				// We need this because FromDomain_Psp return a value not a pointer.
				// When functions return pointers, you can use nil as a return value.
				resp := response.Psp{}
				// mockResponseConverter.On("FromDomain_Psp", mock.Anything).Return(resp, nil)
				// mockResponseConverter.On("FromDomain_Psp", mock.AnythingOfType("psp.PSP")).Return(resp, nil)
				mockResponseConverter.On("FromDomain_Psp", mock.MatchedBy(func(pspArg psp.PSP) bool {
					return pspArg.PspCode == "PspCode"
				})).Return(resp, nil)

			},
		},
		{
			name: "response.FromDomain_Psp(*domainItem) error",
			requestBody: request.PSPObject{
				PspCode: "error",
			},
			userID:         "someuser",
			correlationID:  "",
			expectedStatus: http.StatusInternalServerError,
			expectation_CreateApiLogData: func() {
				mockApiLogGenerator.On("CreateApiLogData", mock.Anything, mock.Anything).Return(map[api_logger.ApiLogDataFieldName]interface{}{})
			},

			expectation_apps_feeSet_PspAdd_OK: func() {
				domainItem := &psp.PSP{
					ID:      func(s uint) *uint { return &s }(uint(10)),
					PspCode: "FromDomain_Psp error",
				}
				mockFeeSetApp.On("PspAdd", mock.Anything, mock.MatchedBy(func(pspArg *psp.PSP) bool {
					return pspArg.PspCode == "error"
				})).Return(domainItem, nil)

			},

			expectation_FromDomain_Psp_Error: func() {
				expectedError := errors.New("Failed: FromDomain_Psp(*domainItem)")

				// We need this because FromDomain_Psp return a value not a pointer.
				// When functions return pointers, you can use nil as a return value.
				resp := response.Psp{}
				//mockResponseConverter.On("", mock.Anything).Return(resp, expectedError)
				mockResponseConverter.On("FromDomain_Psp", mock.MatchedBy(func(pspArg psp.PSP) bool {
					return pspArg.PspCode == "FromDomain_Psp error"
				})).Return(resp, expectedError)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			if tc.expectation_CreateApiLogData != nil {
				tc.expectation_CreateApiLogData()
			}
			if tc.expectation_apps_feeSet_PspAdd_Error != nil {
				tc.expectation_apps_feeSet_PspAdd_Error()
			}
			if tc.expectation_apps_feeSet_PspAdd_OK != nil {
				tc.expectation_apps_feeSet_PspAdd_OK()
			}
			if tc.expectation_FromDomain_Psp_Ok != nil {
				fmt.Println("!!! SET: tc.expectation_FromDomain_Psp_Ok")
				tc.expectation_FromDomain_Psp_Ok()
			}
			if tc.expectation_FromDomain_Psp_Error != nil {
				fmt.Println("!!! SET: tc.expectation_FromDomain_Psp_Error")
				tc.expectation_FromDomain_Psp_Error()
			}

			// It is when we want to extract value from request
			// but we should set that value in request - see below
			// setUserIDMiddleware := func(c *gin.Context) {
			// 	if tc.userID != "" {
			// 		if userID, ok := c.Request.Context().Value("UserID").(string); ok {
			// 			c.Set("UserID", userID)
			// 		}
			// 	}
			// 	c.Next()
			// }
			setUserIDMiddleware := func(c *gin.Context) {
				if tc.userID != "" {
					c.Set("UserID", tc.userID)
				}
				c.Next()
			}
			setCXMiddleware := func(c *gin.Context) {
				if tc.correlationID != "" {
					if correlationID, ok := c.Request.Context().Value("X-Correlation-Id").(string); ok {
						c.Set("X-Correlation-Id", correlationID)
					}
				}
				c.Next()
			}

			var bytes_reader *bytes.Reader

			if tc.invalidJSON == nil {
				body, _ := json.Marshal(tc.requestBody)
				bytes_reader = bytes.NewReader(body)
			} else {
				bytes_reader = bytes.NewReader(tc.invalidJSON)
			}

			req := httptest.NewRequest("POST", "/pspsz", bytes_reader)
			req.Header.Set("Content-Type", "application/json")

			// It is for example, we can ommit it and then just use c.Set("UserID", tc.userID) in setUserIDMiddleware
			// req = req.WithContext(context.WithValue(req.Context(), "UserID", tc.userID))
			// req = req.WithContext(context.WithValue(req.Context(), "X-Correlation-Id", "test-correlation-id"))

			// Set up the response recorder
			w := httptest.NewRecorder()

			// Perform the request
			router := gin.New()
			router.Use(setUserIDMiddleware)
			router.Use(setCXMiddleware)
			router.POST("/pspsz", handler.CreatePSP) // Define the route
			router.ServeHTTP(w, req)

			// Assertions
			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}
