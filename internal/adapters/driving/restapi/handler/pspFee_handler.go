// Package handler to handle http requests.
package handler

import (
	"fmt"
	"github.com/andriykusevol/aktemplategorm/internal/application/pport"

	//"github.com/andriykusevol/aktemplategorm/internal/domain"
	apierror "github.com/andriykusevol/aktemplategorm/internal/adapters/driving/errors"
	api_logger "github.com/andriykusevol/aktemplategorm/internal/adapters/driving/logger"
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driving/restapi/internal"
	"github.com/andriykusevol/aktemplategorm/internal/domain/aggregate/feeset"
	"github.com/andriykusevol/aktemplategorm/internal/domain/entity/psp"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/currency"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/money"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/patch"

	"github.com/gin-gonic/gin"
	"github.com/govalues/decimal"

	"errors"
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driving/restapi/request"
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driving/restapi/response"
	"net/http"
	"strconv"
	"time"
)

type PSPFeeHandler interface {
	CreatePSP(c *gin.Context)
	GetPSP(c *gin.Context)

	PspPatchByID(c *gin.Context)
	PspPatchByrray(c *gin.Context)
	PspPatchByQuery(c *gin.Context)

	PspQueryByMap(c *gin.Context)
	PspQueryByJson(c *gin.Context)
	PspQueryByRequest(c *gin.Context)

	ListPSP(c *gin.Context)
	DeletePSP(c *gin.Context)

	AddMobileProviderFeeSet(c *gin.Context)
	ListMobileProviderFeeSet(c *gin.Context)
	GetMobileProviderFeeSet(c *gin.Context)
	DeleteMobileProviderFeeSet(c *gin.Context)

	AddMobileProviderFeeRange(c *gin.Context)
	ListFeeSetRange(c *gin.Context)
	GetFeeSetRange(c *gin.Context)
	DeleteMobileProviderFeeRange(c *gin.Context)

	PatchFeeRange(c *gin.Context)

	CalculateFeeForAmount(c *gin.Context)
	CalculateListFeeForAmount(c *gin.Context)
}

type appsPSPFeeHandler struct {
	feeSet              pport.FeeSetApp
	apiLogGenerator     api_logger.ApiLoggerIface
	converterFromDomain response.PspConverterFromDomain
}

func NewPSPFeeHandler(
	feeSetApp pport.FeeSetApp,
	apiLogGenerator api_logger.ApiLoggerIface) PSPFeeHandler {
	// This is not testable code, but I do not want to inject it from the di_app.go
	cf := response.NewPspConverterFromDomain()
	return &appsPSPFeeHandler{
		feeSet:              feeSetApp,
		apiLogGenerator:     apiLogGenerator,
		converterFromDomain: cf,
	}
}

// PATCH
func (apps *appsPSPFeeHandler) PspPatchByID(c *gin.Context) {

	ctx := c.Request.Context()
	_ = ctx

	idparam := c.Param("ID")
	if len(idparam) == 0 {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_MISSING_REQUIRED_PARAMETER,
			"Missing Required Parameter: ID",
			apierror.SEVERITY_LOW, errors.New(""),
		)
		c.JSON(http.StatusBadRequest, errObj)
		return
	}

	id, err := internal.StringToUint(idparam)
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_INVALID_DATA_FORMAT,
			"Cannot convert ID parameter to uint",
			apierror.SEVERITY_LOW, errors.New(""),
		)
		c.JSON(http.StatusBadRequest, errObj)
		return
	}

	p := patch.NewPatch(psp.ValidatePatch, psp.ToDomain)

	//We do not have any annotations here.
	if err := c.ShouldBindJSON(p.Data()); err != nil {
		//TODO: handle error
	}

	if p.ValidatePatch() {
		_ = p.ToDomain(ctx)
		updatedpsp, err := apps.feeSet.PspUpdateByID(ctx, id, p)

		if err != nil {
			errObj := apierror.NewApiError(
				"",
				apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
				"Unable to update Psp",
				apierror.SEVERITY_LOW, err,
			)
			c.JSON(http.StatusInternalServerError, errObj)
			return
		}

		if updatedpsp == nil {
			errObj := apierror.NewApiError(
				"",
				apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
				"It looks like you deleted an item",
				apierror.SEVERITY_LOW, errors.New(""),
			)
			c.JSON(http.StatusInternalServerError, errObj)
			return
		}

		psp_response, err := response.FromDomain_Psp(*updatedpsp)
		if err != nil {
			errObj := apierror.NewApiError(
				"",
				apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
				"Cannot convert domain object to response object",
				apierror.SEVERITY_LOW, err,
			)
			c.JSON(http.StatusInternalServerError, errObj)
			return
		}

		c.JSON(http.StatusOK, psp_response)
		return

	} else {
		//TODO: Handle pach is not valid
	}

}

