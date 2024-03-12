package autochat

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"os"
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
	base    pdc_application.BaseApplication
	message *AutochatMessage
	config  *AutochatConfig
}

func NewApplication(
	base pdc_application.BaseApplication,
	message *AutochatMessage,
	config *AutochatConfig,
) *Application {

	return &Application{
		base:    base,
		message: message,
		config:  config,
	}
}

func (app *Application) RunAutoSend() error {

	limitchan := make(chan bool, app.config.Concurrent)
	senderchan, err := app.IterateAkunSender()
	if err != nil {
		return err
	}

	fname := app.base.Path(app.config.ShopLoc)
	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	shops, err := fileLineSplit(file)
	if err != nil {
		return err
	}
	shoplen := len(shops)

	pubapi, err := api_public.NewTokopediaApiPublic()
	if err != nil {
		return err
	}

	autoreport := report.NewAutosendReport(app.base)

	var wg sync.WaitGroup
	for akun := range senderchan {

		_, updateItem := autoreport.CreateItem(akun.GetName(), shoplen)
		updateItem(func(item *report.AutosendReportItem) error {
			item.SellerChatProcessed++
			return nil
		})

		wg.Add(1)
		go func(sender *AutochatSender) {
			defer wg.Done()

			limitchan <- true
			defer func() {
				updateItem(func(item *report.AutosendReportItem) error {
					item.SellerChatDone++
					return nil
				})
				<-limitchan
			}()

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			autosend := NewAutochatSend(sender, pubapi, updateItem, akun.config.Autosend)
			for _, shop := range shops {

				limit := app.config.LimitMessageSend.Get()
				err := autosend.Run(ctx, limit, shop)
				if err != nil {
					updateItem(func(item *report.AutosendReportItem) error {
						item.Error = err.Error()
						return nil
					})
				}
			}

		}(akun)
	}

	wg.Wait()
	return nil
}

func (app *Application) RunAutoReply() error {
	senderchan, err := app.IterateAkunSender()
	if err != nil {
		return err
	}

	parentCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	autoreport := report.NewAutoreplyReport(app.base)

	var wg sync.WaitGroup
	for akun := range senderchan {

		wg.Add(1)
		go func(sender *AutochatSender) {
			defer wg.Done()

			ctx, cancel := context.WithCancel(parentCtx)
			defer cancel()

			autoreply := NewAutochatReply(sender)
			autoreply.Run(ctx, autoreport)
		}(akun)
	}

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
