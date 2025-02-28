package sport

import (
	"context"

	"github.com/andriykusevol/aktemplategorm/internal/domain/aggregate/feeset"

	"github.com/andriykusevol/aktemplategorm/internal/domain/entity/feerange"
	"github.com/andriykusevol/aktemplategorm/internal/domain/entity/psp"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/patch"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/queryfilter"
)

// The application should use this repository to build FeeSet
// then it can call pure domain functions like IdentifyFeeRange(money money.Money) and others.
// This will enshure integrity of our aggregate root (it is an aggregate root for the "psp fees" context).
type FeeSetRepository interface {
	AddPSP(ctx context.Context, psp psp.PSP) (*psp.PSP, error)

	QueryFilterPSP(ctx context.Context, queryFilter queryfilter.QueryFilter) ([]psp.PSP, *queryfilter.Pagination, error)
	QueryPSP(ctx context.Context, domainQuery psp.Query) ([]psp.PSP, *queryfilter.Pagination, error)

	PspUpdateByID(ctx context.Context, id uint, p patch.Patch) (*psp.PSP, error)

	GetPspByID(ctx context.Context, psp_id uint) (*psp.PSP, error)

	GetPspByCode(ctx context.Context, psp_code string) (*psp.PSP, error)

	ListPSP(ctx context.Context) ([]psp.PSP, error)

	DeletePSP(ctx context.Context, psp_id uint) error

	// Create new Fee Set. A set once created cannot be modified! Any modification will create a new set.
	// The FeeSet knows where to add it.
	Add(ctx context.Context, feeSet feeset.FeeSet) error

	// List a Fee Sets for specific PSP.
	// We should pupulate each FeeSet with FeeRange list
	List(ctx context.Context, pspCode string) ([]feeset.FeeSet, error)

	// Get a specific Fee Set.
	Get(ctx context.Context, pspCode string, pspFeeSetID uint) (*feeset.FeeSet, error)
	IsEmpty(ctx context.Context, pspID uint) (bool, error)

	GetActive(ctx context.Context, PspID uint) (*feeset.FeeSet, error)

	// Delete FeeRange
	UpdateStatus(ctx context.Context, pspFeeSetID uint, status bool) error

	// Add Fee Range for specifict Fee Set.
	// The FeeRange knows where to add it.
	AddFeeRange(ctx context.Context, feeRange feerange.FeeRange) error

	// List Fee Ranges for specific Fee Set
	ListFeeRange(ctx context.Context, pspFeeSetID uint) ([]feerange.FeeRange, error)

	//Get a Fee Range by id
	GetFeeRange(ctx context.Context, feerange_id uint) (*feerange.FeeRange, error)

	//==========================================================
	// We do not define functional API here, because it is not a task of this secondary port.
	// The Functional API shild be defined in a primary port and implemented in the application layer.

}
