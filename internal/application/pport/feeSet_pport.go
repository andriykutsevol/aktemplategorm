package pport

import (
	"context"
	"github.com/andriykusevol/aktemplategorm/internal/domain/aggregate/feeset"
	"github.com/andriykusevol/aktemplategorm/internal/domain/entity/feerange"
	"github.com/andriykusevol/aktemplategorm/internal/domain/entity/psp"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/money"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/patch"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/queryfilter"

	//TODO: Bad link for quick bulk fees calculation
	// And now our application is tightly coupled with the presentation layer (rest api)
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driving/restapi/request"
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driving/restapi/response"
)

// The handler recieves an application interface (driving port), but returns handler interface (which is visible in router)
// and the handler is a driving adapter in this case (because it implements application interface)
// But that driving port could also be implemented with GRPC (and that would be an another adapter)

// Application, itself, can take and return well defined domain objects already.

type FeeSetApp interface {
	PspAdd(ctx context.Context, psp *psp.PSP) (*psp.PSP, error)

	GetPSP(ctx context.Context, id uint) (*psp.PSP, error)
	QueryFilterPSP(ctx context.Context, qf queryfilter.QueryFilter) ([]psp.PSP, *queryfilter.Pagination, error)
	QueryPSP(ctx context.Context, domainQuery psp.Query) ([]psp.PSP, *queryfilter.Pagination, error)

	ListPSP(ctx context.Context, psp_id *uint, psp_code *string) ([]psp.PSP, error)
	DeletePSP(ctx context.Context, psp_id uint) error

	PspUpdateByID(ctx context.Context, id uint, p patch.Patch) (*psp.PSP, error)

	AddMobileProviderFeeSet(ctx context.Context, feesetAggregate *feeset.FeeSet) error

	AddMobileProviderFeeRange(ctx context.Context, feeRange feerange.FeeRange) error

	ListFeeRange(ctx context.Context, pspFeeSetID uint) ([]*feerange.FeeRange, error)
	GetFeeRange(ctx context.Context, feerange_id uint) (*feerange.FeeRange, error)

	CalculateFeeForAmount(ctx context.Context, pspCode string, amount money.Money) (totalamount, fees float64, err error)

	BulkCalculateFeeForAmount(ctx context.Context, req request.BulkFeeRequest) (response.BulkFeeResponse, error)

	UpdateFeeRange(ctx context.Context, p patch.Patch)
}