func (apps *appsPSPFeeHandler) PspPatchByrray(c *gin.Context) {

	ctx := c.Request.Context()
	_ = ctx

	p := patch.NewPatch(psp.ValidatePatch, psp.ToDomain)

	//We do not have any annotations here.
	if err := c.ShouldBindJSON(p.Data()); err != nil {
		//TODO: handle error
	}

	p.ValidatePatch()

	// if p.ValidatePatch() {
	// 	h.feeSetApp.UpdateFeeRange(ctx, p)
	// } else {
	// 	//TODO: Handle pach is not valid
	// }

}

func (apps *appsPSPFeeHandler) PspPatchByQuery(c *gin.Context) {

	ctx := c.Request.Context()
	_ = ctx

	p := patch.NewPatch(psp.ValidatePatch, psp.ToDomain)

	//We do not have any annotations here.
	if err := c.ShouldBindJSON(p.Data()); err != nil {
		//TODO: handle error
	}

	p.ValidatePatch()

	// if p.ValidatePatch() {
	// 	h.feeSetApp.UpdateFeeRange(ctx, p)
	// } else {
	// 	//TODO: Handle pach is not valid
	// }

}

// POST
func (apps *appsPSPFeeHandler) CreatePSP(c *gin.Context) {

	correlationID, _ := c.Get("X-Correlation-Id")
	timestamp := api_logger.ApiLogDataField{
		Key:   api_logger.TIMESTAMP,
		Value: time.Now(),
	}
	cid := api_logger.ApiLogDataField{
		Key:   api_logger.CORRELATION_ID,
		Value: correlationID,
	}

	apiLogData := apps.apiLogGenerator.CreateApiLogData(timestamp, cid)
	_ = apiLogData
	// apiLogData[api_logger.API_VERSION]
	// apiLogData[api_logger.COMPONENT]
	// apiLogData[api_logger.ENVIRONMENT]
	// apiLogData[api_logger.TIMESTAMP]
	// apiLogData[api_logger.CORRELATION_ID]
	//------------------------------------------------------

	ctx := c.Request.Context()
	uid, ok := c.Get("UserID")
	if !ok {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_INVALID_AUTHENTICATION_CREDENTIALS,
			"http.StatusUnauthorized zzz",
			apierror.SEVERITY_LOW, errors.New(""),
		)
		c.JSON(http.StatusUnauthorized, errObj)
		return
	}

	userID := uid.(string)

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

	var requestItem request.PSPObject

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

	psp := requestItem.ToDomain_PSPObject(userID)

	domainItem, err := apps.feeSet.PspAdd(ctx, psp)
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
			"Unexpected server error",
			apierror.SEVERITY_LOW, err,
		)
		c.JSON(http.StatusInternalServerError, errObj)
		return
	}

	fmt.Println("555")
	fmt.Println("domainItem.PspCode: ", domainItem.PspCode)

	//psp_response, err := response.FromDomain_Psp(*domainItem)
	domainItemValue := *domainItem
	psp_response, err := apps.converterFromDomain.FromDomain_Psp(domainItemValue)
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
			"Unexpected server error",
			apierror.SEVERITY_LOW, err,
		)
		c.JSON(http.StatusInternalServerError, errObj)
		return
	}

	fmt.Println("666")

	//	If you already have a JSON string and want to send it directly:
	//	c.String(http.StatusOK, jsonString) // Respond with raw JSON string
	//  c.Header("Content-Type", "application/json") // Set the header to application/json
	c.JSON(http.StatusCreated, psp_response)
	return
}

