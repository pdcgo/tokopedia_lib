package datasource

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewSqliteDatabase(fname string) *gorm.DB {
	devmode := os.Getenv("DEV_MODE") != ""

	config := gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	if devmode {
		config.Logger = logger.Default.LogMode(logger.Silent)
	}

	db, err := gorm.Open(sqlite.Open(fname), &config)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
