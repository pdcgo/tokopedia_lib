package autochat

import (
	"context"
	"encoding/csv"
	"errors"
	"io"
	"log"
	"sync"

	"github.com/gocarina/gocsv"
	"github.com/pdcgo/common_conf/pdc_application"
	"github.com/pdcgo/tokopedia_lib/app/autochat/report"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
)

func init() {
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.FieldsPerRecord = -1
		r.LazyQuotes = true
		return r
	})
}

type AutochatMode int

const (
	AUTOCHAT_SEND  = 1
	AUTOCHAT_REPLY = 2
)

type Application struct {
	base      pdc_application.BaseApplication
	message   *AutochatMessage
	config    *AutochatConfig
	akundata  *AkunData
	shopdata  *ShopData
	limitchan chan bool
}

func NewApplication(
	base pdc_application.BaseApplication,
	message *AutochatMessage,
	config *AutochatConfig,
	akundata *AkunData,
	shopdata *ShopData,
) *Application {

	return &Application{
		base:      base,
		message:   message,
		config:    config,
		akundata:  akundata,
		shopdata:  shopdata,
		limitchan: make(chan bool, config.Concurrent),
	}
}

func (app *Application) autoSendOneToMulti(wg *sync.WaitGroup, autosend *AutochatSend, updateItem report.EditAutosendReportItem) {

	defer wg.Done()

	app.limitchan <- true
	defer func() {
		<-app.limitchan
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app.shopdata.Iterate(func(shop *Shop) error {

		updateItem(func(item *report.AutosendReportItem) error {
			item.SellerChatProcessed++
			return nil
		})
		defer func() {
			updateItem(func(item *report.AutosendReportItem) error {
				item.SellerChatDone++
				return nil
			})
		}()

		limit := app.config.LimitMessageSend.Get()
		err := autosend.Run(ctx, limit, shop.ShopName)
		if err != nil {
			updateItem(func(item *report.AutosendReportItem) error {
				item.Error = err.Error()
				return nil
			})
		}

		return nil
	})
}

func (app *Application) autoSendOneToOne(wg *sync.WaitGroup, autosend *AutochatSend, updateItem report.EditAutosendReportItem) {

	defer wg.Done()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app.shopdata.Iterate(func(shop *Shop) error {

		updateItem(func(item *report.AutosendReportItem) error {
			item.SellerChatProcessed++
			return nil
		})
		defer func() {
			updateItem(func(item *report.AutosendReportItem) error {
				item.SellerChatDone++
				return nil
			})
		}()

		for {
			shop, err := app.shopdata.Get()
			if errors.Is(err, ErrNoShopMore) || errors.Is(err, ErrNoShop) {
				break
			}

			func() {
				app.limitchan <- true
				defer func() {
					shop.Status = SHOP_STATUS_DONE
					app.shopdata.Save()
					<-app.limitchan
				}()

				limit := app.config.LimitMessageSend.Get()
				err := autosend.Run(ctx, limit, shop.ShopName)
				if err != nil {
					updateItem(func(item *report.AutosendReportItem) error {
						item.Error = err.Error()
						return nil
					})
				}
			}()
		}

		return nil
	})
}

func (app *Application) RunAutoSend() error {

	shoplen := len(app.shopdata.Data)
	pubapi, err := api_public.NewTokopediaApiPublic()
	if err != nil {
		return err
	}

	autoreport := report.NewAutosendReport(app.base)

	var wg sync.WaitGroup
	app.akundata.IterateAkunSender(app.message, func(akun *Akun, sender *AutochatSender) error {

		_, updateItem := autoreport.CreateItem(sender.GetName(), shoplen)

		wg.Add(1)
		autosend := NewAutochatSend(sender, pubapi, updateItem, app.config.Autosend)

		go func() {

			defer func() {
				akun.Status = AKUN_STATUS_DONE
				app.akundata.Save()
			}()

			if app.config.Autosend.OneToMulti {
				app.autoSendOneToMulti(&wg, autosend, updateItem)
			} else {
				app.autoSendOneToOne(&wg, autosend, updateItem)
			}
		}()

		return nil
	})

	wg.Wait()
	return nil
}

func (app *Application) RunAutoReply() error {

	parentCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	autoreport := report.NewAutoreplyReport(app.base)

	var wg sync.WaitGroup
	app.akundata.IterateAkunSender(app.message, func(akun *Akun, sender *AutochatSender) error {

		wg.Add(1)
		go func(sender *AutochatSender) {
			defer wg.Done()

			ctx, cancel := context.WithCancel(parentCtx)
			defer cancel()

			autoreply := NewAutochatReply(sender)
			autoreply.Run(ctx, autoreport)
		}(sender)

		return nil
	})

	wg.Wait()
	return nil
}

func (app *Application) Run(mode AutochatMode) error {

	switch mode {

	case AUTOCHAT_SEND:
		return app.RunAutoSend()

	case AUTOCHAT_REPLY:
		return app.RunAutoReply()

	default:
		log.Printf("\nmode %d tidak ditemukan", mode)
		return nil
	}
}
