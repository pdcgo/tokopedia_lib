package chat

import (
	"github.com/pdcgo/tokopedia_lib/app/chat/config"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/lib/datasource"
	"gorm.io/gorm"
)

func CreateSqliteDatabase(conf *config.AppConfig) *gorm.DB {

	dst := conf.Path("tokopedia_chat.db")
	sqlitedb := datasource.NewSqliteDatabase(dst)

	sqlitedb.AutoMigrate(
		model.Account{},
		model.AccountData{},
		model.Group{},
		model.Order{},
		model.OrderItem{},
		model.OrderSheet{},
	)

	return sqlitedb
}
