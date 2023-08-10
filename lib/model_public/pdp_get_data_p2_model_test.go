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
		productName := hasil.GetProductName()
		assert.Equal(t, productName, "Jaket Zipper Hoodie Mocca - M")
	})

	t.Run("test get percentage amount", func(t *testing.T) {
		percentageAmount := hasil.GetPercentageAmount()
		assert.Equal(t, percentageAmount, 10)
	})

	t.Run("test get stock", func(t *testing.T) {
		stock := hasil.Getstock()
		assert.Equal(t, stock, 99997)
	})

}
