package feeset

import (
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driven/gorm"
	domain_feeset "github.com/andriykusevol/aktemplategorm/internal/domain/aggregate/feeset"

	//domain_psp "github.com/andriykusevol/aktemplategorm/internal/domain/entity/psp"
	domain_a "github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/auditdbrecord"
	domain_c "github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/currency"

	"github.com/google/uuid"
)

// "FeeSet" bolengs to "PSP", "PspID" is a foreign key.

// By default, the PspID is implicitly used to create a foreign key relationship between the FeeSet and Psp tables,
// and thus must be included in the Feeset struct in order to fill the Psp inner struct.

// To define a belongs to relationship, the foreign key must exist,
// the default foreign key uses the owner’s type name plus its primary field name.

// To define the FeeSet that belongs to Psp, the foreign key should be PspID by convention

type Dto struct {
	ID          uint               `gorm:"primaryKey;autoIncrement"`
	PspID       uint               `gorm:"foreignKey:PspID"`
	AuditRecord orm.AuditRecordDto `gorm:"embedded"`
}

func (Dto) TableName() string {
	return "FeeSet"
}

func DomainToDto_FeeSet(feeset domain_feeset.FeeSet) *Dto {

	// Here we must be sure that our domain object is already self-consistent.
	// Therefore, we do not do any checks.
	auditRecordDTO := orm.DomainToDto_AuditDbRecord(feeset.AuditDBRecord)

	dto := &Dto{
		PspID:       feeset.PspID,
		AuditRecord: *auditRecordDTO,
	}

	return dto

}

func DtoToDomain_FeeSet(dto Dto) *domain_feeset.FeeSet {

	uuid, _ := uuid.FromBytes(dto.AuditRecord.CreatedBy)
	a := domain_a.NewAuditDBRecord(uuid)

	c := domain_c.NewCurrency("XAF")

	domainItem := domain_feeset.NewFeeSet(c, dto.PspID, a)
	domainItem.ID = &dto.ID

	return domainItem

}
