package grab_handler_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/grab_handler"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {

	t.Run("test create cache handler product", func(t *testing.T) {

		pubapi, err := api_public.NewTokopediaApiPublic()
		assert.Nil(t, err)

		url := "https://www.tokopedia.com/milenialbook-1/novel-d-bijis"
		payload, err := model_public.NewPdpGetlayoutQueryVar(url)
		assert.Nil(t, err)

		layout, err := pubapi.PdpGetlayoutQuery(payload)
		assert.Nil(t, err)

		payloadp2 := model_public.NewPdpGetDataP2Var(layout.Data.PdpGetLayout)
		pdp, err := pubapi.PdpGetDataP2(payloadp2)
		assert.Nil(t, err)

		product, err := grab_handler.CreateCacheProduct("test", layout, pdp)

		assert.Nil(t, err)
		assert.Equal(t, "novel d`bijis", product.Name)
		assert.Equal(t, "kondisi bekas halaman lengkap ada cap stempel d halaman depan", product.Desc)
		assert.Equal(t, int64(10255126), product.Shop.Shopid)
		assert.Equal(t, int64(1771154061), product.Productid)
		assert.Equal(t, int64(10000), product.Price)
		assert.Equal(t, int64(751), product.CategoryId)
		assert.Equal(t, "test", product.Namespace)
	})
}
