package toko_libur

import (
	"github.com/pdcgo/common_conf/pdc_application"
	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/urfave/cli/v2"
)

type AppConfig struct {
	BaseConfig pdc_application.BaseApplication
	Libur      bool
	Report     string
	Limit      int
	CloseStart int
	CloseEnd   int
}

func (c *AppConfig) GetReportPath() string {
	return c.BaseConfig.Path(c.Report)
}

func (c *AppConfig) CreateCloseShopInput() *api.CloseShopScheduleInput {

	input := api.CloseShopScheduleInput{
		Action: 1,
	}
	if c.Libur {
		input.Action = 0
		input.CloseStart = c.CloseStart
		input.CloseEnd = c.CloseEnd
	}

	return &input
}

func NewApplicationConfig(cCtx *cli.Context) *AppConfig {
	return &AppConfig{
		BaseConfig: &legacy_source.BaseConfig{
			BaseData: cCtx.String("base"),
		},
		Libur:      cCtx.Bool("libur"),
		Report:     cCtx.String("report"),
		Limit:      cCtx.Int("limit"),
		CloseStart: cCtx.Int("closeStart"),
		CloseEnd:   cCtx.Int("closeEnd"),
	}
}
