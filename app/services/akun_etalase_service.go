package services

import (
	"sync"

	"github.com/pdcgo/tokopedia_lib/lib/api"
)

type AkunEtalaseService struct {
	sync.Mutex
	Etalase    map[string]*api.ShopShowcaseResult
	api        *api.TokopediaApi
	mapservice *EtalaseMapService
}

func NewAkunEtalaseService(
	apiclient *api.TokopediaApi,
	mapservice *EtalaseMapService,
) *AkunEtalaseService {
	service := AkunEtalaseService{
		api:        apiclient,
		mapservice: mapservice,
		Etalase:    make(map[string]*api.ShopShowcaseResult),
	}

	return &service
}

func (service *AkunEtalaseService) GetEtalase(catID int) (*api.ShopShowcaseResult, error) {
	mapitem, err := service.mapservice.GetEtalase(catID)
	if err != nil {
		return nil, err
	}

	return service.GetSellerEtalase(mapitem.EtalaseName)
}

func (service *AkunEtalaseService) GetSellerEtalase(name string) (*api.ShopShowcaseResult, error) {
	service.Lock()
	showcase := service.Etalase[name]
	service.Unlock()

	if showcase == nil {
		_, err := service.api.AddShopShowcase(name)
		if err != nil {
			return showcase, err
		}
		err = service.RefreshShowCase()
		showcase := service.Etalase[name]
		return showcase, err
	}

	return showcase, nil

}

func (service *AkunEtalaseService) RefreshShowCase() error {
	service.Lock()
	defer service.Unlock()

	showcase, err := service.api.ShopShowcase()

	if err != nil {
		return err
	}

	for _, res := range showcase.Data.ShopShowcases.Result {
		service.Etalase[res.Name] = res
	}

	return nil

}
