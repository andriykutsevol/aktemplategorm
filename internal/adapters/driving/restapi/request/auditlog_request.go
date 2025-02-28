package request

// PSPFeeSet struct for PSPFeeSet
// AuditDBRecord struct for AuditDBRecord
type AuditDBRecord struct {
	// It should be created by a trigger function in the database.
	CreatedAt string `json:"CreatedAt" binding:"required"`
	// It should be updated by a trigger function in the database.
	UpdatedAt *string `json:"UpdatedAt,omitempty"`
	// It should be updated by a trigger function in the database.
	DeletedAt *string `json:"DeletedAt,omitempty"`
	// It should be provided by a gin(golang) middleware.
	CreatedBy string `json:"CreatedBy" binding:"required"`
	// It should be provided by a gin(golang) middleware.
	UpdatedBy *string `json:"UpdatedBy,omitempty"`
	// It should be provided by a gin(golang) middleware.
	DeletedBy *string `json:"DeletedBy,omitempty"`
}
