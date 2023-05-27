package api_test

import (
	"strconv"
	"testing"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/stretchr/testify/assert"
)

func TestUploadProductImage(t *testing.T) {
	driver, _ := tokopedia_lib.NewDriverAccount("bethdunn892@outlook.com", "MZT2Zk8U", "FSR3 CTR2 5ZJX XIL5 TVK6 E72R HSRA U5GW")
	driver.Session.Ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"
	driver.DevMode = true
	apiSession, saveSession, _ := driver.CreateApi()
	defer saveSession()

	t.Run("test upload product", func(t *testing.T) {
		result, err := apiSession.UploadProductImage("../../blob.jpeg")
		assert.NotEmpty(t, result)
		assert.Nil(t, err)
	})
}

func TestUploadImageChat(t *testing.T) {
	driver, _ := tokopedia_lib.NewDriverAccount("bethdunn892@outlook.com", "MZT2Zk8U", "FSR3 CTR2 5ZJX XIL5 TVK6 E72R HSRA U5GW")
	driver.Session.Ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"
	driver.DevMode = true
	apiSession, saveSession, _ := driver.CreateApi()
	defer saveSession()

	t.Run("test upload product", func(t *testing.T) {
		shopId := strconv.Itoa(int(apiSession.AuthenticatedData.UserShopInfo.Info.ShopID))
		result, err := apiSession.UploadImageChat(shopId, "../../blob.jpeg")
		assert.NotEmpty(t, result)
		assert.Nil(t, err)
	})
}
