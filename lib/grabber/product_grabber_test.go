package grabber_test

import (
	"context"
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/grabber"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestPageIterate(t *testing.T) {
	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	baseConfig := &legacy_source.BaseConfig{
		BaseData: "../..",
	}
	database := scenario.GetMongoDatabase(t)

	productRepo := mongorepo.NewProductRepo(ctx, database)
	baseGrabber := grabber.NewBaseGrabber(api, baseConfig, productRepo)
	baseGrabber.Filter.GrabBasic.LimitGrab = 1

	t.Run("test product keyword grabber", func(t *testing.T) {
		keywords := []string{
			"mousepad",
			"keyboard gaming",
		}
		grabber := grabber.NewProductListGrabber(baseGrabber, keywords)

		grabber.Run()
	})
	t.Run("test product category grabber first level", func(t *testing.T) {
		// 1759, Fashion Pria
		grabber := grabber.NewCategoryGrabber(baseGrabber, 1759)
		params := grabber.GenerateProductSearchParams()
		params.CategoryId = 1759

		grabber.Run()
	})
}
