package scenario

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
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

func (scen *Scenario) Path(data ...string) string {
	data = append([]string{scen.Base}, data...)
	return filepath.Join(data...)
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

func (scen *Scenario) WithBaseConfig(handler func(base *legacy_source.BaseConfig)) {

	base := legacy_source.BaseConfig{
		BaseData: scen.Base,
	}
	handler(&base)
}

func (scen *Scenario) CreateFile(data []byte, fname string) func() {

	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	assert.Nil(scen.t, err)

	_, err = file.Write(data)
	assert.Nil(scen.t, err)

	file.Close()
	assert.Nil(scen.t, err)

	remove := func() {
		os.Remove(fname)
	}

	return remove
}

func (scen *Scenario) createConfigFile(data any, fname string) func() {

	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	assert.Nil(scen.t, err)

	err = json.NewEncoder(file).Encode(data)
	assert.Nil(scen.t, err)

	file.Close()
	assert.Nil(scen.t, err)

	remove := func() {
		os.Remove(fname)
	}

	return remove
}

func (scen *Scenario) WithUploadConfig(cfgmodifier func(cfg *config.UploadBaseConfig) error, handler func(cfg *config.UploadBaseConfig)) {
	cfg := config.UploadBaseConfig{
		Database: config.DatabaseConfig{
			DbURI:  "mongodb://localhost:27017",
			DbName: "kampretcode2",
		},
	}
	err := cfgmodifier(&cfg)
	assert.Nil(scen.t, err)

	fname := filepath.Join(scen.Base, "data", "config.json")
	remove := scen.createConfigFile(&cfg, fname)
	defer remove()

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

func (scen *Scenario) WithChatSqliteDatabase(handler func(db *gorm.DB)) {
	dbloc := filepath.Join(scen.Base, "tokopedia_chat_test.db")
	db := datasource.NewSqliteDatabase(dbloc)

	defer func() {
		sqldb, _ := db.DB()
		sqldb.Close()
		os.Remove(dbloc)
	}()

	db.AutoMigrate(
		&model.AccountData{},
		&model.Group{},
		&model.Account{},
		&model.OrderItem{},
		&model.OrderSheet{},
		&model.Order{},
	)

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

type FilterTextModifier func(cfg *legacy_source.FilterText) error
type FilterTextHandler func(cfg *legacy_source.FilterText)

func (scen *Scenario) WithFilterText(modifier FilterTextModifier, handler FilterTextHandler) {
	cfg := legacy_source.FilterText{}

	err := modifier(&cfg)
	assert.Nil(scen.t, err)

	fname := filepath.Join(scen.Base, "data", "filter_text_config")
	remove := scen.createConfigFile(&cfg, fname)
	defer remove()

	handler(&cfg)
}

type GrabBasicModifier func(cfg *legacy.GrabBasic) error
type GrabBasicHandler func(cfg *legacy.GrabBasic)

func (scen *Scenario) WithGrabBasic(modifier GrabBasicModifier, handler GrabBasicHandler) {
	cfg := legacy.GrabBasic{}

	err := modifier(&cfg)
	assert.Nil(scen.t, err)

	fname := filepath.Join(scen.Base, "data", "grab_config")
	remove := scen.createConfigFile(&cfg, fname)
	defer remove()

	handler(&cfg)
}

type GrabTokopediaModifier func(cfg *legacy.GrabTokopedia) error
type GrabTokopediaHandler func(cfg *legacy.GrabTokopedia)

func (scen *Scenario) WithGrabTokopedia(modifier GrabTokopediaModifier, handler GrabTokopediaHandler) {
	cfg := legacy.GrabTokopedia{}

	err := modifier(&cfg)
	assert.Nil(scen.t, err)

	fname := filepath.Join(scen.Base, "data", "grab_tokopedia_config")
	remove := scen.createConfigFile(&cfg, fname)
	defer remove()

	handler(&cfg)
}

type MarkupConfigModifier func(cfg *legacy.LegacyMarkupConfig) error
type MarkupConfigHandler func(cfg *legacy.LegacyMarkupConfig)

func (scen *Scenario) WithMarkupConfig(modifier MarkupConfigModifier, handler MarkupConfigHandler) {
	cfg := legacy.LegacyMarkupConfig{}

	err := modifier(&cfg)
	assert.Nil(scen.t, err)

	fname := filepath.Join(scen.Base, "data", "markup_data")
	remove := scen.createConfigFile(&cfg, fname)
	defer remove()

	handler(&cfg)
}
