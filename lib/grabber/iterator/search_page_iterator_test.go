package iterator_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/iterator"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestSearchPageIterator(t *testing.T) {

	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)
	scen := scenario.NewScenario(t)

	searchVar := model_public.NewSearchProductVar()
	searchVar.Query = url.QueryEscape("jantung ayam")
	searchVar.PriceMin = 5000
	searchVar.PriceMax = 15000

	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithBaseConfig(func(base *legacy_source.BaseConfig) {

			t.Run("test search page iterator", func(t *testing.T) {

				ctx := context.Background()
				searchItems := []*model_public.ProductSearch{}

				err := iterator.IterateSearchPage(&iterator.IterateConfig{
					ChuckSize: 10,
				}, api, ctx, searchVar, func(items []*model_public.ProductSearch) error {

					searchItems = append(searchItems, items...)
					assert.LessOrEqual(t, len(items), 10)

					return nil
				})

				assert.Nil(t, err)
				assert.Greater(t, len(searchItems), 10)
			})

			t.Run("test search page iterator context cancel", func(t *testing.T) {

				ctx, cancel := context.WithCancel(context.Background())
				cancel()

				searchItems := []*model_public.ProductSearch{}
				err := iterator.IterateSearchPage(&iterator.IterateConfig{
					ChuckSize: 10,
				}, api, ctx, searchVar, func(items []*model_public.ProductSearch) error {
					searchItems = append(searchItems, items...)
					return nil
				})

				assert.Nil(t, err)
				assert.Equal(t, 0, len(searchItems))
			})

		})
	})
}
