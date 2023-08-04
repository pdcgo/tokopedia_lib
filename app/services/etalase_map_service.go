package services

import (
	"github.com/pdcgo/tokopedia_lib/app/config"
	"github.com/pdcgo/tokopedia_lib/app/shopee/shopee_repo"
	"gorm.io/gorm"
)

type EtalaseMapItem struct {
	ID          uint   `gorm:"primarykey"`
	EtalaseName string `gorm:"index:etalase_map_unique,unique" json:"etalase_name"`
	CategoryID  int    `gorm:"index:etalase_map_unique,unique" json:"category_id"`
}

type EtalaseMapService struct {
	db         *gorm.DB
	shopeePAgg shopee_repo.ProductAggregate
	mapconfig  *config.ShopeeMapper
}

func NewEtalaseMapService(
	db *gorm.DB,
	shopeePAgg shopee_repo.ProductAggregate,
) *EtalaseMapService {

	db.AutoMigrate(
		&EtalaseMapItem{},
	)

	service := EtalaseMapService{
		db:         db,
		shopeePAgg: shopeePAgg,
		mapconfig:  config.NewShopeeMapper(db),
	}

	return &service
}

type ShopeeEtalaseMapItem struct {
	ShopeeID              int64 `json:"shopee_id"`
	TokopediaID           int   `json:"tokpedia_id"`
	Count                 int   `json:"product_count"`
	CategoryNoMapping     bool  `json:"category_no_mapping"`
	ShopeeCategoryName    []string
	TokopediaCategoryName []string
	EtalaseName           string
}

func (service *EtalaseMapService) GetListMap(namespace string) ([]*ShopeeEtalaseMapItem, error) {
	hasil := []*ShopeeEtalaseMapItem{}

	err := service.shopeePAgg.IterCategory(namespace, func(shopeeID int64, count int, name []string) error {
		item := ShopeeEtalaseMapItem{
			ShopeeID:           shopeeID,
			Count:              count,
			ShopeeCategoryName: name,
		}

		mapitem, err := service.mapconfig.GetTokopediaID(shopeeID)
		if err != nil {
			item.CategoryNoMapping = true
			return nil
		}

		item.TokopediaID = mapitem.TokopediaID

		etamap, _ := service.GetEtalase(item.TokopediaID)
		item.EtalaseName = etamap.EtalaseName

		hasil = append(hasil, &item)

		return nil
	})

	return hasil, err
}

func (service *EtalaseMapService) GetEtalase(catID int) (*EtalaseMapItem, error) {
	mapitem := EtalaseMapItem{
		CategoryID: catID,
	}

	err := service.db.Where(&mapitem).First(&mapitem).Error
	return &mapitem, err
}

func (service *EtalaseMapService) DeleteEtalase(name string) error {
	return service.db.Where(&EtalaseMapItem{
		EtalaseName: name,
	}).Delete(&EtalaseMapItem{}).Error
}

type EtalasePayload struct {
	Etalase string `json:"etalase"`
	CatIDs  []int  `json:"cat_ids"`
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

func (service *EtalaseMapService) UpdateBulkMap(payload []*EtalaseMapItem) error {

	for _, item := range payload {
		err := service.db.Transaction(func(tx *gorm.DB) error {
			etalaseName := item.EtalaseName

			tx.Where(&EtalaseMapItem{CategoryID: item.CategoryID}).First(item)
			item.EtalaseName = etalaseName
			return tx.Save(item).Error
		})
		if err != nil {
			return err
		}

	}

	return nil
}
