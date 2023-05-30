package api_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/stretchr/testify/assert"
)

func TestShopProducts(t *testing.T) {
	driver, _ := tokopedia_lib.NewDriverAccount("bethdunn892@outlook.com", "MZT2Zk8U", "FSR3 CTR2 5ZJX XIL5 TVK6 E72R HSRA U5GW")
	driver.DevMode = true
	apiSession, saveSession, _ := driver.CreateApi()
	defer saveSession()

	t.Run("test ShopProduct", func(t *testing.T) {
		variable := model.ShopProductVar{
			Sid:            "1409816",
			Page:           1,
			PerPage:        80,
			EtalaseID:      "etalase",
			Sort:           1,
			UserDistrictID: "2274",
			UserCityID:     "176",
			UserLat:        "",
			UserLong:       "",
		}
		hasil, err := apiSession.ShopProducts(&variable)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
	})
}
