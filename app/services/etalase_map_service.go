package services

import "gorm.io/gorm"

type EtalaseMapService struct {
	db *gorm.DB
}

func NewEtalaseMapService(
	db *gorm.DB,
) *EtalaseMapService {
	service := EtalaseMapService{
		db: db,
	}

	return &service
}
