package pport

import (
	"context"
)

// AdminRegion ...
type AdminRegionApp interface {
	Add(ctx context.Context)
	Update()
	//TODO: this implements filters
	Query()
}
