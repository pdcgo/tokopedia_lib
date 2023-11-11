package api_test

import (
	"strconv"
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
		shopId := strconv.Itoa(int(apiSession.AuthenticatedData.UserShopInfo.Info.ShopID))
		query := model.ProductListVar{
			ShopID: shopId,
			Filter: []model.Filter{
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
			},
			Sort: model.Sort{
				ID:    "DEFAULT",
				Value: "DESC",
			},
			ExtraInfo:   []string{"view", "topads", "rbac", "price-suggestion"},
			WarehouseID: "",
		}

		hasil, err := apiSession.ProductList(&query)
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
		assert.Equal(t, hasil.Data.ProductAddRule.Data.Eligible.Limit, 200)
	})
}

func TestProductAddApi(t *testing.T) {
	// apiSession, saveSession := scenario.GetTokopediaApiClient()
	// defer saveSession()

	// t.Run("product add", func(t *testing.T) {
	// 	input := model.InputNoVariant{
	// 		InputVariable: model.InputVariable{
	// 			Pictures: struct {
	// 				Data []model.Pictures "json:\"data\""
	// 			}{
	// 				Data: []model.Pictures{{
	// 					UploadIds: "efd84e98-1fa6-41a2-a69c-eb969ad61191",
	// 				}},
	// 			},
	// 			ProductName: "Fantech CRYPTO VX7 Mouse Gaming Macro - Hitam",
	// 			Category: model.Category{
	// 				ID: "4011",
	// 			},
	// 			Condition:     "NEW",
	// 			MinOrder:      1,
	// 			PriceCurrency: "IDR",
	// 			Weight:        124,
	// 			WeightUnit:    "GR",
	// 			MustInsurance: true,
	// 			Menus: []struct {
	// 				MenuID string "json:\"menuID\""
	// 			}{},
	// 			Annotations: []string{"5037"},
	// 			Description: "Fantech VX7 Crypto adalah wired gaming mouse dengan gaming optical sensor, sensitivitas 200-8000 DPI, 60 IPS speed, akselerasi 20g, dan juga polling rate 125Hz. Mouse gaming ini juga dilengkapi dengan switch yang memiliki lifecycle hingga 10 juta klik, teflon mouse skates yang besar, 6 tombol yang bisa diatur, serta 4 mode efek pencahayaan. \n",
	// 			Dimention: model.Dimension{
	// 				Width:  12,
	// 				Height: 8,
	// 				Length: 16,
	// 			},
	// 		},
	// 		Sku:    "BlackHead",
	// 		Stock:  12,
	// 		Price:  125000,
	// 		Status: "LIMITED",
	// 	}
	// 	variable := model.ProductAddVar{
	// 		Input: input,
	// 	}
	// 	hasil, err := apiSession.ProductAdd(&variable)
	// 	hasilraw, _ := json.Marshal(hasil)
	// 	t.Log(string(hasilraw), "hasilnya seperti ini")
	// 	assert.NotEmpty(t, hasil)
	// 	if err != nil {
	// 		// cek product name telah dipakai errCode = "2"
	// 		assert.Equal(t, hasil.Data.ProductAddV3.Header.ErrorCode, "2")
	// 	}
	// })
}
