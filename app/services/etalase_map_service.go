package services

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

type EtalaseMapItem struct {
	ID          uint   `gorm:"primarykey"`
	EtalaseName string `gorm:"index:etalase_map_unique,unique"`
	CategoryID  int    `gorm:"index:etalase_map_unique,unique"`
}

type EtalaseMapService struct {
	db *gorm.DB
}

func NewEtalaseMapService(
	db *gorm.DB,
) *EtalaseMapService {

	db.AutoMigrate(
		&EtalaseMapItem{},
	)

	service := EtalaseMapService{
		db: db,
	}

	return &service
}

type EtalasePayload struct {
	Etalase string `json:"etalase"`
	CatIDs  []int  `json:"cat_ids"`
}

func (service *EtalaseMapService) GetEtalase(catID int) (*EtalaseMapItem, error) {
	mapitem := EtalaseMapItem{
		CategoryID: catID,
	}

	err := service.db.Where(&mapitem).First(&mapitem).Error
	log.Println(mapitem)
	return &mapitem, err

}

func (service *EtalaseMapService) DeleteEtalase(name string) error {
	return service.db.Where(&EtalaseMapItem{
		EtalaseName: name,
	}).Delete(&EtalaseMapItem{}).Error
}

func (service *EtalaseMapService) ListEtalase() ([]*EtalasePayload, error) {
	mapsitems := []*EtalaseMapItem{}
	err := service.db.Model(&EtalaseMapItem{}).Find(&mapsitems).Error
	if err != nil {
		return []*EtalasePayload{}, err
	}

	etalasemap := map[string][]int{}

	for _, item := range mapsitems {
		etalasemap[item.EtalaseName] = append(etalasemap[item.EtalaseName], item.CategoryID)
	}

	hasil := []*EtalasePayload{}
	for key, item := range etalasemap {
		hasil = append(hasil, &EtalasePayload{
			Etalase: key,
			CatIDs:  item,
		})
	}

	return hasil, nil
}

func (service *EtalaseMapService) AddMap(payload *EtalasePayload) error {
	return service.db.Transaction(func(tx *gorm.DB) error {
		for _, catID := range payload.CatIDs {
			etalasename := payload.Etalase
			mapitem := EtalaseMapItem{
				EtalaseName: etalasename,
				CategoryID:  catID,
			}

			err := tx.Where(&mapitem).First(&mapitem).Error

			if errors.Is(err, gorm.ErrRecordNotFound) {
				err = tx.Save(&mapitem).Error
			}

			if err != nil {
				return err
			}
		}

		return nil
	})
}
