package response

import (
	domain "github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/auditdbrecord"
	"time"

	"github.com/google/uuid"
)

type AuditDBRecord struct {
	CreatedAt time.Time  `json:"createdat" binding:"required"`
	CreatedBy uuid.UUID  `json:"createdby" binding:"required"`
	UpdatedAt time.Time  `json:"updatedat" binding:"required"`
	UpdatedBy uuid.UUID  `json:"updatedby" binding:"required"`
	DeletedAt *time.Time `json:"deletedat" binding:"required"`
	DeletedBy *uuid.UUID `json:"deletedby" binding:"required"`
}

func FromDomain_AuditDBRecord(a domain.AuditDBRecord) AuditDBRecord {
	response := AuditDBRecord{}
	response.CreatedAt = a.CreatedAt()
	response.CreatedBy = a.CreatedBy()
	response.UpdatedAt = a.UpdatedAt()
	response.UpdatedBy = a.UpdatedBy()
	response.DeletedAt = a.DeletedAt()
	response.DeletedBy = a.DeletedBy()
	return response
}
