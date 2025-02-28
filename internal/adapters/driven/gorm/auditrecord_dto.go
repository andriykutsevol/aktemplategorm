package orm

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/auditdbrecord"
)

// Set of common fields shared between various Data Transfer Objects
type AuditRecordDto struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	CreatedBy []byte `gorm:"type:binary(16)"`
	UpdatedBy []byte `gorm:"type:binary(16)"`
	DeletedBy []byte `gorm:"type:binary(16)"` //Empty []byte represents null (no value)
}

func uuidToBytes(id uuid.UUID) []byte {
	return id[:]
}

func DomainToDto_AuditDbRecord(auditRecord auditdbrecord.AuditDBRecord) *AuditRecordDto {

	auditRecordDto := AuditRecordDto{}
	auditRecordDto.CreatedAt = auditRecord.CreatedAt()
	auditRecordDto.UpdatedAt = auditRecord.UpdatedAt()
	auditRecordDto.SetDeletedAtTime(auditRecord.DeletedAt())
	auditRecordDto.CreatedBy = uuidToBytes(auditRecord.CreatedBy())
	auditRecordDto.UpdatedBy = uuidToBytes(auditRecord.UpdatedBy())

	if auditRecord.DeletedBy() != nil {
		auditRecordDto.DeletedBy = uuidToBytes(*auditRecord.DeletedBy())
	}

	return &auditRecordDto
}

// use this to get deletedAt value
func (d AuditRecordDto) DeletedAtTime() *time.Time {
	if d.DeletedAt.Valid {
		return &d.DeletedAt.Time
	}
	return nil
}

// setters are discouraged in go but in this case we seem to need it
func (d *AuditRecordDto) SetDeletedAtTime(t *time.Time) {
	if t != nil {
		d.DeletedAt.Scan(t)
	}
}
