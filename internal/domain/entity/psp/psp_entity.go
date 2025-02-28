// Package psp ...
package psp

import (
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/auditdbrecord"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/queryfilter"
)

// Making ID as uuid simplifies the whole approach.
// We generate ID on the backend, not in the database,
// Therefore, the ID field can already be made required (not a pointer)
// And if we generate it in the database upon creation, then this field should be
// optional, because upon creation, we do not yet know the ID.

type PSP struct {
	ID             *uint
	PspCode        string
	PspCountryCode *string
	PspShortName   *string
	AuditDBRecord  auditdbrecord.AuditDBRecord
}

type PSPs []*PSP

type Query struct {
	PaginationParam queryfilter.PaginationQuery
	OrderFields     *queryfilter.OrderFields
	IDs             *[]string
	PspCode         *string
	PspCountryCode  *string
}

func NewQuery(pq queryfilter.PaginationQuery) *Query {
	pq.Pagination = true // We always use pagination
	qp := &Query{
		PaginationParam: pq,
	}
	return qp
}

func New(
	pspCode string,
	auditDbRecord auditdbrecord.AuditDBRecord) PSP {

	return PSP{
		PspCode:       pspCode,
		AuditDBRecord: auditDbRecord,
	}
}
