package model_public_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestParsingPDPLayout(t *testing.T) {
	fname := scenario.GetBaseTestAsset("api_response", "pdplayout.json")
	data, err := os.ReadFile(fname)
	assert.Nil(t, err)

	var hasil model_public.PdpGetLayout

	err = json.Unmarshal(data, &hasil)
	assert.Nil(t, err)

	var foundmedia bool
	for _, comp := range hasil.Components {
		switch component := comp.(type) {
		case *model_public.MediaComponent:
			t.Log(component)
			foundmedia = true
		}
	}
	assert.True(t, foundmedia)

	t.Run("test get content", func(t *testing.T) {
		com, err := model_public.GetComponent[model_public.ProductDetailComponent](&hasil)
		assert.Nil(t, err)

		desc, err := com.Data[0].GetContent(model_public.DeskripsiTitle)
		assert.Nil(t, err)
		assert.NotEmpty(t, desc)

	})

	t.Run("test get product name", func(t *testing.T) {
		productName, err := hasil.GetProductName()
		assert.Nil(t, err)
		assert.Equal(t, productName, "Jaket Zipper Hoodie Mocca - M")
	})

	t.Run("test get percentage amount", func(t *testing.T) {
		percentageAmount, err := hasil.GetPercentageAmount()
		assert.Nil(t, err)
		assert.Equal(t, percentageAmount, 10)
	})

	t.Run("test get stock", func(t *testing.T) {
		stock, err := hasil.GetStock()
		assert.Nil(t, err)
		assert.Equal(t, stock, 99997)
	})

	t.Run("test get price", func(t *testing.T) {
		price, err := hasil.GetPrice()
		assert.Nil(t, err)
		assert.Equal(t, price, 66900)
	})

	t.Run("test get price before discount", func(t *testing.T) {
		price, err := hasil.GetPriceBeforeDiscount()
		assert.Nil(t, err)
		assert.Equal(t, price, 89999)
	})

	t.Run("test get images", func(t *testing.T) {
		images, err := hasil.GetImages()
		assert.Nil(t, err)
		assert.Equal(t, images, []string{
			"https://images.tokopedia.net/img/cache/700/product-1/2019/8/1/2999308/2999308_d0a9ad09-823a-4695-a9e0-b874c5290cc4_1000_1000",
			"https://images.tokopedia.net/img/cache/700/product-1/2019/8/1/2999308/2999308_47bb5918-eaaf-4be8-a4ed-d17c23a01395_1000_1000",
			"https://images.tokopedia.net/img/cache/700/product-1/2019/8/1/2999308/2999308_492e946f-1f78-49b1-b1b2-1fd506ac68bb_800_800",
		})
	})

	t.Run("test get description", func(t *testing.T) {
		desc, err := hasil.GetDescription()
		assert.Nil(t, err)
		assert.Equal(t, desc, "***HARAP DI BACA SAMPAI SELESAI***\n\nJaket Standar Distro Kualitas Export.\n")
	})
}
