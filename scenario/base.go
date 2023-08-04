package scenario

import (
	"context"
	"path/filepath"
	"runtime"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetBaseTestAsset(elem ...string) string {
	_, filename, _, _ := runtime.Caller(0)
	basedir := filepath.Dir(filename)

	listPath := []string{basedir, "../test"}

	listPath = append(listPath, elem...)

	return filepath.Join(listPath...)
}

func NewMongoDatabase(ctx context.Context, dburi string, dbname string) *mongo.Database {
	clientOptions := options.Client().ApplyURI(dburi)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		panic(err)
	}

	return client.Database(dbname)
}
