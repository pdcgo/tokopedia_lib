package scenario

import (
	"github.com/pdcgo/tokopedia_lib/app/upload_app"
	"github.com/pdcgo/tokopedia_lib/lib/datasource"
	"gorm.io/gorm"
)

func GetDb() *gorm.DB {
	db := datasource.NewSqliteDatabase(GetBaseTestAsset("tokopedia_test.db"))
	db.AutoMigrate(&upload_app.AkunItem{})
	return db
}
