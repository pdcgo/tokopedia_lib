package services

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

type EtalaseMapItem struct {
	ID          uint   `gorm:"primarykey"`
	EtalaseName string `gorm:"index:etalase_map_unique,unique" json:"etalase_name"`
	CategoryID  int    `gorm:"index:etalase_map_unique,unique" json:"category_id"`
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

func (service *EtalaseMapService) UpdateBulkMap(payload []*EtalaseMapItem) []error {
	hasil := make([]error, len(payload))

	for ind, item := range payload {
		err := service.db.Transaction(func(tx *gorm.DB) error {
			etalaseName := item.EtalaseName

			tx.Where(&EtalaseMapItem{CategoryID: item.CategoryID}).First(item)
			item.EtalaseName = etalaseName
			return tx.Save(item).Error
		})

		hasil[ind] = err
	}

	return hasil
}

func (service *EtalaseMapService) List() ([]*EtalaseMapItem, error) {
	hasil := []*EtalaseMapItem{}
	err := service.db.Find(&hasil).Error
	return hasil, err
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
