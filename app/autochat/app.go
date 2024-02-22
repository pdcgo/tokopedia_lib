package autochat

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/pdcgo/common_conf/pdc_application"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
)

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

	pubapi, err := api_public.NewTokopediaApiPublic()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	for akun := range senderchan {

		wg.Add(1)
		go func(sender *AutochatSender) {
			defer wg.Done()

			limitchan <- true
			defer func() {
				<-limitchan
			}()

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			autosend := NewAutochatSend(sender, pubapi)
			for _, shop := range shops {
				limit := app.config.LimitMessageSend.Get()
				autosend.Run(ctx, limit, shop)
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

	var wg sync.WaitGroup
	for akun := range senderchan {

		wg.Add(1)
		go func(sender *AutochatSender) {
			defer wg.Done()

			ctx, cancel := context.WithCancel(parentCtx)
			defer cancel()

			autoreply := NewAutochatReply(sender)
			autoreply.Run(ctx)
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
		log.Printf("mode %d tidak ditemukan", mode)
		return nil
	}
}
