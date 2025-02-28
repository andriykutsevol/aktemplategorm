package feeset

import (
	"context"
	"errors"
	"fmt"
	"github.com/andriykusevol/aktemplategorm/internal/domain/aggregate/feeset"
	"github.com/andriykusevol/aktemplategorm/internal/domain/sport"

	"github.com/google/uuid"

	//auth_domain "github.com/andriykusevol/aktemplategorm/internal/domain/entity/auth"
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driven/gorm"
	"github.com/andriykusevol/aktemplategorm/internal/domain/entity/feerange"
	"github.com/andriykusevol/aktemplategorm/internal/domain/entity/psp"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/patch"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/queryfilter"

	"gorm.io/gorm"
	// "fmt"
)

type repository struct {
	db *gorm.DB
}

// NewRepository ...
func NewRepository(db *gorm.DB) sport.FeeSetRepository {
	return &repository{db: db}
}

//----------------------GORM Query Context-----------------
// In GORM, query context refers to how the context of a *gorm.DB object
// is modified as you chain methods like Where, Joins, or Preload.

// gorm Query Context:
// Original object
// db := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// // Query context modification
// db1 := db.Where("status = ?", "active")
// db2 := db1.Where("role = ?", "admin")
// The original `db` is not affected

// In this example:
// db is untouched.
// db1 includes the condition status = 'active'.
// db2 includes both status = 'active' and role = 'admin'.
//----------------------------------------------------------

func uuidToBytes(id uuid.UUID) []byte {
	return id[:]
}

func (r *repository) AddPSP(ctx context.Context, psp psp.PSP) (*psp.PSP, error) {
	dto := DomainToDto_PSP(psp)

	err := r.db.Create(&dto).Error
	if err != nil {
		return nil, err
	}

	domainItem, err := DtoToDomain_PSP(dto)
	if err != nil {
		return nil, err
	}

	return domainItem, nil
}

// In general, it is not good idea to allow update the DeletedAt field from PATCH(update)
// So we lose the whole intention of gorm's DeletedAt handling.
// But we still have to use PATCH to restore that soft deleted rectord:
// PATCH /resource/:id/restore

func (r *repository) PspUpdateByID(ctx context.Context, id uint, p patch.Patch) (*psp.PSP, error) {

	var db *gorm.DB = r.db

	// Update attributes with `map`
	// db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	// UPDATE users SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;
	dto := PspDto{ID: &id}
	fmt.Println("---")
	fmt.Println(p.Data())
	err := db.Model(&dto).Unscoped().Updates(p.Data()).Error
	if err != nil {
		return nil, err
	}

	// TODO: Why it does not work with this:
	//err = db.Unscoped().Where(&dto).Find(&dto).Error

	err = db.Unscoped().Where("ID=(?)", id).Find(&dto).Error
	if err != nil {
		return nil, err
	}

	domainItem, err := DtoToDomain_PSP(dto)
	if err != nil {
		return nil, err
	}

	return domainItem, nil

}

func RetrievePSP(ctx context.Context) (*psp.PSP, error) {

	return nil, nil
}

func (r *repository) ListPSP(ctx context.Context) ([]psp.PSP, error) {

	var dtos []PspDto
	err := r.db.Find(&dtos).Error
	if err != nil {
		return nil, err
	}

	psps := ListDtoToDomain_PSP(dtos)

	return psps, nil
}

func (r *repository) DeletePSP(ctx context.Context, psp_id uint) error {

	var dto PspDto
	dto.ID = &psp_id

	err := r.db.Delete(&dto).Error
	if err != nil {
		return err
	}

	return nil
}

//========================================================

//========================================================

func (r *repository) QueryFilterPSP(ctx context.Context, qf queryfilter.QueryFilter) ([]psp.PSP, *queryfilter.Pagination, error) {

	var db *gorm.DB = r.db

	// type User struct {
	// 	Name string
	// 	Age  int
	// }
	// var users []User
	// db.Where(&User{Name: "John", Age: 30}).Find(&users)

	if qf.FilterFields != nil {

		for _, f := range *qf.FilterFields {

			if f.Key == "IDs" {
				db = db.Where("id IN (?)", f.Value)
				continue
			}

			db = db.Where(f.Key+"=?", f.Value)
		}
	}

	if qf.OrderFields != nil {
		for _, f := range *qf.OrderFields {
			s := f.Key + " " + string(f.Direction)
			db = db.Order(s)
		}
	}

	var dtos []PspDto
	db = db.Model(&PspDto{})
	pagination, err := orm.WrapPageQuery(ctx, db, qf.PaginationParam, &dtos)
	if err != nil {
		return nil, nil, err
	}
	psps := ListDtoToDomain_PSP(dtos)
	return psps, pagination, nil
}

