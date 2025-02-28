// Package application ...
package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/andriykusevol/aktemplategorm/internal/application/pport"
	"github.com/andriykusevol/aktemplategorm/internal/domain/aggregate/feeset"
	"github.com/andriykusevol/aktemplategorm/internal/domain/entity/feerange"
	"github.com/andriykusevol/aktemplategorm/internal/domain/entity/psp"
	"github.com/andriykusevol/aktemplategorm/internal/domain/sport"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/patch"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/queryfilter"

	//TODO: Bad link for quick bulk fees calculation
	// And now our application is tightly coupled with the presentation layer (rest api)
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driving/restapi/request"
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driving/restapi/response"

	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/currency"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/money"

	"github.com/govalues/decimal"
)

type feeSetApp struct {
	feeSetRepo sport.FeeSetRepository
}

func NewFeeSetApp(feeSetRepo sport.FeeSetRepository) pport.FeeSetApp {
	fmt.Println("")
	return &feeSetApp{
		feeSetRepo: feeSetRepo,
	}
}

func (app *feeSetApp) PspUpdateByID(ctx context.Context, id uint, p patch.Patch) (*psp.PSP, error) {

	domainItem, err := app.feeSetRepo.PspUpdateByID(ctx, id, p)
	if err != nil {
		return nil, err
	}
	return domainItem, nil
}

func (app *feeSetApp) PspAdd(ctx context.Context, psp *psp.PSP) (*psp.PSP, error) {
	pspvalue := *psp
	domainItem, err := app.feeSetRepo.AddPSP(ctx, pspvalue)
	if err != nil {
		return nil, err
	}
	return domainItem, nil
}

func (app *feeSetApp) GetPSP(ctx context.Context, id uint) (*psp.PSP, error) {

	domainEntity, _ := app.feeSetRepo.GetPspByID(ctx, id)
	return domainEntity, nil
}

func (app *feeSetApp) QueryFilterPSP(ctx context.Context, qf queryfilter.QueryFilter) ([]psp.PSP, *queryfilter.Pagination, error) {

	domainItems, pagination, err := app.feeSetRepo.QueryFilterPSP(ctx, qf)
	if err != nil {
		return nil, nil, err
	}

	return domainItems, pagination, nil
}

func (app *feeSetApp) QueryPSP(ctx context.Context, domainQuery psp.Query) ([]psp.PSP, *queryfilter.Pagination, error) {

	domainItems, pagination, err := app.feeSetRepo.QueryPSP(ctx, domainQuery)
	if err != nil {
		return nil, nil, err
	}

	return domainItems, pagination, nil
}

func (app *feeSetApp) ListPSP(ctx context.Context, psp_id *uint, psp_code *string) ([]psp.PSP, error) {
	//TODO: Handle error.

	var psps []psp.PSP

	if psp_id != nil {
		psp, _ := app.feeSetRepo.GetPspByID(ctx, *psp_id)
		if psp == nil {
			return psps, nil
		} else {
			psps = append(psps, *psp)
			return psps, nil
		}
	}

	if psp_code != nil {
		psp, _ := app.feeSetRepo.GetPspByCode(ctx, *psp_code)
		if psp == nil {
			return psps, nil
		} else {
			psps = append(psps, *psp)
			return psps, nil
		}
	}

	psps, _ = app.feeSetRepo.ListPSP(ctx)
	return psps, nil
}

func (app *feeSetApp) DeletePSP(ctx context.Context, psp_id uint) error {

	err := app.feeSetRepo.DeletePSP(ctx, psp_id)
	if err != nil {
		return err
	}

	return nil

}

func (app *feeSetApp) AddMobileProviderFeeSet(ctx context.Context, feeSet *feeset.FeeSet) error {

	isFeeSetEmpty, err := app.feeSetRepo.IsEmpty(ctx, feeSet.PspID)
	if err != nil {
		return err
	}

	if !isFeeSetEmpty {
		// Get active set
		activeFeeSet, err := app.feeSetRepo.GetActive(ctx, feeSet.PspID)
		if err != nil {
			return err
		}

		if activeFeeSet != nil {
			// Update status of that set
			app.feeSetRepo.UpdateStatus(ctx, *activeFeeSet.ID, false)
			if err != nil {
				return err
			}
		}
	}

	//TODO: Handle error.
	err = app.feeSetRepo.Add(ctx, *feeSet)
	if err != nil {
		return err
	}

	return nil

}

func (app *feeSetApp) AddMobileProviderFeeRange(ctx context.Context, feeRange feerange.FeeRange) error {

	//We pass the feeSet here to make shure of integrity of our aggregate (feeSet)
	err := app.feeSetRepo.AddFeeRange(ctx, feeRange)
	if err != nil {
		return err
	}

	return nil

}

func (app *feeSetApp) ListFeeRange(ctx context.Context, pspFeeSetID uint) ([]*feerange.FeeRange, error) {

	feeRanges, err := app.feeSetRepo.ListFeeRange(ctx, pspFeeSetID)
	if err != nil {
		return nil, err
	}

	for _, feeRange := range feeRanges {
		_ = feeRange
		//TODO: Construct: []*feerange.FeeRange response
	}

	//TODO: not implemented
	return nil, nil
}

