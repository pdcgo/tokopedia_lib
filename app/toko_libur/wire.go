//go:build wireinject
// +build wireinject

package toko_libur

import (
	"github.com/google/wire"
	"github.com/pdcgo/common_conf/pdc_application"
)

func InitApplication(base pdc_application.BaseApplication) (*Application, error) {

	wire.Build(
		NewApplicationConfig,
		NewReport,
		NewApplication,
	)

	return &Application{}, nil
}
