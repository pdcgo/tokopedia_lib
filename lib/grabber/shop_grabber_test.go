package grabber_test

import (
	"context"
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/grabber"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/zeebo/assert"
)

func TestShopGrabber(t *testing.T) {
	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	baseConfig := legacy_source.BaseConfig{
		BaseData: "../..",
	}
	database := scenario.GetMongoDatabase(t)

	productRepo := mongorepo.NewProductRepo(ctx, database)
	baseGrabber := grabber.NewBaseGrabber(api, &baseConfig, productRepo)
	baseGrabber.Filter.GrabBasic.LimitGrab = 1

	t.Run("test grab shop", func(t *testing.T) {
		shops := []string{
			"https://www.tokopedia.com/tokokaosdistropremium",
			"https://www.tokopedia.com/erigo",
		}
		grab := grabber.CreateShopListGrabber(baseGrabber, shops)
		grab.Run()
	})
}