func (r *repository) QueryPSP(ctx context.Context, query psp.Query) ([]psp.PSP, *queryfilter.Pagination, error) {

	var db *gorm.DB = r.db

	if query.OrderFields != nil {
		for _, f := range *query.OrderFields {
			s := f.Key + " " + string(f.Direction)
			db = db.Order(s)
		}
	}

	if query.IDs != nil {
		db = db.Where("id IN (?)", *query.IDs)
	}

	if query.PspCode != nil {
		db = db.Where("PspCode=(?)", *query.PspCode)
	}

	if query.PspCountryCode != nil {
		db = db.Where("PspCountryCode=(?)", *query.PspCountryCode)
	}

	// var dtos []PspDto
	// err := r.db.Scopes(orm.Paginate(qp.PaginationParam)).Find(&dtos).Error
	// if err != nil {
	// 	return nil, nil, err
	// }
	// psps := ListDtoToDomain_PSP(dtos)
	var dtos []PspDto
	db = db.Model(&PspDto{})
	pagination, err := orm.WrapPageQuery(ctx, db, query.PaginationParam, &dtos)
	if err != nil {
		return nil, nil, err
	}
	psps := ListDtoToDomain_PSP(dtos)

	return psps, pagination, nil
}

func (r *repository) GetPspByCode(ctx context.Context, psp_code string) (*psp.PSP, error) {

	var db *gorm.DB = r.db

	var dto PspDto
	err := db.Where(&PspDto{PspCode: psp_code}).Find(&dto).Error
	if err != nil {
		return nil, err
	}
	domainItem, err := DtoToDomain_PSP(dto)
	if err != nil {
		return nil, err
	}
	return domainItem, nil
}

func (r *repository) GetPspByID(ctx context.Context, psp_id uint) (*psp.PSP, error) {

	var db *gorm.DB = r.db

	var dto PspDto
	err := db.Where(&PspDto{ID: &psp_id}).First(&dto).Error
	if err != nil {
		return nil, err
	}

	domainItem, err := DtoToDomain_PSP(dto)
	if err != nil {
		return nil, err
	}
	return domainItem, nil

}

//===========================================================================================

func (r *repository) Add(ctx context.Context, feeSet feeset.FeeSet) error {
	// Domain interface implementation
	var db *gorm.DB = r.db

	dto := DomainToDto_FeeSet(feeSet)

	err := db.Create(&dto).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) List(ctx context.Context, pspCode string) ([]feeset.FeeSet, error) {

	return nil, nil
}

// TODO: define pspCode, or refactor to remove it.
// But checking for the pspCode could help to keep integrity.
func (r *repository) Get(ctx context.Context, pspCode string, pspFeeSetID uint) (*feeset.FeeSet, error) {

	return nil, nil
}

func (r *repository) IsEmpty(ctx context.Context, pspID uint) (bool, error) {
	var db *gorm.DB = r.db

	var dtos []Dto
	var count int64

	err := db.Model(&dtos).Where("PspID = ?", pspID).Count(&count).Error
	if err != nil {
		return true, err
	}

	if count == 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func (r *repository) GetActive(ctx context.Context, pspID uint) (*feeset.FeeSet, error) {

	var dtos []Dto

	err := r.db.Where("PspID = ? AND IsActive = ?", pspID, 1).Find(&dtos).Error
	if err != nil {
		return nil, err
	}

	if len(dtos) > 1 {
		return nil, errors.New("You have more than one active FeeSet for the PSP.")
	}

	if len(dtos) == 0 {
		return nil, errors.New("Psp Not Found OR Active FeeSet not found for this Psp.")
	}

	domainItem := DtoToDomain_FeeSet(dtos[0])
	return domainItem, nil
}

func (r *repository) UpdateStatus(ctx context.Context, pspFeeSetID uint, status bool) error {

	err := r.db.Model(&Dto{}).Where("ID = ?", pspFeeSetID).
		Update("IsActive", status).Error
	if err != nil {
		return err
	}

	return nil
}

// TODO: Refactor to receive FeeRange
// We pass the feeSet here to make shure of integrity of our aggregate (feeSet)
func (r *repository) AddFeeRange(ctx context.Context, feeRange feerange.FeeRange) error {
	// You have a propper feeSet here, with business invariants, etc...

	dto, err := DomainToDto_FeeRange(feeRange)
	if err != nil {
		return err
	}

	err = r.db.Create(&dto).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) ListFeeRange(ctx context.Context, pspFeeSetID uint) ([]feerange.FeeRange, error) {

	//TODO: We're still have a pspFeeSetID as a constant (uint 1)

	var dtos []FeeRangeDto
	err := r.db.Where("FeeSetID = ?", pspFeeSetID).Find(&dtos).Error
	if err != nil {
		return nil, err
	}

	feeranges := ListDtoToDomain_FeeRange(dtos)

	return feeranges, nil
}

func (r *repository) GetFeeRange(ctx context.Context, feerange_id uint) (*feerange.FeeRange, error) {

	var dto FeeRangeDto

	err := r.db.First(&dto, "ID = ?", feerange_id).Error
	if err != nil {
		return nil, err
	}

	domainItem := DtoToDomain_FeeRange(dto)

	return domainItem, nil
}
