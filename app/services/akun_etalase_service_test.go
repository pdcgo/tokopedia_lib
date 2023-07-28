package services_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/services"
	"github.com/pdcgo/tokopedia_lib/app/shopee/shopee_repo"
	"github.com/pdcgo/tokopedia_lib/app/upload_app/config"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func TestAkunEtalase(t *testing.T) {
	scen := scenario.NewScenario(t)

	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithUploadConfig(func(cfg *config.UploadBaseConfig) error { return nil }, func(cfg *config.UploadBaseConfig) {
			scen.WithSqliteDatabase(func(db *gorm.DB) {
				scen.WithMongoDatabase(cfg, func(mongodb *mongo.Database) {
					agg := shopee_repo.NewProductAggregate(mongodb.Collection("item"))
					mapsrv := services.NewEtalaseMapService(db, agg)

					t.Run("test add update bulkmap mapping", func(t *testing.T) {
						err := mapsrv.UpdateBulkMap([]*services.EtalaseMapItem{
							{
								EtalaseName: "Gamis",
								CategoryID:  20,
							},
						})

						assert.Nil(t, err)
					})

					t.Run("test akun etalase service", func(t *testing.T) {
						apiclient, saveSession := scenario.GetTokopediaApiClient()
						defer saveSession()
						akunsrv := services.NewAkunEtalaseService(apiclient, mapsrv)

						t.Run("test refresh showcase", func(t *testing.T) {
							err := akunsrv.RefreshShowCase()
							assert.Nil(t, err)
						})
						t.Run("test create etalase", func(t *testing.T) {
							showcase, err := akunsrv.GetEtalase(20)

							assert.NotEmpty(t, showcase)
							assert.Nil(t, err)
						})

						t.Run("test create etalase map not found", func(t *testing.T) {
							_, err := akunsrv.GetEtalase(21)

							assert.NotNil(t, err)
						})
					})

				})

			})
		})
	})

}