func (app *feeSetApp) GetFeeRange(ctx context.Context, feerange_id uint) (*feerange.FeeRange, error) {

	feeRange, err := app.feeSetRepo.GetFeeRange(ctx, feerange_id)
	if err != nil {
		return nil, err
	}

	return feeRange, nil
}

// TODO: This should be in the domain logic
// Inclusive Range: min <= value <= max
func isInRange(value, min, max decimal.Decimal) bool {
	// Check if value is between min and max (inclusive)
	result := value.Cmp(min) >= 0 && value.Cmp(max) <= 0
	return result
}

// Exclusive Range: min < value < max
func isInExclusiveRange(value, min, max decimal.Decimal) bool {
	// Check if value is strictly between min and max (exclusive)
	return value.Less(max) && !value.Less(min)
}

func (app *feeSetApp) CalculateFeeForAmount(ctx context.Context, pspCode string, amount money.Money) (totalamount, fees float64, err error) {

	//TODO: Implement supporting of single Active PSPFeeSet for PSP privider.
	// It implies that, if the PSP supports multiple currencies, and have only single active FeeSet,
	// the FeeRanges have to contain currencies, but the FeeSet is not.
	//pspFeeSetID := uint(1)

	psp, err := app.feeSetRepo.GetPspByCode(ctx, pspCode)
	if err != nil {
		return 0, 0, err
	}

	activeFeeSet, err := app.feeSetRepo.GetActive(ctx, *psp.ID)
	if err != nil {
		return 0, 0, err
	}

	//feeRanges := a.ListFeeRange(ctx, pspFeeSetID)
	feeRanges, err := app.feeSetRepo.ListFeeRange(ctx, *activeFeeSet.ID)
	if err != nil {
		return 0, 0, err
	}

	if len(feeRanges) == 0 {
		return 0, 0, errors.New("Amount is not covered by any fee range")
	}

	//TODO: This should be placed to the domain logic.
	for _, feeRange := range feeRanges {
		//TODO: Handle an error.

		if isInRange(amount.Amount(), feeRange.From, *feeRange.To) {
			amount_with_fees, _ := amount.Amount().Add(feeRange.FeeFixed)
			totalamount, _ := amount_with_fees.Float64()
			fees, _ := feeRange.FeeFixed.Float64()
			return totalamount, fees, nil
		}
	}

	//TODO!: What if amount is not covered by any range?
	// Now we're just return the def_max_total_fee.
	fees, ok := feeRanges[0].MaxTotalFee.Float64()
	if !ok {
		return 0, 0, errors.New("Cannot convert Decimal to Float64")
	}

	totalamount, _ = amount.Amount().Float64()
	return totalamount, fees, nil
}

// TODO: We're skipping the domain things for this right now
// and use request and responses from the presentation layer (bad link)
func (app *feeSetApp) BulkCalculateFeeForAmount(ctx context.Context, req request.BulkFeeRequest) (response.BulkFeeResponse, error) {

	bulkResponse := response.BulkFeeResponse{
		Results: []response.FeeReponse{}, // Initialize with an empty slice
	}

	for _, requestItem := range req {

		//TODO: and you see the consiquences - we have to build domain object in the aplication
		cur := currency.NewCurrency("XAF")

		amoundDecimal, err := decimal.Parse(requestItem.Amount)
		if err != nil {
			return bulkResponse, err
		}

		m := money.NewMoney(amoundDecimal, cur)

		totalamount, fees, err := app.CalculateFeeForAmount(ctx, requestItem.PspCode, m)
		if err != nil {
			return bulkResponse, err
		}

		amount, ok := amoundDecimal.Float64()
		if !ok {
			return bulkResponse, errors.New("Cannot convert Decimal to Float64")
		}

		resp := response.FeeReponse{
			Amount:      amount,
			FeeAmount:   fees,
			TotalAmount: totalamount,
			PspCode:     requestItem.PspCode,
		}

		bulkResponse.Results = append(bulkResponse.Results, resp)

	}

	return bulkResponse, nil
}

func (app *feeSetApp) UpdateFeeRange(ctx context.Context, p patch.Patch) {

	// We can also validate a state of the patch from any place.
	//p.ValidatePatch()

	// for key, value := range *p.Data() {
	// 	//fmt.Printf("Key: %s, Value: %v, Type: %s\n", key, value, reflect.TypeOf(value))
	// 	_ = value

	// 	// !!! We cannot catch if the field was ommited here.
	// 	// Exists, Present := (*p.Data())[key]				// The Present will be always true.

	// 	// That is why we check required fields in _ValidatePatch function (in the domain)
	// 	// Remember, that for patch we do not have an annotations (we do not use request stucture at all)
	// 	// If the field in not in there, we do nothing with it.

	// 	// Handle fields that was explicitly set to null from the json payload.
	// 	Exists, _ := (*p.Data())[key]
	// 	if Exists == nil {
	// 		fmt.Println(key, ": field was explicitly set to null")
	// 	} else {
	// 		fmt.Println(key, ": field has a valid value")
	// 	}
	// }
}
