package api_test

import (
	"strconv"
	"testing"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/stretchr/testify/assert"
)

func TestProductApi(t *testing.T) {
	driver, _ := tokopedia_lib.NewDriverAccount("bethdunn892@outlook.com", "MZT2Zk8U", "FSR3 CTR2 5ZJX XIL5 TVK6 E72R HSRA U5GW")
	driver.DevMode = true
	apiSession, saveSession, _ := driver.CreateApi()
	defer saveSession()

	t.Run("product list meta", func(t *testing.T) {
		hasil, err := apiSession.ProductListMeta()
		assert.NotEmpty(t, hasil)
		assert.Nil(t, err)
	})

}

func TestProductAddApi(t *testing.T) {
	driver, _ := tokopedia_lib.NewDriverAccount("bethdunn892@outlook.com", "MZT2Zk8U", "FSR3 CTR2 5ZJX XIL5 TVK6 E72R HSRA U5GW")
	driver.DevMode = true
	t.Log("muktar jangan test add product di otomatis. akunnya bukan milik kita")
	// apiSession, saveSession, _ := driver.CreateApi()
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

func TestProductList(t *testing.T) {
	driver, _ := tokopedia_lib.NewDriverAccount("bethdunn892@outlook.com", "MZT2Zk8U", "FSR3 CTR2 5ZJX XIL5 TVK6 E72R HSRA U5GW")
	driver.DevMode = true
	apiSession, saveSession, _ := driver.CreateApi()
	defer saveSession()

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
		assert.Nil(t, err)
	})
}

func TestGetProductV3(t *testing.T) {
	driver, _ := tokopedia_lib.NewDriverAccount("bethdunn892@outlook.com", "MZT2Zk8U", "FSR3 CTR2 5ZJX XIL5 TVK6 E72R HSRA U5GW")
	driver.DevMode = true
	apiSession, saveSession, _ := driver.CreateApi()
	defer saveSession()

	t.Run("test getProductV3", func(t *testing.T) {
		variable := model.GetProductV3Var{
			ProductID: "9781591960",
			Options: model.Options{
				Basic:       true,
				Menu:        true,
				Shop:        true,
				Category:    true,
				Wholesale:   true,
				Preorder:    true,
				Picture:     true,
				Sku:         true,
				Lock:        true,
				Variant:     true,
				Video:       true,
				Edit:        true,
				TxStats:     true,
				Dimension:   true,
				CustomVideo: true,
			},
			ExtraInfo: struct {
				Event bool "json:\"event\""
			}{
				Event: false,
			},
		}
		hasil, err := apiSession.GetProductV3(&variable)
		assert.NotEmpty(t, hasil)
		assert.Nil(t, err)
		t.Log(hasil)
	})
}
