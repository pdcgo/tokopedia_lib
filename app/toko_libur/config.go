package toko_libur

import (
	"encoding/json"
	"os"

	"github.com/pdcgo/common_conf/pdc_application"
	"github.com/pdcgo/tokopedia_lib/lib/api"
)

type TokopediaTokoLiburConfig struct {
	BaseConfig pdc_application.BaseApplication `json:"-"`
	Libur      bool                            `json:"libur"`
	Report     string                          `json:"report"`
	Limit      int                             `json:"limit"`
	CloseStart int                             `json:"closeStart"`
	CloseEnd   int                             `json:"closeEnd"`
}

func (c *TokopediaTokoLiburConfig) GetReportPath() string {
	return c.BaseConfig.Path(c.Report)
}

func (c *TokopediaTokoLiburConfig) CreateCloseShopInput() *api.CloseShopScheduleInput {

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

func NewApplicationConfig(base pdc_application.BaseApplication) *TokopediaTokoLiburConfig {

	config := TokopediaTokoLiburConfig{
		BaseConfig: base,
		Report:     "tokopedia_toko_libur.csv",
		Limit:      3,
	}

	fname := base.Path("data", "tokopedia_libur_config")
	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return &config
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.Decode(&config)

	return &config
}

func SaveApplicationConfig(base pdc_application.BaseApplication, config *TokopediaTokoLiburConfig) error {

	fname := base.Path("data", "tokopedia_libur_config")
	file, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewEncoder(file)
	return decoder.Encode(&config)
}
