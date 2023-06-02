package scenario

import (
	"github.com/pdcgo/tokopedia_lib/lib/datasource"
	"github.com/pdcgo/tokopedia_lib/lib/repo"
	"gorm.io/gorm"
)

func GetDb() *gorm.DB {
	db := datasource.NewSqliteDatabase(GetBaseTestAsset("tokopedia_test.db"))
	db.AutoMigrate(&repo.AkunItem{})
	return db
}
