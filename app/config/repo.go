package config

import (
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type RowConfigItem struct {
	Key   string `gorm:"primaryKey;autoIncrement:false"`
	Value string
}

type ConfigRepo struct {
	db *gorm.DB
}

func NewConfigRepo(db *gorm.DB) *ConfigRepo {
	db.AutoMigrate(&RowConfigItem{})
	return &ConfigRepo{
		db: db,
	}
}

type Config interface {
	KeyConfig() string
}

func (repo *ConfigRepo) GetConfig(data Config) error {
	key := data.KeyConfig()

	var item RowConfigItem

	err := repo.db.Model(&RowConfigItem{}).Find(&RowConfigItem{
		Key: key,
	}).First(&item).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	return json.Unmarshal([]byte(item.Value), data)
}

func (repo *ConfigRepo) Save(data Config) error {

	databyte, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	key := data.KeyConfig()
	item := RowConfigItem{
		Key:   key,
		Value: string(databyte),
	}

	err = repo.db.Save(&item).Error
	return err
}