func (apps *appsPSPFeeHandler) GetPSP(c *gin.Context) {

	ctx := c.Request.Context()
	_ = ctx

	correlationID, _ := c.Get("X-Correlation-Id")
	timestamp := api_logger.ApiLogDataField{
		Key:   api_logger.TIMESTAMP,
		Value: time.Now(),
	}
	cid := api_logger.ApiLogDataField{
		Key:   api_logger.CORRELATION_ID,
		Value: correlationID,
	}

	apiLogData := apps.apiLogGenerator.CreateApiLogData(timestamp, cid)
	_ = apiLogData
	// apiLogData[api_logger.API_VERSION]
	// apiLogData[api_logger.COMPONENT]
	// apiLogData[api_logger.ENVIRONMENT]
	// apiLogData[api_logger.TIMESTAMP]
	// apiLogData[api_logger.CORRELATION_ID]
	//------------------------------------------------------

	idparam := c.Param("ID")
	id, err := strconv.ParseUint(idparam, 10, 64)
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
			"Unexpected server error",
			apierror.SEVERITY_LOW, err,
		)
		c.JSON(http.StatusInternalServerError, errObj)
		return
	}

	domainItem, err := apps.feeSet.GetPSP(ctx, uint(id))
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
			"Unexpected server error",
			apierror.SEVERITY_LOW, err,
		)
		c.JSON(http.StatusInternalServerError, errObj)
		return
	}

	psp_response, err := response.FromDomain_Psp(*domainItem)
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
			"Unexpected server error",
			apierror.SEVERITY_LOW, err,
		)
		c.JSON(http.StatusInternalServerError, errObj)
		return
	}

	c.JSON(http.StatusCreated, psp_response)
	return
}

func (apps *appsPSPFeeHandler) PspQueryByMap(c *gin.Context) {

	ctx := c.Request.Context()
	_ = ctx

	// type Values map[string][]string
	queryParams := c.Request.URL.Query()

	domainQueryFilter, errObj := request.ToDomain_QueryMapFilter(queryParams,
		psp.ValidateFilterFields, psp.ValidateOrderFields)
	if errObj != nil {
		c.JSON(http.StatusBadRequest, errObj)
		return
	}

	//validate filter map
	err := domainQueryFilter.ValidateFilterFields()
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_INVALID_DATA_FORMAT,
			"The provided data is invalid. Please make sure you have the correct filter fields",
			apierror.SEVERITY_LOW, err,
		)
		c.JSON(http.StatusBadRequest, errObj)
		return
	}
	err = domainQueryFilter.ValidateOrderFields()
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_INVALID_DATA_FORMAT,
			"The provided data is invalid. Please make sure you have the correct sort fields",
			apierror.SEVERITY_LOW, err,
		)
		c.JSON(http.StatusBadRequest, errObj)
		return
	}

	domainItems, domainPagination, err := apps.feeSet.QueryFilterPSP(ctx, *domainQueryFilter)
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
			"Unexpected server error",
			apierror.SEVERITY_LOW, err,
		)
		c.JSON(http.StatusInternalServerError, errObj)
		return
	}

	responseItems, _ := response.FromDomain_PaginatedPspList(domainItems, *domainPagination)
	c.JSON(http.StatusOK, responseItems)

}

