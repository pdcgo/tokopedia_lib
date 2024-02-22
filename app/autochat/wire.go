//go:build wireinject
// +build wireinject

package autochat

import (
	"github.com/google/wire"
	"github.com/pdcgo/common_conf/pdc_application"
)

func InitApplication(base pdc_application.BaseApplication) (*Application, error) {

	wire.Build(
		NewAutochatMessage,
		NewAutochatConfig,
		NewApplication,
	)

	return &Application{}, nil
}
