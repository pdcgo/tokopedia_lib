package scenario

import (
	"context"
	"sync"
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/app/chat_app/chat_store"
	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func GetDatabase() *gorm.DB {
	dbname := GetBaseTestAsset("database_test.db")
	return chat_store.NewChatDb(dbname)
}

var mongoDB sync.Once
var db *mongo.Database

func GetMongoDatabase(t *testing.T) *mongo.Database {

	var err error

	mongoDB.Do(func() {
		CheckConfig(t)
		CheckMongoActive(t)

		db = mongorepo.NewDatabase(context.Background(), CfgTest.MongoUri, "kampretcode2")
		if err != nil {
			t.Skip("skipping database test because : ", err)
		}
	})

	return db
}