// GET
func (apps *appsPSPFeeHandler) PspQueryByJson(c *gin.Context) {
	ctx := c.Request.Context()
	_ = ctx

	var queryfilter request.QueryFilter
	if err := c.ShouldBindQuery(&queryfilter); err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_DATA_VALIDATION_FAILURE,
			"Unable to bind query parameters: "+err.Error(),
			apierror.SEVERITY_LOW, err,
		)
		c.JSON(http.StatusBadRequest, errObj)
		return
	}
	fmt.Println("")

	domainQueryFilter, apierr := queryfilter.ToDomain_QueryFilter(
		psp.ValidateFilterFields, psp.ValidateOrderFields)
	if apierr != nil {
		c.JSON(http.StatusBadRequest, apierr)
		return
	}

	//validate filter map
	err := domainQueryFilter.ValidateFilterFields()
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_INVALID_DATA_FORMAT,
			"The provided data is invalid. Please make sure you have the correct filter fields",
			apierror.SEVERITY_LOW, err,
		)
		c.JSON(http.StatusOK, errObj)
		return
	}
	err = domainQueryFilter.ValidateOrderFields()
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_INVALID_DATA_FORMAT,
			"The provided data is invalid. Please make sure you have the correct sort fields",
			apierror.SEVERITY_LOW, err,
		)
		c.JSON(http.StatusOK, errObj)
		return
	}

	domainItems, domainPagination, err := apps.feeSet.QueryFilterPSP(ctx, *domainQueryFilter)
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
			"Unexpected server error",
			apierror.SEVERITY_LOW, err,
		)
		c.JSON(http.StatusInternalServerError, errObj)
		return
	}

	responseItems, _ := response.FromDomain_PaginatedPspList(domainItems, *domainPagination)
	c.JSON(http.StatusOK, responseItems)
	return

}

// GET
func (apps *appsPSPFeeHandler) PspQueryByRequest(c *gin.Context) {

	ctx := c.Request.Context()
	_ = ctx

	var requestQuery request.PSPQuery
	if err := c.ShouldBindQuery(&requestQuery); err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_DATA_VALIDATION_FAILURE,
			"Query parameters are wrong: "+err.Error(),
			apierror.SEVERITY_LOW, err,
		)
		c.JSON(http.StatusBadRequest, errObj)
		return
	}

	domainQuery := requestQuery.ToDomain_PSPQuery()

	fmt.Println("=======Sort========")
	fmt.Println(requestQuery.Sort)

	fmt.Println("=======Filter========")
	fmt.Println(requestQuery.IDs)
	fmt.Println(requestQuery.PspCode)
	fmt.Println(requestQuery.PspCountryCode)

	fmt.Println("=======Pagination========")
	fmt.Println(requestQuery.PaginationParam.CurrentPage)
	fmt.Println(requestQuery.PaginationParam.PageSize)
	fmt.Println(requestQuery.PaginationParam.OnlyCount)

	domainItems, domainPagination, err := apps.feeSet.QueryPSP(ctx, *domainQuery)
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
			"Unexpected server error",
			apierror.SEVERITY_LOW, err,
		)
		c.JSON(http.StatusInternalServerError, errObj)
		return
	}

	responseItems, _ := response.FromDomain_PaginatedPspList(domainItems, *domainPagination)

	c.JSON(http.StatusOK, responseItems)
	return
}

// GET
// !!! Instead of this we have a Query requests for now.
func (apps *appsPSPFeeHandler) ListPSP(c *gin.Context) {

	ctx := c.Request.Context()

	var psp_id *uint
	var psp_code *string

	psp_id_query := c.Query("pspID")

	if len(psp_id_query) != 0 {

		id, err := internal.StringToUint(psp_id_query)
		if err != nil {
			errObj := apierror.NewApiError(
				"",
				apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
				"Unexpected server error",
				apierror.SEVERITY_LOW, errors.New(""),
			)
			c.JSON(http.StatusInternalServerError, errObj)
			return
		}
		psp_id = &id
	} else {
		psp_id = nil
	}

	psp_code_query := c.Query("pspCode")

	if len(psp_code_query) != 0 {
		psp_code = &psp_code_query
	} else {
		psp_code = nil
	}

	psps, err := apps.feeSet.ListPSP(ctx, psp_id, psp_code)
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
			"Unexpected server error",
			apierror.SEVERITY_LOW, errors.New(""),
		)
		c.JSON(http.StatusInternalServerError, errObj)
		return
	}

	response, err := response.FromDomainList_PspList(psps)
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
			"Unexpected server error",
			apierror.SEVERITY_LOW, errors.New(""),
		)
		c.JSON(http.StatusInternalServerError, errObj)
		return
	}

	//	If you already have a JSON string and want to send it directly:
	//	c.String(http.StatusOK, jsonString) // Respond with raw JSON string
	//	c.Header("Content-Type", "application/json") // Set the header to application/json

	c.JSON(http.StatusOK, response)
	return
}

