package datasource

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqliteDatabase(fname string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(fname), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
