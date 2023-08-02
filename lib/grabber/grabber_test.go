package grabber_test

import (
	"context"
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/grab_handler"
	"github.com/pdcgo/tokopedia_lib/lib/grabber"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestGrabCategoryCsv(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	base := &legacy_source.BaseConfig{
		BaseData: "../..",
	}

	database := scenario.GetMongoDatabase(t)

	productRepo := mongorepo.NewProductRepo(ctx, database)
	tasker := legacy.NewGrabTasker(base.Path("data/tasker.json"))
	cacheHandler := grab_handler.NewCacheProductHandler(productRepo)

	baseGrab := grabber.NewBaseGrabber(api, base, tasker, cacheHandler)
	t.Run("test categorycsv", func(t *testing.T) {
		grabber := grabber.NewCategoryCsvGrabber(baseGrab)
		err := grabber.Run()
		assert.Nil(t, err)
	})
}

// func TestParseProductDetailParamsFromUrl(t *testing.T) {
// 	url := "https://www.tokopedia.com/acomeindone/acome-keyboard-mouse-wireless-portable-1600dpi-silikon-akm2000-keyboard-only?extParam=cmp%3D1%26ivf%3Dfalse&src=topads"
// 	params, err := grabber.ParseProductDetailParamsFromUrl(url)
// 	assert.Nil(t, err)
// 	assert.Equal(t, params, &model_public.PdpGetlayoutQueryVar{
// 		ShopDomain: "acomeindone",
// 		ProductKey: "acome-keyboard-mouse-wireless-portable-1600dpi-silikon-akm2000-keyboard-only",
// 		APIVersion: 1,
// 		ExtParam:   "cmp%3D1%26ivf%3Dfalse",
// 	})
// }

// func TestGenerateShopCoreInfoParamsFormUrl(t *testing.T) {
// 	t.Run("test parse with url", func(t *testing.T) {
// 		shopUrl := "https://www.tokopedia.com/acomeindone"
// 		params, err := grabber.GenerateShopCoreInfoParamsFormUrl(shopUrl)
// 		assert.Nil(t, err)
// 		assert.Equal(t, params, &model_public.ShopCoreInfoVar{
// 			ID:     0,
// 			Domain: "acomeindone",
// 		})
// 	})
// }
