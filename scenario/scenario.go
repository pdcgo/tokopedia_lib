package scenario

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/pdcgo/tokopedia_lib/app/upload_app/config"
	"github.com/pdcgo/tokopedia_lib/lib/datasource"
	"github.com/pdcgo/tokopedia_lib/lib/repo"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Scenario struct {
	t    *testing.T
	Base string
}

func NewScenario(t *testing.T) *Scenario {
	scen := Scenario{
		Base: GetBaseTestAsset(""),
	}

	return &scen
}

func (scen *Scenario) WithBase(handler func(dirbase string, scen *Scenario)) {

	id := uuid.New()
	mockbase := filepath.Join(scen.Base, "temp", id.String())
	err := os.MkdirAll(mockbase, os.ModePerm)
	assert.Nil(scen.t, err)

	err = os.MkdirAll(filepath.Join(mockbase, "data"), os.ModePerm)
	assert.Nil(scen.t, err)

	newscen := Scenario{
		t:    scen.t,
		Base: mockbase,
	}

	defer func() {
		os.RemoveAll(mockbase)
	}()

	handler(mockbase, &newscen)

}

func (scen *Scenario) WithUploadConfig(cfgmodifier func(cfg *config.UploadBaseConfig) error, handler func(cfg *config.UploadBaseConfig)) {
	cfg := config.UploadBaseConfig{
		Database: config.DatabaseConfig{
			DbURI:  "mongodb://root:password@localhost",
			DbName: "kampretcode2",
		},
	}
	err := cfgmodifier(&cfg)
	assert.Nil(scen.t, err)

	fname := filepath.Join(scen.Base, "data", "config.json")
	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)

	assert.Nil(scen.t, err)
	err = json.NewEncoder(file).Encode(&cfg)
	assert.Nil(scen.t, err)

	file.Close()
	assert.Nil(scen.t, err)

	defer func() {
		os.Remove(fname)
	}()

	handler(&cfg)
}

func (scen *Scenario) WithMongoDatabase(cfg *config.UploadBaseConfig, handler func(mongodb *mongo.Database)) {
	db := NewMongoDatabase(context.TODO(), cfg.Database.DbURI, cfg.Database.DbName)
	handler(db)
}

func (scen *Scenario) WithSqliteDatabase(handler func(db *gorm.DB)) {
	dbloc := filepath.Join(scen.Base, "tokopedia_test.db")
	db := datasource.NewSqliteDatabase(dbloc)

	db.AutoMigrate(&repo.AkunItem{})

	handler(db)
}

func copyFile(src string, dst string) {
	r, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer r.Close()
	w, err := os.Create(dst)
	if err != nil {
		panic(err)
	}
	defer w.Close()
	w.ReadFrom(r)
}

func (scen *Scenario) WithCopySqliteDatabase(handler func(db *gorm.DB)) {
	dbloc := filepath.Join(scen.Base, "tokopedia_test.db")
	dbdest := filepath.Join(scen.Base, "tokopedia_test_temp.db")
	copyFile(dbloc, dbdest)
	db := datasource.NewSqliteDatabase(dbdest)
	db.AutoMigrate(&repo.AkunItem{})

	defer func() {
		os.Remove(dbdest)
	}()

	handler(db)
}
