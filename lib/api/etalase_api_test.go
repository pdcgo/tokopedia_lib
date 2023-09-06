package api_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestEtalase(t *testing.T) {
	api, saveSession := scenario.GetTokopediaApiClient()
	defer saveSession()

	etalaseName := "pakaian gamis"
	t.Run("test create etalase", func(t *testing.T) {
		hasil, err := api.AddShopShowcase(etalaseName)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
	})

	t.Run("test get etalase", func(t *testing.T) {
		hasil, err := api.ShopShowcase()
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)

		found := false
		for _, eta := range hasil.Data.ShopShowcases.Result {
			if eta.Name == etalaseName {
				found = true

				t.Run("test delete etalase", func(t *testing.T) {
					hasil, err := api.DeleteShopShowcase(eta.ID)

					assert.Nil(t, err)
					assert.NotEmpty(t, hasil)
				})
			}
		}

		assert.True(t, found)
	})

	t.Run("test get etalase query", func(t *testing.T) {
		hasil, err := api.ShopShowcasesQuery(7125740)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
		assert.NotEmpty(t, hasil.Data.ShopShowcases.Result)
	})
}
