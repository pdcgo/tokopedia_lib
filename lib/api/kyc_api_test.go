package api_test

import (
	"log"
	"testing"

	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestKycApi(t *testing.T) {

	api, saveSession := scenario.GetTokopediaApiClient()
	defer saveSession()

	t.Run("test get info kyc", func(t *testing.T) {

		hasil, err := api.GetInfoKyc()
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
		assert.NotEmpty(t, hasil.TypeList)
		log.Println(hasil)
	})

	// t.Run("test submit kyc", func(t *testing.T) {

	// 	driver, _ := tokopedia_lib.NewDriverAccount("bachtiarhamidah36@outlook.com", "blitar123", "F445LWKVRB3LCEBATTJPJ3PVOK6XZMB5")
	// 	// driver, _ := tokopedia_lib.NewDriverAccount("mrwicaksono690@gmail.com", "Semogaberkah", "KQ2YZZJDS2DAC7Y6K4HCTTXSS5B7N4IN")

	// 	api, save, _ := driver.CreateApi()
	// 	defer save()

	// 	imgKtp, err := os.Open(scenario.GetBaseTestAsset("ktp", "imgKtp.jpg"))
	// 	if !assert.Nil(t, err) {
	// 		defer imgKtp.Close()
	// 	}

	// 	imgSelfie, err := os.Open(scenario.GetBaseTestAsset("ktp", "imgSelfie.jpg"))
	// 	if !assert.Nil(t, err) {
	// 		defer imgSelfie.Close()
	// 	}

	// 	hasil, err := api.SubmitKyc(imgKtp, imgSelfie)
	// 	assert.Nil(t, err)
	// 	assert.NotEmpty(t, hasil)
	// 	log.Println(hasil)
	// })
}
