package request

import (
	"github.com/andriykusevol/aktemplategorm/internal/domain/entity/feerange"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/auditdbrecord"

	"github.com/google/uuid"
	"github.com/govalues/decimal"
)

type PSPFeeRange struct {
	// Link to fee set this range is part of
	FeeSetID uint `json:"FeeSetID" binding:"required"`
	// lower bound this fee is configured to
	//TODO: get rid of those int32?
	From *float64 `json:"From" binding:"required"`

	// upper bound this fee is configured to
	To *float64 `json:"To,omitempty"`

	// Fixed fee amount
	FeeFixed float64 `json:"FeeFixed" binding:"required"`
	// Fixed fee amount
	FeePercentage *float64 `json:"FeePercentage" binding:"required"`

	// set by the system on record creation unless updated by user
	MaxTotalFee *uint64 `json:"MaxTotalFee,omitempty"`
}

// Here we construct just an entity. Later we'll use it in application layer.
func (s *PSPFeeRange) ToDomain_PSPFeeRange(userID string) *feerange.FeeRange {

	//TODO: handle error
	createdBy, _ := uuid.Parse(userID)

	a := auditdbrecord.NewAuditDBRecord(createdBy)

	domainItem := feerange.NewFeeRange(s.FeeSetID, a)

	//TODO: Should use domain logic here, to ensure the FeeRange object integrity?
	//TODO: Handle error
	from, _ := decimal.NewFromFloat64(*s.From)

	domainItem.From = from

	if s.To != nil {
		to, _ := decimal.NewFromFloat64(*s.To)
		domainItem.To = &to
	}

	//TODO: Handle errrors
	feefixed, _ := decimal.NewFromFloat64(s.FeeFixed)
	domainItem.FeeFixed = feefixed

	feepercentage, _ := decimal.NewFromFloat64(*s.FeePercentage)
	domainItem.FeePercentage = &feepercentage

	maxtotalfee, _ := decimal.New(int64(*s.MaxTotalFee), 0)
	domainItem.MaxTotalFee = maxtotalfee

	return domainItem

}
