package integration_test

import (
	"context"
	"testing"

	shopeeupapp "github.com/pdcgo/go_v2_shopeelib/app/upload_app"
	mongolib "github.com/pdcgo/go_v2_shopeelib/lib/mongo"
	"github.com/pdcgo/tokopedia_lib/app/upload_app/shopee_flow"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/datasource"
	"github.com/pdcgo/tokopedia_lib/lib/repo"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func addAccount(db *gorm.DB) error {
	akun := repo.AkunItem{
		Username:   "pdcthoni@gmail.com",
		Password:   "SilentIsMyMantra",
		Secret:     "IULIWGH6TIK3CZBKHGE27DBRLQ5LR5WQ",
		Collection: "default",
		AkunUploadStatus: repo.AkunUploadStatus{
			LimitUpload: 100,
			Active:      true,
		},
	}

	return db.Save(&akun).Error
}

func TestUploadFlow(t *testing.T) {
	scenario.CheckMongoActive(t)

	cfg := scenario.CheckConfig(t)

	sqlpath := scenario.GetBaseTestAsset("tokopedia_data.db")
	rootBase := scenario.GetBaseTestAsset()

	sqlitedb := datasource.NewSqliteDatabase(sqlpath)
	sqlitedb.AutoMigrate(repo.AkunItem{})

	err := addAccount(sqlitedb)
	assert.Nil(t, err)
	concurent := shopeeupapp.UploadConcurencyConfig{
		AccountConcurency: 1,
		ProductPerAccount: 2,
	}

	mdb := mongolib.NewDatabase(context.Background(), cfg.MongoUri, "kampretcode2")

	publicapi, err := api_public.NewTokopediaApiPublic()

	assert.Nil(t, err)
	flow := shopee_flow.NewShopeeToTopedFlow(rootBase, context.Background(), mdb, sqlitedb, &concurent, publicapi)
	flow.AkunIterator.Reset()
	t.Run("test getting double account", func(t *testing.T) {
		akun, _, _, _ := flow.AkunIterator.Get()
		akun2, _, _, _ := flow.AkunIterator.Get()

		assert.NotEqual(t, akun.Username, akun2.Username)
	})

	err = flow.Run()

	assert.Nil(t, err)
}
