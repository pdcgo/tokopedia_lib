package main

import (
	_ "embed"
	"errors"
	"log"
	"os"
	"sync"
	"time"

	"github.com/pdcgo/common_conf/auth"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/report"
)

//go:embed ..\..\logger_credentials.json
var cred []byte

var devmode = os.Getenv("DEV_MODE") != ""
var Version = "development"
var LogName = "golang_tokopedia_cekbot"

var concurent = make(chan int, 50)
var waitallakun sync.WaitGroup
var loginMutex sync.Mutex

func LisensiLogin(botID int) bool {
	client := auth.NewAuthClient("https://pdcoke.com/v2/login")
	config := pdc_common.GetConfig()

	err := client.Login(config.Lisensi.Email, config.Lisensi.Pwd, botID, string(config.Version))
	if err != nil {
		log.Println("[", config.Lisensi.Email, "]", err)
		return false
	}

	return true
}

func setupPdcLogger() {
	fname := "config.yml"
	pdc_common.SetConfig(fname, Version, LogName, cred)
	pdc_common.InitializeLogger()

	if !LisensiLogin(7) {
		time.Sleep(time.Minute)
		panic(errors.New("gagal login"))
	}
}

func cekbot(driver *report.CekReport) {
	loginMutex.Lock()
	defer loginMutex.Unlock()
	api, saveSession, _ := driver.CreateApi()

	driver.ShopName = api.AuthenticatedData.UserShopInfo.Info.ShopName

	concurent <- 1
	go func() {
		defer func() {
			saveSession()
			<-concurent
		}()

		var waitall sync.WaitGroup

		// checking product
		waitall.Add(1)
		go func() {
			defer waitall.Done()
			hasil, err := api.ProductListMeta()
			if err != nil {
				pdc_common.ReportError(err)
			}

			for _, tab := range hasil.Data.ProductListMeta.Data.Tab {
				switch tab.ID {
				case "ACTIVE":
					driver.ProductActive = tab.Value
				case "INACTIVE":
					driver.ProductInActive = tab.Value
				case "VIOLATION":
					driver.ProductViolation = tab.Value
				}
			}

		}()

		// getting shopscore
		waitall.Add(1)
		go func() {
			defer waitall.Done()
			hasil, err := api.GetShopScoreLevel()
			if err != nil {
				pdc_common.ReportError(err)
			}

			driver.ShopScore = hasil[0].Data.ShopScoreLevel.Result.ShopScore

		}()

		// getting notification
		waitall.Add(1)
		go func() {
			defer waitall.Done()
			hasil, err := api.NotificationCounter()
			if err != nil {
				pdc_common.ReportError(err)
			}

			driver.UreadChat = hasil[0].Data.Notifications.Chat.UnreadsSeller
			driver.NewOrder = hasil[0].Data.Notifications.SellerOrderStatus.NewOrder

		}()

		// getting merchant status
		waitall.Add(1)
		go func() {
			defer waitall.Done()
			hasil, err := api.GoldGetPMOSStatus()
			if err != nil {
				pdc_common.ReportError(err)
			}

			driver.ExtendStatus = hasil[0].Data.GoldGetPMOSStatus.Data.PowerMerchant.AutoExtend.Status

			switch hasil[0].Data.GoldGetPMOSStatus.Data.PowerMerchant.Status {
			case "active":
				driver.PmStatus = "aktif"
			case "idde":
				driver.PmStatus = "tidak aktif"
			case "inactive":
				driver.PmStatus = "moderasi"
			}

		}()

		// getting ShopInfo
		waitall.Add(1)
		go func() {
			defer waitall.Done()
			hasil, err := api.ShopInfoByID()
			if err != nil {
				pdc_common.ReportError(err)
			}

			switch hasil[0].Data.ShopInfoByID.Result[0].StatusInfo.StatusName {
			case "Moderated Permanently":
				driver.Status = "Moderasi Permanen"
			case "Moderated":
				driver.Status = "Moderasi Sementara"
			}

		}()

		waitall.Wait()
		waitallakun.Done()
		log.Println(driver.Username, "Berhasil Checking..")

	}()

}

func main() {
	setupPdcLogger()
	// proxy := tokpedproxy.NewInspectProxy("localhost:8082", context.Background())
	// go proxy.RunProxy()

	// driver, _ := tokopedia_lib.NewDriverAccount("bethdunn892@outlook.com", "MZT2Zk8U", "FSR3 CTR2 5ZJX XIL5 TVK6 E72R HSRA U5GW")
	// driver.DevMode = devmode
	// api, saveSession, _ := driver.CreateApi()
	// defer saveSession()

	// hasil, _ := api.IsAutheticated()
	// log.Println(hasil)
	// time.Sleep(time.Hour)

	// driver.Proxy = proxy.Addr

	akuns, save, err := report.NewCekReport("cekbot.txt")
	defer save()
	if err != nil {
		pdc_common.ReportError(err)
		time.Sleep(time.Minute)
		return
	}

	for _, driver := range akuns {
		driver.DevMode = devmode
		waitallakun.Add(1)
		go cekbot(driver)

	}

	waitallakun.Wait()
	log.Println("cekbot selesai..")
}
