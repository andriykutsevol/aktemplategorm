package auditdbrecord

import (
	"github.com/google/uuid"
	"time"
)

// AuditDBRecord structure corresponds to AuditDBRecord scheme in the API definition.
type AuditDBRecord struct {
	createdAt time.Time
	updatedAt time.Time
	createdBy uuid.UUID
	updatedBy uuid.UUID
	deletedAt *time.Time
	deletedBy *uuid.UUID
}

// In new record - created by, always equal to updated by.
// In thise case we use gorm, and it automatically set CreatedAt and UpdatedAt when create record.
// so we do not need to set it manually here.
func NewAuditDBRecord(createdBy uuid.UUID) AuditDBRecord {
	return AuditDBRecord{
		createdBy: createdBy,
		updatedBy: createdBy,
	}
}

func (a *AuditDBRecord) SetUpdatedAt(updatedAt time.Time, updatedBy uuid.UUID) {
	a.updatedAt = updatedAt
	a.updatedBy = updatedBy
}

func (a *AuditDBRecord) SetDeletedAt(deletedAt time.Time, deletedBy uuid.UUID) {
	a.deletedAt = &deletedAt
	a.deletedBy = &deletedBy
}

func (a *AuditDBRecord) SetCreatedAt(createdAt time.Time) {
	a.createdAt = createdAt
}
func (a *AuditDBRecord) CreatedAt() time.Time {
	return a.createdAt
}
func (a *AuditDBRecord) UpdatedAt() time.Time {
	return a.updatedAt
}
func (a *AuditDBRecord) DeletedAt() *time.Time {
	return a.deletedAt
}

func (a *AuditDBRecord) DeletedBy() *uuid.UUID {
	return a.deletedBy
}

// --------------------------------------------------------
// TODO: optimize size of UUID in the database.
func (a *AuditDBRecord) CreatedBy() uuid.UUID {
	return a.createdBy
}

func (a *AuditDBRecord) UpdatedBy() uuid.UUID {
	return a.updatedBy
}

//--------------------------------------------------------
