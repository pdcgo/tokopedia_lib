//go:build wireinject
// +build wireinject

package toko_libur

import (
	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

func InitApplication(cCtx *cli.Context) (*Application, error) {

	wire.Build(
		NewApplicationConfig,
		NewReport,
		NewApplication,
	)

	return &Application{}, nil
}
