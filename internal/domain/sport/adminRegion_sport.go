// Package administrativeregion ...
package sport

// Repository is a administrativeregion port
type AdminRegionRepository interface {

	//---------------------------------
	// Management API
	//---------------------------------

	// API to add new entries
	Add() //TODO: What is an "entiy" in this case?
	// API to update existing entries
	Update() //TODO: What is an "entiy" in this case?

	//---------------------------------
	// Functional API
	//---------------------------------

	//TODO: Filters, pagination

	// Provide list regions only in a country
	ListRetionsInCoutry()

	// Provide list of cities in a region
	ListCitiesInRegion()

	// Provide list regions and cities in a country
	ListRetionsAndCitiesInCoutry()

	//---------------------------------
	// For requests with filters we use "Retrieve" prefix.

}