func (apps *appsPSPFeeHandler) DeletePSP(c *gin.Context) {

	ctx := c.Request.Context()

	psp_id := c.Query("pspID")
	if len(psp_id) == 0 {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_MISSING_REQUIRED_PARAMETER,
			"Missing Required Parameter: psp_code",
			apierror.SEVERITY_LOW, errors.New(""),
		)
		c.JSON(http.StatusBadRequest, errObj)
		return
	}

	pid, err := internal.StringToUint(psp_id)
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
			"Unexpected server error",
			apierror.SEVERITY_LOW, err,
		)
		c.JSON(http.StatusInternalServerError, errObj)
		return
	}

	//It's a GET request so our application recieves just parameters.
	//TODO: set DeletedBy from auth middleware.
	err = apps.feeSet.DeletePSP(ctx, pid)

	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
			"Unexpected server error",
			apierror.SEVERITY_LOW, err,
		)
		c.JSON(http.StatusInternalServerError, errObj)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
	return
}

//---------------------------------------------------------------
// PSPFee Set Management API
//---------------------------------------------------------------

// POST
func (apps *appsPSPFeeHandler) AddMobileProviderFeeSet(c *gin.Context) {
	// Create new Fee Set. A set once created cannot be modified! Any modification will create a new set.

	correlationID, _ := c.Get("X-Correlation-Id")

	timestamp := api_logger.ApiLogDataField{
		Key:   api_logger.TIMESTAMP,
		Value: time.Now(),
	}
	cid := api_logger.ApiLogDataField{
		Key:   api_logger.CORRELATION_ID,
		Value: correlationID,
	}
	apiLogData := apps.apiLogGenerator.CreateApiLogData(timestamp, cid)
	_ = apiLogData
	// apiLogData[api_logger.API_VERSION]
	// apiLogData[api_logger.COMPONENT]
	// apiLogData[api_logger.ENVIRONMENT]
	// apiLogData[api_logger.TIMESTAMP]
	// apiLogData[api_logger.CORRELATION_ID]
	//------------------------------------------------------

	ctx := c.Request.Context()
	uid, ok := c.Get("UserID")
	if !ok {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_INVALID_AUTHENTICATION_CREDENTIALS,
			"http.StatusUnauthorized: Cannot find UserID in a token.",
			apierror.SEVERITY_LOW, errors.New(""),
		)
		c.JSON(http.StatusUnauthorized, errObj)
		return
	}
	userID := uid.(string)

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

	var requestItem request.PSPFeeSet

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

	// TODO: set the AuditDBRecord.CreatedBy from the gin.Contex
	// (which in turn have to be set by gin middleware based on authentication method)

	feesetAggregate := requestItem.ToDomain_PSPFeeSet(userID)
	err := apps.feeSet.AddMobileProviderFeeSet(ctx, feesetAggregate)
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
			"Unexpected server error",
			apierror.SEVERITY_LOW, err,
		)
		c.JSON(http.StatusInternalServerError, errObj)
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
	return
}

func (apps *appsPSPFeeHandler) ListMobileProviderFeeSet(c *gin.Context) {
	// List a Fee Sets for specific PSP.
}

func (apps *appsPSPFeeHandler) GetMobileProviderFeeSet(c *gin.Context) {
	// Get a specific Fee Set.
}

func (apps *appsPSPFeeHandler) DeleteMobileProviderFeeSet(c *gin.Context) {
	// Set a Fee Set as deleted (is_active = FALSE)
}

//---------------------------------------------------------------
// PSPFee Range Management API
//---------------------------------------------------------------

