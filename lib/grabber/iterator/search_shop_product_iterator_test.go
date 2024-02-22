package iterator_test

import (
	"context"
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/iterator"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestSearchShopProductIterator(t *testing.T) {

	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)
	scen := scenario.NewScenario(t)

	// shop: https://windomieofficialww.tokopedia.com/
	searchVar := model_public.NewShopProductVar(4544078)

	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithBaseConfig(func(base *legacy_source.BaseConfig) {

			t.Run("test search shop product iterator", func(t *testing.T) {

				ctx := context.Background()
				searchItems := []*model_public.ShopProductData{}

				err := iterator.IterateProductShopPage(api, ctx, searchVar, func(items []*model_public.ShopProductData) error {
					searchItems = append(searchItems, items...)
					return nil
				})

				assert.Nil(t, err)
				assert.Greater(t, len(searchItems), 80)
			})

			t.Run("test search shop product iterator context cancel", func(t *testing.T) {

				ctx, cancel := context.WithCancel(context.Background())
				cancel()

				searchItems := []*model_public.ShopProductData{}
				err := iterator.IterateProductShopPage(api, ctx, searchVar, func(items []*model_public.ShopProductData) error {
					searchItems = append(searchItems, items...)
					return nil
				})

				assert.Nil(t, err)
				assert.Equal(t, 0, len(searchItems))
			})

		})
	})
}
