package api_test

import (
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	"github.com/pdcgo/tokopedia_lib/app/web/api"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

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
