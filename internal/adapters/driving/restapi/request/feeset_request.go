package request

import (
	"github.com/google/uuid"
	//"time"

	"github.com/andriykusevol/aktemplategorm/internal/domain/aggregate/feeset"
	domain_a "github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/auditdbrecord"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/currency"
)

// PSPFeeSet struct for PSPFeeSet
type PSPFeeSet struct {
	PspID int `json:"PspID" binding:"required"`
}

func (r *PSPFeeSet) ToDomain_PSPFeeSet(userID string) *feeset.FeeSet {
	//TODO: Validations

	//TODO: Handle error
	createdBy, _ := uuid.Parse(userID)
	a := domain_a.NewAuditDBRecord(createdBy)

	cur := currency.NewCurrency("XAF")
	domainItem := feeset.NewFeeSet(cur, uint(r.PspID), a)

	return domainItem
}
