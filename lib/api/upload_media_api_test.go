package api_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestUploadProductImage(t *testing.T) {
	apiSession, saveSession := scenario.GetTokopediaApiClient()
	defer saveSession()

	t.Run("test upload image dengan url", func(t *testing.T) {
		result, err := apiSession.UploadImageFromUrl("https://down-id.img.susercontent.com/file/sg-11134201-22110-jehurfqhmckv02")
		assert.NotEmpty(t, result)
		assert.Nil(t, err)
	})
}

// func TestUploadImageChat(t *testing.T) {
// 	apiSession, saveSession := scenario.GetTokopediaApiClient()
// 	defer saveSession()

// 	t.Run("test upload product", func(t *testing.T) {
// 		shopId := strconv.Itoa(int(apiSession.AuthenticatedData.UserShopInfo.Info.ShopID))
// 		result, err := apiSession.UploadImageChat(shopId, "../../blob.jpeg")
// 		assert.NotEmpty(t, result)
// 		assert.Nil(t, err)
// 	})
// }
