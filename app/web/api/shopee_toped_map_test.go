package api_test

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	"github.com/pdcgo/tokopedia_lib/app/config"
	"github.com/pdcgo/tokopedia_lib/app/web/api"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/category_mapper"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/pdcgo/v2_gots_sdk"
	"github.com/pdcgo/v2_gots_sdk/pdc_api"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestTokopediaCollectionList(t *testing.T) {
	db := scenario.GetDb()
	mongodb := scenario.GetMongoDatabase(t)

	pubapi, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	scenario.RunWebSdk(t, func(sdk *v2_gots_sdk.ApiSdk, sendApi scenario.SendRequest) error {
		prepo := mongorepo.NewProductRepo(mongodb)
		api.RegisterShopeeTopedMap(sdk.Group("tokopedia"), db, prepo, category_mapper.NewMapper(pubapi))

		t.Run("test api category mapper list", func(t *testing.T) {
			res := sendApi(&pdc_api.Api{
				Method:       http.MethodGet,
				RelativePath: "/tokopedia/mapper/category",
				Query: &api.TokopediaMapQuery{
					Namespace: "default",
				},
			})

			assert.Equal(t, http.StatusOK, res.Result().StatusCode)

			data, err := io.ReadAll(res.Result().Body)
			assert.Nil(t, err)

			hasil := []*api.TokopediaMapItem{}

			err = json.Unmarshal(data, &hasil)
			assert.Nil(t, err)

			assert.NotEmpty(t, hasil)
		})

		t.Run("test api tokopedia category mapper update", func(t *testing.T) {
			res := sendApi(&pdc_api.Api{
				Method:       http.MethodGet,
				RelativePath: "/tokopedia/mapper/category",
				Query: &api.TokopediaMapQuery{
					Namespace: "default",
				},
			})

			data, err := io.ReadAll(res.Result().Body)
			assert.Nil(t, err)

			dataBefore := []*api.TokopediaMapItem{}
			err = json.Unmarshal(data, &dataBefore)
			assert.Nil(t, err)
			assert.NotEmpty(t, dataBefore)

			randId := rand.Intn(100)
			tokopediaID := dataBefore[0].TokopediaID
			res = sendApi(&pdc_api.Api{
				Method:       http.MethodPut,
				RelativePath: "/tokopedia/mapper/map",
				Payload: []config.ShopeeMapItem{{
					ShopeeID:    int64(randId),
					TokopediaID: tokopediaID,
				}},
			})

			assert.Equal(t, http.StatusOK, res.Result().StatusCode)

			t.Run("test check data valid ketika sudah diupdate", func(t *testing.T) {
				res = sendApi(&pdc_api.Api{
					Method:       http.MethodGet,
					RelativePath: "/tokopedia/mapper/category",
					Query: &api.TokopediaMapQuery{
						Namespace: "default",
					},
				})

				data, err = io.ReadAll(res.Result().Body)
				assert.Nil(t, err)

				dataAfter := []*api.TokopediaMapItem{}
				err = json.Unmarshal(data, &dataAfter)
				assert.Nil(t, err)
				assert.NotEmpty(t, dataAfter)

				assert.NotEmpty(t, dataAfter)

				for _, catmap := range dataAfter {
					if catmap.TokopediaID == tokopediaID {
						t.Log(catmap, dataBefore[0], "asdasdasdasd")
						assert.Equal(t, int64(randId), catmap.ShopeeID)
					}

				}

			})

			t.Run("test update tidak hilang", func(t *testing.T) {

				defer db.Delete(&config.ShopeeMapItem{ShopeeID: 1000})

				res = sendApi(&pdc_api.Api{
					Method:       http.MethodPut,
					RelativePath: "/tokopedia/mapper/map",
					Payload: []config.ShopeeMapItem{
						{ShopeeID: 1000, TokopediaID: 1000},
						{ShopeeID: 2000, TokopediaID: 1000},
					},
				})
				assert.Equal(t, res.Code, 200)

				data := []*config.ShopeeMapItem{}
				err = db.Find(&data, &config.ShopeeMapItem{TokopediaID: 1000}).Error
				assert.Nil(t, err)
				assert.Equal(t, 2, len(data))
			})

		})

		rand.Intn(100)

		return nil
	})

}

func TestSavingMapItem(t *testing.T) {

	db := scenario.GetDb()

	t.Run("test set tokopedia id", func(t *testing.T) {
		item := api.NewShopeeMapSuggestItem(db, &mongorepo.ProductCategoryAgg{
			ID: 123,
		})

		err := item.SetTokopediaID(1233)
		assert.Nil(t, err)

		t.Run("test double", func(t *testing.T) {
			item := api.NewShopeeMapSuggestItem(db, &mongorepo.ProductCategoryAgg{
				ID: 12323,
			})

			err := item.SetTokopediaID(1233)
			assert.Nil(t, err)
		})
	})

	scen := scenario.NewScenario(t)
	scen.Base = scenario.GetBaseTestAsset("base_toni")

	scen.WithCopySqliteDatabase(func(db *gorm.DB) {
		t.Run("test set tokopedia id toni", func(t *testing.T) {
			item := api.NewShopeeMapSuggestItem(db, &mongorepo.ProductCategoryAgg{
				ID: 123,
			})

			err := item.SetTokopediaID(1233)
			assert.Nil(t, err)

			t.Run("test double toni", func(t *testing.T) {
				item := api.NewShopeeMapSuggestItem(db, &mongorepo.ProductCategoryAgg{
					ID: 12323,
				})

				err := item.SetTokopediaID(1233)
				assert.Nil(t, err)
			})
		})
	})
}
