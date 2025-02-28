package feeset

import (
	"errors"
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driven/gorm"
	domain_psp "github.com/andriykusevol/aktemplategorm/internal/domain/entity/psp"
	domain_a "github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/auditdbrecord"

	"github.com/google/uuid"
)

type PspDto struct {
	ID             *uint              `gorm:"primaryKey;autoIncrement"`
	PspCode        string             `gorm:"size:16"`
	PspShortName   *string            `gorm:"size:50"`
	PspCountryCode *string            `gorm:"size:3"`
	AuditRecord    orm.AuditRecordDto `gorm:"embedded"`
}

func (PspDto) TableName() string {
	return "Psp"
}

func DomainToDto_PSP(psp domain_psp.PSP) PspDto {

	auditRecordDTO := orm.DomainToDto_AuditDbRecord(psp.AuditDBRecord)

	dto := PspDto{
		PspCode:        psp.PspCode,
		PspCountryCode: psp.PspCountryCode,
		PspShortName:   psp.PspShortName,
		AuditRecord:    *auditRecordDTO,
	}

	return dto
}

func DtoToDomain_PSP(dto PspDto) (*domain_psp.PSP, error) {

	if dto.ID == nil {
		return nil, errors.New("PSP not found")
	}

	uid, _ := uuid.FromBytes(dto.AuditRecord.CreatedBy)

	a := domain_a.NewAuditDBRecord(uid)
	a.SetCreatedAt(dto.AuditRecord.CreatedAt)
	updatedbyuuid, _ := uuid.FromBytes(dto.AuditRecord.UpdatedBy)
	a.SetUpdatedAt(dto.AuditRecord.UpdatedAt, updatedbyuuid)

	if !dto.AuditRecord.DeletedAt.Time.IsZero() {
		a.SetDeletedAt(dto.AuditRecord.DeletedAt.Time, uuid.UUID(dto.AuditRecord.DeletedBy))
	}

	domainItem := domain_psp.New(dto.PspCode, a)
	domainItem.ID = dto.ID

	if dto.PspCountryCode != nil {
		domainItem.PspCountryCode = dto.PspCountryCode
	}

	if dto.PspShortName != nil {
		domainItem.PspShortName = dto.PspShortName
	}

	return &domainItem, nil
}

func ListDtoToDomain_PSP(dtos []PspDto) []domain_psp.PSP {

	psps := []domain_psp.PSP{}
	for _, dto := range dtos {
		domain, err := DtoToDomain_PSP(dto)
		if err != nil {
			continue
		}
		psps = append(psps, *domain)
	}
	return psps
}
