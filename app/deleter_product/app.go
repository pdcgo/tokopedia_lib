package deleter_product

import (
	"errors"
	"log"
	"sync"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/rs/zerolog"
)

type DeleteRunner struct {
	limitGuard chan int
	Config     *TokopediaDeleteConfig
}

func NewDeleteRunner(cfg *TokopediaDeleteConfig) *DeleteRunner {
	deleter := DeleteRunner{
		limitGuard: make(chan int, cfg.LimitConcurent),
		Config:     cfg,
	}
	return &deleter
}

var ErrDeleteLimitExcedeed = errors.New("delete limit excedeed")

func (runner *DeleteRunner) RunDeleteAkun(akun *AkunDeleteItem, reports chan *DeleteReportItem) error {
	driver, err := tokopedia_lib.NewDriverAccount(akun.Username, akun.Password, akun.Secret)
	if err != nil {
		return err
	}

	sapi, saveSession, err := driver.CreateApi()
	if err != nil {
		pdc_common.ReportErrorCustom(err, func(event *zerolog.Event) *zerolog.Event {
			return event.Str("err_from", "create api")
		})
		return err
	}

	defer saveSession()

	status := runner.Config.StatusProduct
	username := sapi.AuthenticatedData.User.Email

	filterhandler := runner.Config.GenerateFilter()
	count := 0

	// running yang pelanggaran
	if status == model.ViolationStatus {
		runner.Config.LimitProduct = 1000000
		filterhandler = func(product *model.SellerProductItem) (bool, string) {
			return true, ""
		}
	}

	// running yang Inactive
	if status == model.InActiveStatus {
		runner.Config.LimitProduct = 1000000
		filterhandler = func(product *model.SellerProductItem) (bool, string) {
			return true, ""
		}
	}

	err = IterateProduct(sapi, &IterateFilter{
		CategoryID: runner.Config.CategoryID,
		PageSize:   50,
		Status:     runner.Config.StatusProduct,
	}, func(page int, product *model.SellerProductItem, delete func() int) error {
		if product.Status == model.DeletedStatus {
			return nil
		}

		cek, _ := filterhandler(product)
		if cek {
			count = delete()

			reports <- &DeleteReportItem{
				Username: username,
				Judul:    product.Name,
				Url:      product.URL,
				Status:   product.Status,
			}
			log.Println(username, count, "/", runner.Config.LimitProduct, "deleted ", product.Name)
		}

		if count == runner.Config.LimitProduct {
			return ErrDeleteLimitExcedeed
		}
		return nil
	})

	if errors.Is(err, ErrDeleteLimitExcedeed) {
		log.Println(username, "delete selesai ...")
		return nil
	}

	if err != nil {
		pdc_common.ReportErrorCustom(err, func(event *zerolog.Event) *zerolog.Event {
			return event.Str("err_from", "iterate product")
		})
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
