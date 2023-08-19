package api_test

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	"github.com/pdcgo/tokopedia_lib/app/web/api"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/category_mapper"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/pdcgo/v2_gots_sdk"
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

		res := sendApi(&v2_gots_sdk.Api{
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
