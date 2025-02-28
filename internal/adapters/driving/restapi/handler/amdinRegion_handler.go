// Package handler to handle http requests.
package handler

import (
	"github.com/andriykusevol/aktemplategorm/internal/application/pport"
	"github.com/gin-gonic/gin"
)

// AdminRegion interface is not a port, It's just an interface for router.
type AdminRegionHandler interface {
	AddRegion(c *gin.Context)
	ListRegion(c *gin.Context)
	UpdateRegion(c *gin.Context)
	DeleteRegion(c *gin.Context)
	AddCity(c *gin.Context)
	DeleteCity(c *gin.Context)
	ListRegionInCountry(c *gin.Context)
	ListCityInCountryByRegion(c *gin.Context)
	ListCityInRegion(c *gin.Context)
	GetCityInRegion(c *gin.Context)
}

type appsAdminRegionHandler struct {
	adminRegion pport.AdminRegionApp
}

func NewAdminRegionHandler(adminRegionApp pport.AdminRegionApp) AdminRegionHandler {
	return &appsAdminRegionHandler{adminRegion: adminRegionApp}
}

//---------------------------------------------------------------
// Management API
//---------------------------------------------------------------

func (apps *appsAdminRegionHandler) AddRegion(c *gin.Context) {
	// Add a new Region
	ctx := c.Request.Context()
	apps.adminRegion.Add(ctx)

}

func (apps *appsAdminRegionHandler) ListRegion(c *gin.Context) {
	// List all regions

}

func (apps *appsAdminRegionHandler) UpdateRegion(c *gin.Context) {
	// Update an existing Region

}

func (apps *appsAdminRegionHandler) DeleteRegion(c *gin.Context) {
	// Soft Delete Region (CASCADE)

}

func (apps *appsAdminRegionHandler) AddCity(c *gin.Context) {
	// Add a new City to Region

}

func (apps *appsAdminRegionHandler) DeleteCity(c *gin.Context) {
	// Delete City by it's ID.

}

//---------------------------------------------------------------
// Functional API
//---------------------------------------------------------------

func (apps *appsAdminRegionHandler) ListRegionInCountry(c *gin.Context) {
	// List regions in a country (country code ISO 3166-1 ALPHA-2)
	// Allows the requestor to get a list of all regions for a given country. Sort Regions alphabetically by Region_name_EN

}

func (apps *appsAdminRegionHandler) ListCityInCountryByRegion(c *gin.Context) {
	// List Cities in a Country grouped by Region. Sort Cities alphabetically
	// Provide list cities in a country grouped by region. Sort Cities alphabetically

}

func (apps *appsAdminRegionHandler) ListCityInRegion(c *gin.Context) {
	// Provide list of cities in a region. Sort Cities alphabetically

}

func (apps *appsAdminRegionHandler) GetCityInRegion(c *gin.Context) {
	// Get City in Region by ID

}
