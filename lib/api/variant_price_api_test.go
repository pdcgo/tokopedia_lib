package api_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestVariantPriceApi(t *testing.T) {
	api, saveSession := scenario.GetTokopediaApiClient()
	defer saveSession()

	// Tas Ransel Wanita
	catId := 1917

	t.Run("test variant price validation valid", func(t *testing.T) {
		var variant model.Variant
		sourcebytes, err := os.ReadFile(scenario.GetBaseTestAsset("assets/variant_price_valid.json"))
		assert.Nil(t, err)

		err = json.Unmarshal(sourcebytes, &variant)
		assert.Nil(t, err)

		hasil, err := api.VariantPriceValidation(catId, &variant)
		assert.NotEmpty(t, hasil)
		assert.Nil(t, err)

		assert.Zero(t, hasil.GetPriceGab())
	})

	t.Run("test variant price validation invalid", func(t *testing.T) {
		var variant model.Variant
		sourcebytes, err := os.ReadFile(scenario.GetBaseTestAsset("assets/variant_price_invalid.json"))
		assert.Nil(t, err)

		err = json.Unmarshal(sourcebytes, &variant)
		assert.Nil(t, err)

		hasil, err := api.VariantPriceValidation(catId, &variant)
		assert.NotEmpty(t, hasil)
		assert.Nil(t, err)

		assert.Equal(t, hasil.GetPriceGab(), 3)
	})
}
