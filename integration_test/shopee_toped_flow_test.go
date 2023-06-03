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

func addAccount(db *gorm.DB) {
	akun := repo.AkunItem{
		Username: "pdcthoni@gmail.com",
		Password: "SilentIsMyMantra",
		Secret:   "IULIWGH6TIK3CZBKHGE27DBRLQ5LR5WQ",
		AkunUploadStatus: repo.AkunUploadStatus{
			LimitUpload: 100,
			Active:      true,
		},
	}

	db.Save(&akun)
}

func TestUploadFlow(t *testing.T) {
	scenario.CheckMongoActive(t)

	cfg := scenario.CheckConfig(t)

	sqlpath := scenario.GetBaseTestAsset("tokopedia_data.db")
	rootBase := scenario.GetBaseTestAsset()

	sqlitedb := datasource.NewSqliteDatabase(sqlpath)
	sqlitedb.AutoMigrate(repo.AkunItem{})

	concurent := shopeeupapp.UploadConcurencyConfig{
		AccountConcurency: 1,
		ProductPerAccount: 100,
	}

	mdb := mongolib.NewDatabase(context.Background(), cfg.MongoUri, "kampretcode2")

	publicapi, err := api_public.NewTokopediaApiPublic()

	assert.Nil(t, err)
	flow := shopee_flow.NewShopeeToTopedFlow(rootBase, context.Background(), mdb, sqlitedb, &concurent, publicapi)

	err = flow.Run()
	assert.Nil(t, err)
}
