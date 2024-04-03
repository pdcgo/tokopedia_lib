package toko_libur

import (
	"fmt"
	"log"
	"sync"

	"github.com/pdcgo/tokopedia_lib/lib/api"
)

type Application struct {
	Config *TokopediaTokoLiburConfig
	Report *Report
}

func NewApplication(config *TokopediaTokoLiburConfig, report *Report) *Application {
	return &Application{
		Config: config,
		Report: report,
	}
}

var statuses = map[bool]string{
	false: "Toko Buka",
	true:  "Toko Libur",
}

func (app *Application) ProcessLibur(item *ReportItem, tapi *api.TokopediaApi) error {
	log.Printf("[ %s ] processing...", item.Username)

	info, err := tapi.ShopInfoByID()
	if err != nil {
		return err
	}

	tokolibur := info.Data.ShopInfoByID.Result[0].ClosedInfo.Detail.Status > 0
	editlibur := app.Config.Libur

	item.StatusBefore = statuses[tokolibur]
	item.StatusAfter = statuses[editlibur]
	if tokolibur == editlibur {
		item.Message = "Tidak ada perubahan"
		return nil
	}

	input := app.Config.CreateCloseShopInput()
	closedInfo, err := tapi.CloseShopSchedule(input)
	if closedInfo.Data != nil {
		item.Message = closedInfo.Data.CloseShopSchedule.Message
		if !closedInfo.Data.CloseShopSchedule.Success {
			return fmt.Errorf("gagal update status toko")
		}
	}

	return err
}

func (app *Application) Run() error {

	var wg sync.WaitGroup
	var limit = make(chan int, app.Config.Limit)

	app.Report.Iterate(func(item *ReportItem, updateitem func(handler func(item *ReportItem) error) error) error {
		driver, err := item.CreateDriver()
		if err != nil {
			updateitem(func(item *ReportItem) error { return item.ReportError(err) })
			return nil
		}

		tapi, saveSession, err := driver.CreateApi()
		if err != nil {
			updateitem(func(item *ReportItem) error { return item.ReportError(err) })
			return nil
		}

		wg.Add(1)
		limit <- 1
		go updateitem(func(item *ReportItem) error {
			defer saveSession()
			defer wg.Done()
			defer func() {
				<-limit
			}()

			err := app.ProcessLibur(item, tapi)
			if err != nil {
				item.ReportError(err)
			}
			return nil
		})
		return nil
	})

	wg.Wait()
	return nil
}