// POST
func (apps *appsPSPFeeHandler) AddMobileProviderFeeRange(c *gin.Context) {
	// Add Fee Range for specifict Fee Set.

	correlationID, _ := c.Get("X-Correlation-Id")

	timestamp := api_logger.ApiLogDataField{
		Key:   api_logger.TIMESTAMP,
		Value: time.Now(),
	}
	cid := api_logger.ApiLogDataField{
		Key:   api_logger.CORRELATION_ID,
		Value: correlationID,
	}
	apiLogData := apps.apiLogGenerator.CreateApiLogData(timestamp, cid)
	_ = apiLogData
	// apiLogData[api_logger.API_VERSION]
	// apiLogData[api_logger.COMPONENT]
	// apiLogData[api_logger.ENVIRONMENT]
	// apiLogData[api_logger.TIMESTAMP]
	// apiLogData[api_logger.CORRELATION_ID]
	//------------------------------------------------------

	ctx := c.Request.Context()

	uid, ok := c.Get("UserID")
	if !ok {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_INVALID_AUTHENTICATION_CREDENTIALS,
			"http.StatusUnauthorized",
			apierror.SEVERITY_LOW, errors.New(""),
		)
		c.JSON(http.StatusUnauthorized, errObj)
		return
	}
	userID := uid.(string)

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

	var requestItem request.PSPFeeRange

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

	feeRangeEntity := requestItem.ToDomain_PSPFeeRange(userID)

	err := apps.feeSet.AddMobileProviderFeeRange(ctx, *feeRangeEntity)

	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
			"Unexpected server error",
			apierror.SEVERITY_LOW, errors.New(""),
		)
		c.JSON(http.StatusInternalServerError, errObj)
		return
	}

	c.JSON(http.StatusCreated, gin.H{})

}

func (apps *appsPSPFeeHandler) ListFeeSetRange(c *gin.Context) {
	//TODO: Look at the ListPSP
	// List Fee Ranges for specific Fee Set
	ctx := c.Request.Context()

	//TODO: this should be done by restapi package.
	setId, _ := strconv.Atoi(c.Param("pspfeeset_id"))

	apps.feeSet.ListFeeRange(ctx, uint(setId))

}

// GET
func (apps *appsPSPFeeHandler) GetFeeSetRange(c *gin.Context) {

	ctx := c.Request.Context()
	_ = ctx

	feerange_id := c.Query("feerange_id")
	if len(feerange_id) == 0 {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_MISSING_REQUIRED_PARAMETER,
			"Missing Required Parameter: psp_code",
			apierror.SEVERITY_LOW, errors.New(""),
		)
		c.JSON(http.StatusBadRequest, errObj)
		return
	}

	f, err := internal.StringToUint(feerange_id)
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
			"Unexpected server error",
			apierror.SEVERITY_LOW, err,
		)
		c.JSON(http.StatusInternalServerError, errObj)
		return
	}

	//It's a GET request so our application recieves just parameters.
	feerange, err := apps.feeSet.GetFeeRange(ctx, f)
	_ = feerange
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
			"Unexpected server error",
			apierror.SEVERITY_LOW, err,
		)
		c.JSON(http.StatusInternalServerError, errObj)
		return
	}

	response := response.FromDomain_FeeRange(feerange)

	c.JSON(http.StatusOK, response)

}

func (apps *appsPSPFeeHandler) PatchFeeRange(c *gin.Context) {

	ctx := c.Request.Context()
	_ = ctx

	p := patch.NewPatch(feeset.ValidatePatch, feeset.ToDomain)

	//We do not have any annotations here.
	if err := c.ShouldBindJSON(p.Data()); err != nil {
		//TODO: handle error
	}

	// if p.ValidatePatch() {
	// 	h.feeSetApp.UpdateFeeRange(ctx, p)
	// } else {
	// 	//TODO: Handle pach is not valid
	// }

}

func (apps *appsPSPFeeHandler) DeleteMobileProviderFeeRange(c *gin.Context) {
	// Delete row from the Fee Range.
}

//---------------------------------------------------------------
// PSPFee Set Functional API

