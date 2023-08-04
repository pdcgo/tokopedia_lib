package scenario

import (
	"github.com/pdcgo/tokopedia_lib/app/config"
	"github.com/pdcgo/tokopedia_lib/lib/datasource"
	"github.com/pdcgo/tokopedia_lib/lib/repo"
	"gorm.io/gorm"
)

// TODO: bakalan deprecated soale ada test scenario
func GetDb() *gorm.DB {
	db := datasource.NewSqliteDatabase(GetBaseTestAsset("tokopedia_test.db"))
	db.AutoMigrate(&repo.AkunItem{}, &config.ShopeeMapItem{})
	return db
}
