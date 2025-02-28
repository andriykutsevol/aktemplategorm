package request

import (
	"encoding/json"
	"errors"
	apierror "github.com/andriykusevol/aktemplategorm/internal/adapters/driving/errors"
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driving/restapi/internal"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/queryfilter"
	"strconv"
)

// =========================================================
// We define two functions here
// It is used by the:
// 		PspQueryByMap

// 		func ToDomain_QueryMapFilter(queryParams map[string][]string,
// 			vff func(queryfilter.FilterFields) error,
// 			vof func(queryfilter.OrderFields) error,
//		) (*queryfilter.QueryFilter, *apierror.ApiError)

// and
// It is used by the:
// 		PspQueryByJson

// 		func (qf *QueryFilter) ToDomain_QueryFilter(			// QueryFilter structure here is because we can define it for resquest
// 			vff func(queryfilter.FilterFields) error,
// 			vof func(queryfilter.OrderFields) error,
// 		) (*queryfilter.QueryFilter, *apierror.ApiError)

// The both returns QueryFilter domain object.

// ==========================================================
// It is used by the:
// 		PspQueryByMap
//
// ==========================================================

// Here we delete from map in a function, let's see will it affect caller's map.
// The answer is - we need to pass a POINTER here to modify caller's map
func paginationBuild(req map[string]interface{}) (map[string]interface{}, *queryfilter.PaginationQuery, error) {

	pq := queryfilter.PaginationQuery{}
	var errorString string

	// --------------------------------------------------------------------------------

	if req["CurrentPage"] == nil { // Yes, it will be nil if it does not present in c.Request.URL.Query()
		pq.CurrentPage = 1 // Set default value from openapi
	} else {
		if val, ok := req["CurrentPage"].(string); ok {
			v, err := internal.StringToUint(val)
			if err != nil {
				errorString += "Wrong pagination query: CurrentPage; "
			} else {
				pq.CurrentPage = v
			}
		} else {
			errorString += "Wrong pagination query: CurrentPage; "
		}
		delete(req, "CurrentPage")
	}

	// --------------------------------------------------------------------------------

	if req["PageSize"] == nil {
		pq.PageSize = 10 // Set default value from openapi
	} else {
		if val, ok := req["PageSize"].(string); ok {
			v, err := internal.StringToUint(val)
			if err != nil {
				errorString += "Wrong pagination query: PageSize; "
			} else {
				pq.PageSize = v
			}
		} else {
			errorString += "Wrong pagination query: PageSize; "
		}
		delete(req, "PageSize")
	}

	// --------------------------------------------------------------------------------

	if req["OnlyCount"] == nil {
		pq.OnlyCount = false // Set default value from openapi
	} else {

		if val, ok := req["OnlyCount"].(string); ok {
			v, err := strconv.ParseBool(val)
			if err != nil {
				errorString += "Wrong pagination query: OnlyCount; "
			} else {
				pq.OnlyCount = v
			}
		} else {
			errorString += "Wrong pagination query: OnlyCount; "
		}
		delete(req, "OnlyCount")

	}

	// --------------------------------------------------------------------------------
	if len(errorString) > 0 {
		return req, nil, errors.New(errorString)
	}
	return req, &pq, nil
}

func sortBuild(req map[string]interface{}) (map[string]interface{}, *queryfilter.OrderFields, error) {

	var orderFields *queryfilter.OrderFields
	var errorString string

	if req["Sort"] != nil {
		var sortarr []string

		if val, ok := req["Sort"].([]string); ok {
			sortarr = val
		} else if val, ok := req["Sort"].(string); ok {
			sortarr = append(sortarr, val)
		} else {
			errorString += "Wrong query: Sort "
		}

		var err error
		orderFields, err = queryfilter.OrderFieldsFromStrings(sortarr)
		if err != nil {
			errorString += "Wrong query: Sort "
		}
		delete(req, "Sort")
	}

	if len(errorString) > 0 {
		return req, nil, errors.New(errorString)
	}

	return req, orderFields, nil
}

