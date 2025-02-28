// Package administrativeregion ...
package administrativeregion

import (
	"github.com/andriykusevol/aktemplategorm/internal/domain/sport"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// NewRepository ...
func NewRepository(db *gorm.DB) sport.AdminRegionRepository {
	return &repository{db: db}
}

//---------------------------------
// Management API
//---------------------------------

// Work with the database.
func (r *repository) Add() {
	// Domain interface implementation

}

// Work with the database.
func (r *repository) Update() {
	// Domain interface implementation

}

//---------------------------------
// Functional API
//---------------------------------

// Work with the database.
func (r *repository) ListRetionsInCoutry() {
	// Domain interface implementation
}

// Work with the database.
func (r *repository) ListCitiesInRegion() {
	// Domain interface implementation

}

// Work with the database.
func (r *repository) ListRetionsAndCitiesInCoutry() {
	// Domain interface implementation

}
