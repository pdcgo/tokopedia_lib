package services_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/services"
	"github.com/pdcgo/tokopedia_lib/app/shopee/shopee_repo"
	"github.com/pdcgo/tokopedia_lib/app/upload_app/config"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestEtalase(t *testing.T) {
	db := scenario.GetDb()

	scen := scenario.NewScenario(t)
	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithUploadConfig(func(cfg *config.UploadBaseConfig) error { return nil }, func(cfg *config.UploadBaseConfig) {
			t.Log(cfg)
			scen.WithMongoDatabase(cfg, func(mongodb *mongo.Database) {

				shopeeAgg := shopee_repo.NewProductAggregate(mongodb.Collection("item"))
				service := services.NewEtalaseMapService(db, shopeeAgg)

				t.Run("test getting list etalase", func(t *testing.T) {
					_, err := service.GetListMap("default")
					assert.Nil(t, err)
				})

				t.Run("testing add map etalase", func(t *testing.T) {
					err := service.UpdateBulkMap([]*services.EtalaseMapItem{
						{EtalaseName: "test etalase", CategoryID: 1},
					})

					assert.Nil(t, err)

					err = service.UpdateBulkMap([]*services.EtalaseMapItem{
						{EtalaseName: "test etalase2", CategoryID: 2},
					})
					assert.Nil(t, err)

					t.Run("testing double add map", func(t *testing.T) {
						err := service.UpdateBulkMap([]*services.EtalaseMapItem{
							{EtalaseName: "test etalase", CategoryID: 1},
						})

						assert.Nil(t, err)
					})

				})

				t.Run("list etalase", func(t *testing.T) {
					hasil, err := service.ListEtalase()

					assert.NotEmpty(t, hasil)
					assert.Nil(t, err)
				})

				t.Run("test etalase delete ada", func(t *testing.T) {
					err := service.DeleteEtalase("test_etalase")
					assert.Nil(t, err)

					_, err = service.GetEtalase(1)

					assert.Nil(t, err)
				})

				t.Run("delete etalase tidak ada", func(t *testing.T) {
					err := service.DeleteEtalase("test_etalase")
					assert.Nil(t, err)

					datas, err := service.GetEtalase(4)
					t.Log(datas)
					assert.NotNil(t, err)
				})
			})
		})
	})

}
