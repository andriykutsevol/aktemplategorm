// Package application ...
package app

import (
	"context"
	"github.com/andriykusevol/aktemplategorm/internal/application/pport"
	"github.com/andriykusevol/aktemplategorm/internal/domain/sport"
	"github.com/andriykusevol/aktemplategorm/pkg/logger"
)

type adminRegionApp struct {
	adminregion sport.AdminRegionRepository
}

// NewAdminRegion ...
func NewAdminRegion(adminregion sport.AdminRegionRepository) pport.AdminRegionApp {
	return &adminRegionApp{
		adminregion: adminregion,
	}
}

func (a *adminRegionApp) Add(ctx context.Context) {
	l := logger.Logger()
	l.Info("Logging from the Application Layer")
	_ = ctx

}

func (a *adminRegionApp) Update() {

}

func (a *adminRegionApp) Query() {

}
