package response

import (
	domain "github.com/andriykusevol/aktemplategorm/internal/domain/entity/feerange"
)

// ResponseFeeRange
type FeeRangeReponse struct {
	ID            *uint   `json:"id"`
	From          float64 `json:"from"`
	To            float64 `json:"to"`
	FeeFixed      float64 `json:"fee_fixed"`
	FeePercentage float64 `json:"fee_percentage"`
	MaxTotalFee   int     `json:"max_total_fee"`
}

// ResponseFromDomain_FeeRange
func FromDomain_FeeRange(feeRange *domain.FeeRange) FeeRangeReponse {

	from, _ := feeRange.From.Float64()
	to, _ := feeRange.To.Float64()
	feefixed, _ := feeRange.FeeFixed.Float64()
	feepercentage, _ := feeRange.FeePercentage.Float64()
	maxtotalfee, _ := feeRange.MaxTotalFee.Float64()

	feeRangeResponse := FeeRangeReponse{
		ID:            feeRange.ID,
		From:          from,
		To:            to,
		FeeFixed:      feefixed,
		FeePercentage: feepercentage,
		MaxTotalFee:   int(maxtotalfee),
	}
	return feeRangeResponse
}
