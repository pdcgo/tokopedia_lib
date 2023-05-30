package api_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/stretchr/testify/assert"
)

func TestShopCoreInfo(t *testing.T) {
	driver, _ := tokopedia_lib.NewDriverAccount("bethdunn892@outlook.com", "MZT2Zk8U", "FSR3 CTR2 5ZJX XIL5 TVK6 E72R HSRA U5GW")
	driver.DevMode = true
	apiSession, saveSession, _ := driver.CreateApi()
	defer saveSession()

	t.Run("test shopCoreInfo", func(t *testing.T) {
		variable := model.ShopCoreInfoVar{
			ID:     0,
			Domain: "lenovo-tangerang",
		}
		hasil, err := apiSession.ShopCoreInfo(&variable)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
	})
}

func TestShopStatisticQuery(t *testing.T) {
	driver, _ := tokopedia_lib.NewDriverAccount("bethdunn892@outlook.com", "MZT2Zk8U", "FSR3 CTR2 5ZJX XIL5 TVK6 E72R HSRA U5GW")
	driver.DevMode = true
	apiSession, saveSession, _ := driver.CreateApi()
	defer saveSession()

	t.Run("test shopCoreInfo", func(t *testing.T) {
		variable := model.ShopStatisticQueryVar{
			ShopID:    11534215,
			ShopIDStr: "11534215",
		}
		hasil, err := apiSession.ShopStatisticQuery(&variable)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
	})
}

func TestGetShopOperationalHourStatus(t *testing.T) {
	driver, _ := tokopedia_lib.NewDriverAccount("bethdunn892@outlook.com", "MZT2Zk8U", "FSR3 CTR2 5ZJX XIL5 TVK6 E72R HSRA U5GW")
	driver.DevMode = true
	apiSession, saveSession, _ := driver.CreateApi()
	defer saveSession()

	t.Run("test shopCoreInfo", func(t *testing.T) {
		variable := model.ShopIdVar{
			ShopID: "11534215",
		}
		hasil, err := apiSession.GetShopOperationalHourStatus(&variable)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
	})
}

func TestVoucherListQuery(t *testing.T) {
	driver, _ := tokopedia_lib.NewDriverAccount("bethdunn892@outlook.com", "MZT2Zk8U", "FSR3 CTR2 5ZJX XIL5 TVK6 E72R HSRA U5GW")
	driver.DevMode = true
	apiSession, saveSession, _ := driver.CreateApi()
	defer saveSession()

	t.Run("test shopCoreInfo", func(t *testing.T) {
		variable := model.ShopIdVarInt{
			ShopID: 11534215,
		}
		hasil, err := apiSession.VoucherListQuery(&variable)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
	})
}

func TestShopNote(t *testing.T) {
	driver, _ := tokopedia_lib.NewDriverAccount("bethdunn892@outlook.com", "MZT2Zk8U", "FSR3 CTR2 5ZJX XIL5 TVK6 E72R HSRA U5GW")
	driver.DevMode = true
	apiSession, saveSession, _ := driver.CreateApi()
	defer saveSession()

	t.Run("test shopCoreInfo", func(t *testing.T) {
		variable := model.ShopNoteVar{
			ID:  "0",
			Sid: "11534215",
		}
		hasil, err := apiSession.ShopNote(&variable)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
	})
}
