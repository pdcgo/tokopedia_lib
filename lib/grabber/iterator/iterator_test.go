package iterator_test

import (
	"context"
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/csv"
	"github.com/pdcgo/tokopedia_lib/lib/grabber"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/iterator"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/stretchr/testify/assert"
)

func TestIteratorCategoryCsv(t *testing.T) {
	base := legacy_source.BaseConfig{
		BaseData: "../../..",
	}

	err := iterator.IterateCategoryCsv(&base, func(item *csv.CategoryCsv) error {
		assert.NotNil(t, item)
		return nil
	})
	assert.Nil(t, err)
}

func TestIteratorkeywords(t *testing.T) {
	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)
	base := legacy_source.BaseConfig{
		BaseData: "../../..",
	}

	ctx, cancel := context.WithCancel(context.Background())
	grabTokopedia := legacy.NewGrabTokopedia(&base)

	searchVar := grabber.CreateGrabSearchVar(grabTokopedia)
	searchVar.Query = "keyboard"

	err = iterator.IterateSearchPage(api, ctx, searchVar, func(items []*model_public.ProductSearch) error {
		assert.NotEqual(t, len(items), 0)
		cancel()
		return nil
	})
	assert.Nil(t, err)
}

func TestIteratorCategory(t *testing.T) {
	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)
	base := legacy_source.BaseConfig{
		BaseData: "../../..",
	}

	ctx, cancel := context.WithCancel(context.Background())
	grabTokopedia := legacy.NewGrabTokopedia(&base)

	searchVar := grabber.CreateGrabSearchVar(grabTokopedia)
	searchVar.CategoryId = 340

	err = iterator.IterateSearchPage(api, ctx, searchVar, func(items []*model_public.ProductSearch) error {
		assert.NotEqual(t, len(items), 0)
		cancel()
		return nil
	})
	assert.Nil(t, err)
}

func TestIteratorShopProducts(t *testing.T) {
	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)
	ctx, cancel := context.WithCancel(context.Background())

	searchVar := model_public.NewShopProductVar("53089753")
	err = iterator.IterateProductShopPage(api, ctx, searchVar, func(items []*model_public.ShopProductData) error {
		assert.NotEqual(t, len(items), 0)
		cancel()
		return nil
	})
	assert.Nil(t, err)
}
