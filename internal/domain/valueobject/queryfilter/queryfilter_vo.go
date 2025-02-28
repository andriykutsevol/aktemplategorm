package queryfilter

import (
	"fmt"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/domain_error"
	"strings"
)

// We use this domain structure in response directly
type Pagination struct {
	CurrentPage uint  `json:"CurrentPage"`
	PageSize    uint  `json:"PageSize"`
	PagesCount  uint  `json:"PagesCount"`
	HasMore     bool  `json:"HasMore"`
	ItemsCount  int64 `json:"ItemsCount"`
}

// We use this domain structure to bind query parameters directly.
type PaginationQuery struct {
	Pagination  bool `form:"-"` // Pagination (We alsway do paginagion for Query request)
	OnlyCount   bool `form:"OnlyCount,default=false"`
	CurrentPage uint `form:"CurrentPage,default=1"`
	PageSize    uint `form:"PageSize,default=10" binding:"max=100"`
}

func (a PaginationQuery) GetCurrent() uint {
	return a.CurrentPage
}

func (a PaginationQuery) GetPageSize() uint {
	page_size := a.PageSize
	if a.PageSize == 0 {
		page_size = 100
	}
	return page_size
}

type QueryFilter struct {
	PaginationParam      PaginationQuery
	FilterFields         *FilterFields
	validateFilterFields func(FilterFields) error
	OrderFields          *OrderFields
	validateOrderFields  func(OrderFields) error
}

func NewQueryFilter(qp PaginationQuery,
	qf map[string]any,
	vff func(FilterFields) error,
	qs []string,
	vof func(OrderFields) error,
) (*QueryFilter, error) {

	filterFields := FilterFieldsFromMap(qf)
	orderFields, err := OrderFieldsFromSlice(qs)
	if err != nil {
		return nil, err
	}

	qp.Pagination = true

	return &QueryFilter{
		PaginationParam:      qp,
		FilterFields:         filterFields,
		validateFilterFields: vff,
		OrderFields:          orderFields,
		validateOrderFields:  vof,
	}, nil
}

func NewQueryMapFilter(qp PaginationQuery,
	filterFields *FilterFields,
	vff func(FilterFields) error,
	orderFields *OrderFields,
	vof func(OrderFields) error,
) (*QueryFilter, error) {

	qp.Pagination = true

	return &QueryFilter{
		PaginationParam:      qp,
		FilterFields:         filterFields,
		validateFilterFields: vff,
		OrderFields:          orderFields,
		validateOrderFields:  vof,
	}, nil
}

func (qf QueryFilter) ValidateFilterFields() error {
	if qf.FilterFields == nil {
		return nil
	}
	return qf.validateFilterFields(*qf.FilterFields)
}

func (qf QueryFilter) ValidateOrderFields() error {
	if qf.OrderFields == nil {
		return nil
	}
	return qf.validateOrderFields(*qf.OrderFields)
}

//-----------------------------------------

type FilterField struct {
	Key   string
	Value any
}
type FilterFields []FilterField

func FilterFieldsFromMap(m map[string]any) *FilterFields {

	if len(m) == 0 {
		return nil
	}

	filterFields := FilterFields{}

	for k, v := range m {
		filterFields = append(filterFields, FilterField{
			Key:   k,
			Value: v,
		})
	}

	return &filterFields
}

//========================================================
// Order

type OrderField struct {
	Key       string
	Direction OrderDirection
}
type OrderFields []OrderField

type OrderDirection string

const (
	OrderByASC  OrderDirection = "ASC"
	OrderByDESC OrderDirection = "DESC"
)

func StringTo_OrderDirection(s string) (OrderDirection, error) {
	switch s {
	case "ASC":
		return OrderByASC, nil
	case "DESC":
		return OrderByDESC, nil
	default:
		return "", domainerror.ErrBadSortDirectionField
	}
}

func ToString_OrderFields(ofs OrderFields) string {
	orders := make([]string, len(ofs))

	for i, of := range ofs {
		orders[i] = fmt.Sprintf("%s %s", of.Key, of.Direction)
	}

	return strings.Join(orders, ",")
}

func OrderFieldsFromStrings(s []string) (*OrderFields, error) {

	orderFields := OrderFields{}

	for _, v := range s {
		var orderDirection OrderDirection
		if strings.HasPrefix(v, "-") {
			orderDirection = "DESC"
			v = strings.TrimPrefix(v, "-")
		} else {
			orderDirection = "ASC"
		}
		orderField := NewOrderField(v, orderDirection)
		orderFields = append(orderFields, *orderField)
	}

	return &orderFields, nil
}

func OrderFieldsFromSlice(s []string) (*OrderFields, error) {

	if len(s) == 0 {
		return nil, nil
	}
	//we need the slice to be even, to have pairs of fields in form: ["country","ASC"]
	if len(s)%2 != 0 {
		return nil, domainerror.ErrSortFieldNotInPair
	}

	orderFields := OrderFields{}

	for i := 0; i < len(s)-1; i += 2 {
		// Access elements in pairs
		key := s[i]
		direction := s[i+1]

		orderDirection, err := StringTo_OrderDirection(direction)
		if err != nil {
			return nil, err
		}

		orderField := NewOrderField(key, orderDirection)
		orderFields = append(orderFields, *orderField)

	}
	return &orderFields, nil
}

func NewOrderFields(orderFields ...*OrderField) []*OrderField {
	return orderFields
}

func NewOrderField(key string, d OrderDirection) *OrderField {
	return &OrderField{
		Key:       key,
		Direction: d,
	}
}
