package deleter_product

import (
	"errors"
	"log"
	"sync"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/model"
)

type DeleteRunner struct {
	limitGuard chan int
	Config     *DeleteConfig
}

func NewDeleteRunner(cfg *DeleteConfig) *DeleteRunner {
	deleter := DeleteRunner{
		limitGuard: make(chan int, cfg.LimitConcurent),
		Config:     cfg,
	}
	return &deleter
}

var ErrDeleteLimitExcedeed = errors.New("delete limit excedeed")

func (runner *DeleteRunner) RunDeleteAkun(akun *AkunDeleteItem) error {
	driver, err := tokopedia_lib.NewDriverAccount(akun.Username, akun.Password, akun.Secret)

	if err != nil {
		return err
	}

	sapi, saveSession, err := driver.CreateApi()

	if err != nil {
		return err
	}

	defer saveSession()
	filterhandler := runner.Config.GenerateFilter()
	count := 0

	queryFilter := []model.Filter{
		{
			ID:    "status",
			Value: []string{string(runner.Config.StatusProduct)},
		},
	}
	if runner.Config.CategoryID != "" {
		queryFilter = append(queryFilter, model.Filter{
			ID:    "category",
			Value: []string{runner.Config.CategoryID},
		})
	}

	err = IterateProduct(sapi, func(page int, product *model.SellerProductItem, delete func() int) error {
		if filterhandler(product) {
			count = delete()
			log.Println(sapi.AuthenticatedData.User.Email, count, "/", runner.Config.LimitProduct, "deleted ", product.Name)
		}

		if count == runner.Config.LimitProduct {
			return ErrDeleteLimitExcedeed
		}
		return nil
	}, queryFilter...)

	if errors.Is(err, ErrDeleteLimitExcedeed) {
		log.Println(sapi.AuthenticatedData.User.Email, "delete selesai ...")
		return nil
	}

	return err
}

func (runner *DeleteRunner) Run() {

	var wg sync.WaitGroup

	for _, ak := range runner.Config.Akuns {
		runner.limitGuard <- 1

		akun := ak
		wg.Add(1)
		go func() {
			defer func() {
				<-runner.limitGuard
				wg.Done()
			}()

			err := runner.RunDeleteAkun(akun)
			if err != nil {
				pdc_common.ReportError(err)
			}
		}()
	}

	wg.Wait()
}
