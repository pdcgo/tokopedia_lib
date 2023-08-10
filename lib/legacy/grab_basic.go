package legacy

import (
	"sync"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/helper"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
)

type GrabBasic struct {
	legacy.GrabBasic
	UsePriceDiscount bool `json:"use_price_discount"`
}

var grabBasicOnce sync.Once
var grabBasic = &GrabBasic{
	GrabBasic: legacy.GrabBasic{
		Concurrent:     100,
		LimitGrab:      100,
		LastReviewDays: 7,
		Penjualan:      100,
		Prosentase:     20,
		Stock:          0,
	},
	UsePriceDiscount: false,
}

func NewGrabBasic(base *legacy_source.BaseConfig) *GrabBasic {
	grabBasicOnce.Do(func() {
		path := base.Path("data", "grab_config")
		helper.JsonLoad(path, grabBasic)
	})

	return grabBasic
}

func SaveGrabBasic(base legacy_source.BaseConfig, setting *GrabBasic) error {
	path := base.Path("data", "grab_config")
	err := helper.JsonDump(path, setting)
	return err
}
