package feeset

import (
	"github.com/andriykusevol/aktemplategorm/internal/domain/entity/feerange"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/auditdbrecord"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/currency"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/money"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/patch"
)

type FeeSet struct {
	ID *uint
	// value objects are just columns in your table (they have no ID)
	Currency  currency.Currency
	FeeRanges *[]feerange.FeeRange //Optional Field.
	PspID     uint
	// value objects are just columns in your table (they have no ID)
	AuditDBRecord auditdbrecord.AuditDBRecord
}

func NewFeeSet(cur currency.Currency, pspID uint, auditDbRecord auditdbrecord.AuditDBRecord) *FeeSet {
	return &FeeSet{
		Currency:      cur,
		PspID:         pspID,
		AuditDBRecord: auditDbRecord,
	}
}

func (f *FeeSet) FindFeeRangeByID(id uint) *feerange.FeeRange {
	for _, feeRange := range *f.FeeRanges {
		if feeRange.ID == &id {
			return &feeRange
		}
	}
	return nil
}

// Maybe you'll want to check whether the feeRange is applicable to this particular feeSet.
// For that you can use this function (and implement checking :) - the integrity insurance of an aggregate.
func (f *FeeSet) AppendFeeRange(feerange feerange.FeeRange) error {
	*f.FeeRanges = append(*f.FeeRanges, feerange)
	return nil
}

// Identify the fee range into which a record fits.
func (f *FeeSet) IdentifyFeeRange(money money.Money) *feerange.FeeRange {
	for _, feeRange := range *f.FeeRanges {
		_ = feeRange
	}
	return nil
}

func ValidatePatch(data map[string]any) bool {

	for key, value := range data {
		//TODO: Manually check for required fields, and business logic.
		_ = key
		_ = value
	}

	return true
}

func ToDomain(p *patch.Patch) error {
	return nil
}
