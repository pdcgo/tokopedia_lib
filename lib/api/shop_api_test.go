package api_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/stretchr/testify/assert"
)

func TestShopApi(t *testing.T) {
	driver, _ := tokopedia_lib.NewDriverAccount("bethdunn892@outlook.com", "MZT2Zk8U", "FSR3 CTR2 5ZJX XIL5 TVK6 E72R HSRA U5GW")
	driver.DevMode = true
	api, saveSession, _ := driver.CreateApi()
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

}
