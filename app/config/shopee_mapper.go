package config

import (
	"errors"

	"gorm.io/gorm"
)

type ShopeeMapperConfig struct {
	UseMapper bool `json:"use_mapper"`
}

func (cfg *ShopeeMapperConfig) KeyConfig() string {
	return "UPLOAD_CONFIG_MAPPER"
}

type ShopeeMapItem struct {
	ShopeeID    int64 `gorm:"primaryKey;autoIncrement:false"  json:"shopee_id"`
	TokopediaID int   `json:"tokopedia_id"`
}

type ShopeeMapper struct {
	db *gorm.DB
}

func NewShopeeMapper(db *gorm.DB) *ShopeeMapper {
	return &ShopeeMapper{
		db: db,
	}
}

func (maper *ShopeeMapper) GetTokopediaID(shopeID int64) (*ShopeeMapItem, error) {
	itemap := ShopeeMapItem{
		ShopeeID: shopeID,
	}
	err := maper.db.Where(&itemap).First(&itemap).Error
	return &itemap, err
}

var ErrMapNotFound = errors.New("mapping di db tidak ada")

func (maper *ShopeeMapper) GetShopeeID(tokopediaID int) (*ShopeeMapItem, error) {
	itemap := ShopeeMapItem{}

	err := maper.db.Where(&ShopeeMapItem{
		TokopediaID: tokopediaID,
	}).First(&itemap).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &itemap, ErrMapNotFound
	}

	return &itemap, err
}
