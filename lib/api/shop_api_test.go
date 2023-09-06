package api_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestShopApi(t *testing.T) {

	api, saveSession := scenario.GetTokopediaApiClient()
	defer saveSession()

	t.Run("test shop info api", func(t *testing.T) {
		hasil, err := api.ShopInfoByID()
		assert.NotEmpty(t, hasil)
		assert.Nil(t, err)
	})
	t.Log("---------------")
	t.Run("test gold merchant api", func(t *testing.T) {
		ghasil, err := api.GoldGetPMOSStatus()
		assert.NotEmpty(t, ghasil)
		assert.Nil(t, err)
	})

	t.Log("---------------")
	t.Run("test shopscore", func(t *testing.T) {
		ghasil, err := api.GetShopScoreLevel()
		assert.NotEmpty(t, ghasil)
		assert.Nil(t, err)
	})

	t.Log("---------------")
	t.Run("test shop location all", func(t *testing.T) {
		ghasil, err := api.GetShopLocationAll(7125740)
		assert.NotEmpty(t, ghasil)
		assert.Nil(t, err)
	})
}
