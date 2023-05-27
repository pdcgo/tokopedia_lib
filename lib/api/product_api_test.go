package api_test

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/api"
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
	driver.Session.Ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"
	driver.DevMode = true
	apiSession, saveSession, _ := driver.CreateApi()
	defer saveSession()

	t.Run("product add", func(t *testing.T) {
		input := api.InputNoVariant{
			InputVariable: api.InputVariable{
				Pictures: struct {
					Data []struct {
						UploadIds string "json:\"uploadIds\""
					} "json:\"data\""
				}{
					Data: []struct {
						UploadIds string "json:\"uploadIds\""
					}{{
						UploadIds: "efd84e98-1fa6-41a2-a69c-eb969ad61191",
					}},
				},
				ProductName: "Fantech CRYPTO VX7 Mouse Gaming Macro - Hitam",
				Category: struct {
					ID string "json:\"id\""
				}{
					ID: "4011",
				},
				Condition:     "NEW",
				MinOrder:      1,
				PriceCurrency: "IDR",
				Weight:        124,
				WeightUnit:    "GR",
				MustInsurance: true,
				Menus: []struct {
					MenuID string "json:\"menuID\""
				}{},
				Annotations: []string{"5037"},
				Description: "Fantech VX7 Crypto adalah wired gaming mouse dengan gaming optical sensor, sensitivitas 200-8000 DPI, 60 IPS speed, akselerasi 20g, dan juga polling rate 125Hz. Mouse gaming ini juga dilengkapi dengan switch yang memiliki lifecycle hingga 10 juta klik, teflon mouse skates yang besar, 6 tombol yang bisa diatur, serta 4 mode efek pencahayaan. \n",
				Dimention: struct {
					Width  int "json:\"width\""
					Height int "json:\"height\""
					Length int "json:\"length\""
				}{
					Width:  12,
					Height: 8,
					Length: 16,
				},
			},
			Sku:    "BlackHead",
			Stock:  12,
			Price:  125000,
			Status: "LIMITED",
		}
		variable := api.ProductAddVar{
			Input: input,
		}
		hasil, err := apiSession.ProductAdd(&variable)
		hasilraw, _ := json.Marshal(hasil)
		t.Log(string(hasilraw), "hasilnya seperti ini")
		assert.NotEmpty(t, hasil)
		if err != nil {
			// cek product name telah dipakai errCode = "2"
			assert.Equal(t, hasil.Data.ProductAddV3.Header.ErrorCode, "2")
		}
	})
}

func TestProductList(t *testing.T) {
	driver, _ := tokopedia_lib.NewDriverAccount("bethdunn892@outlook.com", "MZT2Zk8U", "FSR3 CTR2 5ZJX XIL5 TVK6 E72R HSRA U5GW")
	driver.Session.Ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"
	driver.DevMode = true
	apiSession, saveSession, _ := driver.CreateApi()
	defer saveSession()

	t.Run("test get product list", func(t *testing.T) {
		shopId := strconv.Itoa(int(apiSession.AuthenticatedData.UserShopInfo.Info.ShopID))
		query := api.ProductListVar{
			ShopID: shopId,
			Filter: []struct {
				ID    string   "json:\"id\""
				Value []string "json:\"value\""
			}{
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
			Sort: struct {
				ID    string "json:\"id\""
				Value string "json:\"value\""
			}{
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
