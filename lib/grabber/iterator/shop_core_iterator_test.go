package iterator_test

import (
	"context"
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/core_concept"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/iterator"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/stretchr/testify/assert"
)

type ShopProd struct {
	t      *testing.T
	url    string
	shopID int
}

func (shop *ShopProd) GetShopID() int {
	return shop.shopID
}

func (shop *ShopProd) SetStatistic(data *model_public.ShopStatisticQueryResp) error {
	assert.NotEmpty(shop.t, data)
	return nil
}

func (shop *ShopProd) GetShopUrl() string {
	return shop.url
}

func (shop *ShopProd) SetShopCore(data *model_public.ShopCoreInfoResp) error {
	assert.NotEmpty(shop.t, data)
	shopdata := data.Data.Result[0]

	shop.shopID = data.Data.Result[0].ShopCore.ShopID

	assert.NotEmpty(shop.t, shopdata.ShopAssets.Avatar)

	return nil

}

func TestShopCoreIterator(t *testing.T) {
	ctxErr := core_concept.NewTaskContext(context.Background())

	productsChan := make(chan []*ShopProd, 1)

	go func() {
		defer close(productsChan)
		productsChan <- []*ShopProd{
			{
				t:   t,
				url: "https://www.tokopedia.com/gmbest-store",
			},
			{
				t:   t,
				url: "https://www.tokopedia.com/burangrangarchery",
			},
		}
	}()

	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	shopcore, err := iterator.BatchShopCore(ctxErr, productsChan, nil, 10, 5, api)
	assert.Nil(t, err)

	t.Run("test iterate statistic", func(t *testing.T) {
		defer ctxErr.Cancel()
		statpipe, err := iterator.BatchShopStatistic(ctxErr, shopcore, 10, 5, api)
		assert.Nil(t, err)
		for shop := range statpipe {
			t.Log(shop)

		}
	})
Parent:
	for {
		select {
		case err := <-ctxErr.Err:
			assert.Nil(t, err)
		case <-ctxErr.Ctx.Done():
			break Parent
		}

	}

}