// It is for the case when we do not use json in a query params,
// but use separate fields, and we are still do not want to creare separate request in the backend.
// In this case we convert raw map[string][]string from query to the queryfilter_vo (domain object)
func ToDomain_QueryMapFilter(queryParams map[string][]string,
	vff func(queryfilter.FilterFields) error,
	vof func(queryfilter.OrderFields) error) (*queryfilter.QueryFilter, *apierror.ApiError) {

	//-------------------------------------------------
	// Build map
	req := make(map[string]interface{})
	for key, values := range queryParams {
		if len(values) > 1 {
			// If multiple values exist, return them as an array
			req[key] = values
		} else {
			// If only one value exists, return it as a single value
			req[key] = values[0]
		}
	}

	//-------------------------------------------------
	// Pagination build

	// fmt.Println("--------ToDomain_QueryMapFilter------")
	// for key, values := range req {
	// 	fmt.Println(key, ":::", values)
	// }

	req, pq, err := paginationBuild(req)
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_INVALID_DATA_FORMAT,
			"Error in buildPagination(req)",
			apierror.SEVERITY_LOW, err,
		)
		return nil, &errObj
	}

	//-------------------------------------------------
	// Sort build

	req, orderFields, err := sortBuild(req)
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_INVALID_DATA_FORMAT,
			"Error in sortBuild(req)",
			apierror.SEVERITY_LOW, err,
		)
		return nil, &errObj
	}

	// We tread all other fields as filter parameters.
	//-------------------------------------------------
	// Filter build

	filterFields := queryfilter.FilterFieldsFromMap(req)
	//fmt.Println("filterfields: ", filterFields)

	domainItem, err := queryfilter.NewQueryMapFilter(*pq, filterFields, vff, orderFields, vof)
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_INVALID_DATA_FORMAT,
			"The provided data is invalid. Domain error.",
			apierror.SEVERITY_LOW, err,
		)
		return nil, &errObj
	}

	return domainItem, nil
}

//==========================================================
// It is used by the:
// 		PspQueryByJson
//==========================================================

type QueryFilter struct {
	PaginationQuery queryfilter.PaginationQuery
	Filter          *string `form:"filter"`
	Sort            *string `form:"sort"`
}

func (qf *QueryFilter) ToDomain_QueryFilter(
	vff func(queryfilter.FilterFields) error,
	vof func(queryfilter.OrderFields) error,
) (*queryfilter.QueryFilter, *apierror.ApiError) {

	var filterMap map[string]any
	var sortSlice []string

	if qf.Filter != nil {
		filterQuery := []byte(*qf.Filter)
		err := json.Unmarshal(filterQuery, &filterMap)
		if err != nil {
			errObj := apierror.NewApiError(
				"",
				apierror.ERROR_CODE_INVALID_DATA_FORMAT,
				"The provided data is invalid. Did you encode the filter query parameter?",
				apierror.SEVERITY_LOW, err,
			)
			return nil, &errObj
		}
	}

	if qf.Sort != nil {
		sortQuery := []byte(*qf.Sort)

		err := json.Unmarshal(sortQuery, &sortSlice)

		if err != nil {
			errObj := apierror.NewApiError(
				"",
				apierror.ERROR_CODE_INVALID_DATA_FORMAT,
				"The provided data is invalid. Did you encode the sort query parameter?",
				apierror.SEVERITY_LOW, err,
			)
			return nil, &errObj
		}

	}

	domainItem, err := queryfilter.NewQueryFilter(qf.PaginationQuery, filterMap, vff, sortSlice, vof)
	if err != nil {
		errObj := apierror.NewApiError(
			"",
			apierror.ERROR_CODE_INVALID_DATA_FORMAT,
			"The provided data is invalid. Domain error.",
			apierror.SEVERITY_LOW, err,
		)
		return nil, &errObj
	}

	return domainItem, nil
}
