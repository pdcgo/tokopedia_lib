package services

import (
	"context"
	"errors"
	"regexp"
	"sync"
	"time"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/rs/zerolog"
	"github.com/sethvargo/go-retry"
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
	re := regexp.MustCompile(`[^A-Za-z0-9\s]+`)
	respace := regexp.MustCompile(`[\s]+`)

	name = re.ReplaceAllString(name, "")
	name = respace.ReplaceAllString(name, " ")

	var showcase *api.ShopShowcaseResult

	b := retry.NewFibonacci(time.Second)
	err := retry.Do(context.Background(), retry.WithMaxRetries(3, b), func(ctx context.Context) error {
		service.Lock()
		showcase = service.Etalase[name]
		service.Unlock()

		if showcase == nil {
			_, err := service.api.AddShopShowcase(name)
			if err != nil {
				return retry.RetryableError(err)
			}
			err = service.RefreshShowCase()
			if err != nil {
				return retry.RetryableError(err)
			}
			err = errors.New("showcase setelah refresh tidak ada")
			return retry.RetryableError(err)
		}

		return nil
	})

	if err != nil {
		pdc_common.ReportErrorCustom(err, func(event *zerolog.Event) *zerolog.Event {

			return event.
				Interface("map", service.Etalase).
				Str("etalase_name", name)
		})
	}

	return showcase, err

}

func (service *AkunEtalaseService) RefreshShowCase() error {
	service.Lock()
	defer service.Unlock()

	showcase, err := service.api.ShopShowcase()

	if err != nil {
		return err
	}

	for _, res := range showcase.Data.ShopShowcases.Result {
		showcase := res
		service.Etalase[res.Name] = showcase
	}

	return nil

}
