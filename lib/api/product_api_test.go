package api_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestProductApi(t *testing.T) {

	apiSession, saveSession := scenario.GetTokopediaApiClient()
	defer saveSession()

	t.Run("product list meta", func(t *testing.T) {
		hasil, err := apiSession.ProductListMeta()
		assert.NotEmpty(t, hasil)
		assert.Nil(t, err)
	})

	t.Run("test get product list", func(t *testing.T) {
		shopId := apiSession.AuthenticatedData.UserShopInfo.Info.ShopID
		filter := []model.Filter{
			{
				ID:    "pageSize",
				Value: []string{"20"},
			}, {
				ID:    "keyword",
				Value: []string{""},
			}, {
				ID:    "status",
				Value: []string{},
			},
			{
				ID:    "page",
				Value: []string{"1"},
			},
		}
		query := model.NewProductListVar(shopId, filter)
		hasil, err := apiSession.ProductList(query)
		assert.NotEmpty(t, hasil)
		assert.NotEmpty(t, hasil.Data.ProductList.Data)
		assert.Nil(t, err)
	})

	t.Run("test get product v3", func(t *testing.T) {
		productId := int64(4514276124)
		variable := model.NewProductV3Var(productId)
		hasil, err := apiSession.GetProductV3(variable)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
		assert.Equal(t, hasil.Data.GetProductV3.ProductID, productId)
		t.Log(hasil)
	})

	t.Run("test product add rule", func(t *testing.T) {

		hasil, err := apiSession.GetProductAddRule()
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
		assert.Equal(t, hasil.Data.ProductAddRule.Data.Eligible.Limit, 2000)
	})
}

// func TestProductAddApi(t *testing.T) {
// 	apiSession, saveSession := scenario.GetTokopediaApiClient()
// 	defer saveSession()

// 	payload := uploader.PayloadUpload{}

// 	fname := scenario.GetBaseTestAsset("payload.json")
// 	b, err := os.ReadFile(fname)
// 	assert.Nil(t, err)

// 	err = json.Unmarshal(b, &payload)
// 	assert.Nil(t, err)

// 	hasil, err := apiSession.ProductAdd(payload.GetProductAddVar())
// 	assert.Nil(t, err)
// 	hasilraw, err := json.Marshal(hasil)
// 	assert.Nil(t, err)
// 	t.Log(hasilraw)
// }
