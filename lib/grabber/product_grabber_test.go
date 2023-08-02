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
	baseGrabber.Filter.GrabBasic.LimitGrab = 150

	t.Run("test product keyword grabber", func(t *testing.T) {
		keywords := []string{
			"mousepad",
			"keyboard gaming",
		}
		grabber := grabber.NewProductListGrabber(baseGrabber, keywords)
		params := grabber.GenerateProductSearchParams()
		params.Query = keywords[0]
		products, err := grabber.GetProducts(params)
		assert.Nil(t, err)
		assert.NotEqual(t, len(products), 0)
	})
	t.Run("test product category grabber first level", func(t *testing.T) {
		// 1759, Fashion Pria
		grabber := grabber.NewCategoryGrabber(baseGrabber, 1759)
		params := grabber.GenerateProductSearchParams()
		params.CategoryId = 1759

		products, err := grabber.GetProducts(params)
		assert.Nil(t, err)
		assert.NotEqual(t, len(products), 0)
		assert.Equal(t, products[0].CategoryID, 1759)
	})
	t.Run("test product category grabber last level", func(t *testing.T) {
		// 297, "Komputer & Laptop"
		// 338, "Aksesoris Komputer & Laptop",
		// 340, "Keyboard"
		grabber := grabber.NewCategoryGrabber(baseGrabber, 340)
		params := grabber.GenerateProductSearchParams()
		params.CategoryId = 340

		products, err := grabber.GetProducts(params)
		assert.Nil(t, err)
		assert.NotEqual(t, len(products), 0)
		assert.Equal(t, products[0].Category, 340)
	})
}
