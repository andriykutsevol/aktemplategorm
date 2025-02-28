package feerange

import (
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/auditdbrecord"
	"github.com/govalues/decimal"
)

type FeeRange struct {
	ID            *uint
	FeeSetID      uint             // Link to fee set this range is part of
	From          decimal.Decimal  // Lower bound this fee is configured to
	To            *decimal.Decimal // Upper bound this fee is configure to
	FeeFixed      decimal.Decimal  // Fixed fee amount
	FeePercentage *decimal.Decimal // Fixed fee amount
	MaxTotalFee   decimal.Decimal  // Set by the system on record creation unless updated by user.
	AuditDBRecord auditdbrecord.AuditDBRecord
}

func NewFeeRange(feesetID uint, auditDbRecord auditdbrecord.AuditDBRecord) *FeeRange {

	//Initialize default values
	// It is for the case when we do not set any values explicitly.
	def_from, _ := decimal.NewFromFloat64(0.00)
	def_feeFixed, _ := decimal.NewFromFloat64(0.00)
	def_feePercentage, _ := decimal.NewFromFloat64(0.00)
	def_max_total_fee, _ := decimal.NewFromFloat64(999999999000.00)

	return &FeeRange{
		FeeSetID:      feesetID,
		From:          def_from,
		FeeFixed:      def_feeFixed,
		FeePercentage: &def_feePercentage,
		MaxTotalFee:   def_max_total_fee,
		AuditDBRecord: auditDbRecord,
	}

}