// Get
func (apps *appsPSPFeeHandler) CalculateFeeForAmount(c *gin.Context) {
	// Calculate and return the withdrawal fees for given amount
	ctx := c.Request.Context()

	correlationID, _ := c.Get("X-Correlation-Id")

	timestamp := api_logger.ApiLogDataField{
		Key:   api_logger.TIMESTAMP,
		Value: time.Now(),
	}
	cid := api_logger.ApiLogDataField{
		Key:   api_logger.CORRELATION_ID,
		Value: correlationID,
	}
	apiLogData := apps.apiLogGenerator.CreateApiLogData(timestamp, cid)
	_ = apiLogData
	// apiLogData[api_logger.API_VERSION]
	// apiLogData[api_logger.COMPONENT]
	// apiLogData[api_logger.ENVIRONMENT]
	// apiLogData[api_logger.TIMESTAMP]
	// apiLogData[api_logger.CORRELATION_ID]
	//------------------------------------------------------

	// pspCode := c.Param("psp_code")
	// amount := c.Param("amount")

	pspCode := c.Query("psp_code")
	if len(pspCode) == 0 {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_MISSING_REQUIRED_PARAMETER,
			"Missing Required Parameter: psp_code",
			apierror.SEVERITY_LOW, errors.New(""),
		)
		c.JSON(http.StatusBadRequest, errObj)
		return
	}

	amount := c.Query("amount")
	if len(pspCode) == 0 {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_MISSING_REQUIRED_PARAMETER,
			"Missing Required Parameter: amount",
			apierror.SEVERITY_LOW, errors.New(""),
		)
		c.JSON(http.StatusBadRequest, errObj)
		return
	}

	amoundDecimal, err := decimal.Parse(amount)

	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
			"Cannot parse amount",
			apierror.SEVERITY_LOW, err,
		)
		c.JSON(http.StatusInternalServerError, errObj)
		return
	}

	cur := currency.NewCurrency("XAF")
	feeMoney := money.NewMoney(amoundDecimal, cur)
	totalamount, fees, err := apps.feeSet.CalculateFeeForAmount(ctx, pspCode, feeMoney)
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
			"Unexpected server error",
			apierror.SEVERITY_LOW, err,
		)
		c.JSON(http.StatusInternalServerError, errObj)
		return
	}

	amount_resp, ok := amoundDecimal.Float64()
	if !ok {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
			"Cannot parse amountDecimal",
			apierror.SEVERITY_LOW, err,
		)
		c.JSON(http.StatusInternalServerError, errObj)
		return
	}

	response := response.NewFeeResponse(amount_resp, fees, totalamount, pspCode)
	c.JSON(http.StatusOK, response)
	return

}

// POST
func (apps *appsPSPFeeHandler) CalculateListFeeForAmount(c *gin.Context) {
	// Calculate and return the withdrawal fees for a list of amounts.

	correlationID, _ := c.Get("X-Correlation-Id")

	timestamp := api_logger.ApiLogDataField{
		Key:   api_logger.TIMESTAMP,
		Value: time.Now(),
	}
	cid := api_logger.ApiLogDataField{
		Key:   api_logger.CORRELATION_ID,
		Value: correlationID,
	}
	apiLogData := apps.apiLogGenerator.CreateApiLogData(timestamp, cid)
	_ = apiLogData
	// apiLogData[api_logger.API_VERSION]
	// apiLogData[api_logger.COMPONENT]
	// apiLogData[api_logger.ENVIRONMENT]
	// apiLogData[api_logger.TIMESTAMP]
	// apiLogData[api_logger.CORRELATION_ID]
	//------------------------------------------------------

	ctx := c.Request.Context()

	var resp response.BulkFeeResponse
	var requestItem request.BulkFeeRequest

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

	resp, err := apps.feeSet.BulkCalculateFeeForAmount(ctx, requestItem)

	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
			"Unexpected server error",
			apierror.SEVERITY_LOW, errors.New(""),
		)
		c.JSON(http.StatusInternalServerError, errObj)
		return
	}

	c.JSON(http.StatusOK, resp)
	return
}
