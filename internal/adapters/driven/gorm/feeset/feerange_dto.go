package feeset

import (
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driven/gorm"
	domain_feerange "github.com/andriykusevol/aktemplategorm/internal/domain/entity/feerange"
	domain_a "github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/auditdbrecord"
	"github.com/govalues/decimal"

	"github.com/google/uuid"
)

type FeeRangeDto struct {
	ID            uint `gorm:"primaryKey;autoIncrement"`
	FeeSetID      uint `gorm:"foreignKey:FeeSetID"`
	From          decimal.Decimal
	To            *decimal.Decimal
	FeeFixed      decimal.Decimal
	FeePercentage *decimal.Decimal
	MaxTotalFee   decimal.Decimal

	AuditRecord orm.AuditRecordDto `gorm:"embedded"`
}

func (FeeRangeDto) TableName() string {
	return "FeeRange"
}

func DomainToDto_FeeRange(feerange domain_feerange.FeeRange) (*FeeRangeDto, error) {

	auditRecordDTO := orm.DomainToDto_AuditDbRecord(feerange.AuditDBRecord)

	dto := &FeeRangeDto{
		FeeSetID:      feerange.FeeSetID,
		From:          feerange.From,
		To:            feerange.To,
		FeeFixed:      feerange.FeeFixed,
		FeePercentage: feerange.FeePercentage,
		MaxTotalFee:   feerange.MaxTotalFee,
		AuditRecord:   *auditRecordDTO,
	}

	return dto, nil

}

func DtoToDomain_FeeRange(dto FeeRangeDto) *domain_feerange.FeeRange {

	//TODO: use domain logic here.
	uuid, _ := uuid.FromBytes(dto.AuditRecord.CreatedBy)

	a := domain_a.NewAuditDBRecord(uuid)
	domainItem := domain_feerange.NewFeeRange(dto.FeeSetID, a)

	domainItem.ID = &dto.ID
	domainItem.From = dto.From
	domainItem.To = dto.To
	domainItem.FeeFixed = dto.FeeFixed

	return domainItem
}

func ListDtoToDomain_FeeRange(dtos []FeeRangeDto) []domain_feerange.FeeRange {

	feeranges := []domain_feerange.FeeRange{}

	for _, dto := range dtos {
		domain := DtoToDomain_FeeRange(dto)
		feeranges = append(feeranges, *domain)
	}

	return feeranges
}
