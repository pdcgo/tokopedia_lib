package deleter_product

import (
	"errors"
	"log"
	"sync"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/api"
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

func (runner *DeleteRunner) RunDeleteViolation(sapi *api.TokopediaApi) error {

	queryFilter := []model.Filter{
		{
			ID:    "status",
			Value: []string{string(runner.Config.StatusProduct)},
		},
	}

	err := IterateProduct(sapi, func(page int, product *model.SellerProductItem, delete func() int) error {
		count := delete()
		log.Println(sapi.AuthenticatedData.User.Email, count, "/", runner.Config.LimitProduct, "deleted ", product.Name)

		return nil
	}, queryFilter...)

	return err
}

func (runner *DeleteRunner) RunDeleteAkun(akun *AkunDeleteItem, reports chan *DeleteReportItem) error {
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

	queryFilter := []model.Filter{}
	if runner.Config.StatusProduct != "" {
		queryFilter = append(queryFilter, model.Filter{
			ID:    "status",
			Value: []string{string(runner.Config.StatusProduct)},
		})
	}

	// log.Println(queryFilter)
	// time.Sleep(time.Hour)

	// running yang pelanggaran
	if runner.Config.StatusProduct == model.ViolationStatus {
		runner.Config.LimitProduct = 1000000
		filterhandler = func(product *model.SellerProductItem) (bool, string) {
			return true, ""
		}
	}

	// running yang Inactive
	if runner.Config.StatusProduct == model.InActiveStatus {
		runner.Config.LimitProduct = 1000000
		filterhandler = func(product *model.SellerProductItem) (bool, string) {
			return true, ""
		}
	}

	if runner.Config.CategoryID != "" {
		queryFilter = append(queryFilter, model.Filter{
			ID:    "category",
			Value: []string{runner.Config.CategoryID},
		})
	}

	err = IterateProduct(sapi, func(page int, product *model.SellerProductItem, delete func() int) error {
		cek, _ := filterhandler(product)
		if cek {
			count = delete()

			reports <- &DeleteReportItem{
				Username: sapi.AuthenticatedData.User.Email,
				Judul:    product.Name,
				Url:      product.URL,
				Status:   product.Status,
			}
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

func (runner *DeleteRunner) Run(fname string) {

	var wg sync.WaitGroup

	reports := []*DeleteReportItem{}
	defer func() {
		SaveReport(fname, reports)
	}()

	reportchan := NewReport()
	go func() {
		for report := range reportchan {
			reports = append(reports, report)
		}
	}()

	for _, ak := range runner.Config.Akuns {
		runner.limitGuard <- 1

		akun := ak
		wg.Add(1)
		go func() {
			defer func() {
				<-runner.limitGuard
				wg.Done()
			}()

			err := runner.RunDeleteAkun(akun, reportchan)
			if err != nil {
				pdc_common.ReportError(err)
			}
		}()
	}

	wg.Wait()
}
