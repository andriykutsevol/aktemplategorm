package patch

import (
	"context"
	auth_domain "github.com/andriykusevol/aktemplategorm/internal/domain/entity/auth"

	"github.com/google/uuid"
)

type Patch struct {
	data     map[string]any
	validate func(map[string]any) bool
	toDomain func(*Patch) error
}

func NewPatch(
	validate func(map[string]any) bool,
	toDomain func(*Patch) error) Patch {
	return Patch{
		data:     make(map[string]any),
		validate: validate,
		toDomain: toDomain,
	}
}

func (p Patch) Data() *map[string]any {
	return &p.data
}

func uuidToBytes(id uuid.UUID) []byte {
	return id[:]
}

// https://gorm.io/docs/update.html
func (p Patch) ValidatePatch() bool {
	isValid := p.validate(p.data)
	return isValid
}

func (p *Patch) ToDomain(ctx context.Context) error {
	deletedat, exists := p.data["DeletedAt"]
	if (exists) && (deletedat != nil) {
		_, exists := p.data["DeletedBy"]
		if !exists {
			ctxUid := ctx.Value(auth_domain.UserIDCtx{})
			tmpUuid, _ := uuid.Parse(ctxUid.(string))
			p.data["DeletedBy"] = uuidToBytes(tmpUuid)
		}
	} else if (exists) && (deletedat == nil) {
		p.data["DeletedBy"] = nil
	}

	p.toDomain(p)
	return nil
}
