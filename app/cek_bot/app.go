package cek_bot

import (
	_ "embed"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pdcgo/common_conf/auth"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/pdcgo/tokopedia_lib/lib/report"
	"github.com/rs/zerolog"
	"github.com/urfave/cli/v2"
)

var concurent = make(chan int, 50)
var Waitallakun sync.WaitGroup
var loginMutex sync.Mutex

func LisensiLogin(botID int) bool {
	development := os.Getenv("DEVEL_AUTH") != ""

	if development {
		return true
	}

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
	// fname := "data/config.json"
	// pdc_common.SetConfig(fname, config.Version, "golang_tokopedia_checkbot", config.Cred)
	// pdc_common.InitializeLogger()

	if !LisensiLogin(2) {
		time.Sleep(time.Minute)
		panic(errors.New("gagal login"))
	}
}

func Cekbot(driver *report.CekReport) {
	loginMutex.Lock()
	defer loginMutex.Unlock()
	log.Println("login", driver.Username)
	apiclient, saveSession, err := driver.CreateApi()
	log.Println("login finish", driver.Username)
	if err != nil {
		pdc_common.ReportErrorCustom(err, func(event *zerolog.Event) *zerolog.Event {
			return event.Str("username", driver.Username)
		})
		Waitallakun.Done()
		return
	}

	driver.ShopName = apiclient.AuthenticatedData.UserShopInfo.Info.ShopName

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
			hasil, err := apiclient.ProductListMeta()
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
			hasil, err := apiclient.GetShopScoreLevel()
			if err != nil {
				pdc_common.ReportError(err)
			}

			driver.ShopScore = hasil.Data.ShopScoreLevel.Result.ShopScore

		}()

		// getting notification
		waitall.Add(1)
		go func() {
			defer waitall.Done()
			hasil, err := apiclient.NotificationCounter()
			if err != nil {
				pdc_common.ReportError(err)
			}

			driver.UreadChat = hasil.Data.Notifications.Chat.UnreadsSeller
			driver.NewOrder = hasil.Data.Notifications.SellerOrderStatus.NewOrder

		}()

		// getting merchant status
		waitall.Add(1)
		go func() {
			defer waitall.Done()
			hasil, err := apiclient.GoldGetPMOSStatus()
			if err != nil {
				pdc_common.ReportError(err)
			}

			driver.ExtendStatus = hasil.Data.GoldGetPMOSStatus.Data.PowerMerchant.AutoExtend.Status

			switch hasil.Data.GoldGetPMOSStatus.Data.PowerMerchant.Status {
			case "active":
				driver.PmStatus = "aktif"
			case "idde":
				driver.PmStatus = "tidak aktif"
			case "inactive":
				driver.PmStatus = "moderasi"
			}

		}()

		// product add rule
		waitall.Add(1)
		go func() {
			defer waitall.Done()
			hasil, err := apiclient.GetProductAddRule()
			if err != nil {
				pdc_common.ReportError(err)
			}

			driver.ProductTotal = hasil.Data.ProductAddRule.Data.Eligible.TotalProduct
			driver.ProductLimit = hasil.Data.ProductAddRule.Data.Eligible.Limit
		}()

		// getting ShopInfo
		waitall.Add(1)
		go func() {
			defer waitall.Done()
			hasil, err := apiclient.ShopInfoByID()
			if err != nil {
				pdc_common.ReportError(err)
			}

			switch hasil.Data.ShopInfoByID.Result[0].StatusInfo.StatusName {
			case "Moderated Permanently":
				driver.Status = "Moderasi Permanen"
			case "Moderated":
				driver.Status = "Moderasi Sementara"
			}

		}()

		// getting penalty
		waitall.Add(1)
		go func() {
			defer waitall.Done()
			enddate := time.Now()
			stdate := enddate.AddDate(0, -2, 0)

			shopID := strconv.FormatInt(apiclient.AuthenticatedData.UserShopInfo.Info.ShopID, 10)

			hasil, err := apiclient.ShopScorePenaltySummary(&api.ShopScorePenaltySummaryVar{
				StartDate: stdate,
				EndDate:   enddate,
				ShopID:    shopID,
				Source:    "icarus",
			})

			if err != nil {
				pdc_common.ReportError(err)
				return
			}
			penalty := strconv.Itoa(hasil.Data.ShopScorePenaltySummary.Result.Penalty)
			penaltyAmount := strconv.Itoa(hasil.Data.ShopScorePenaltySummary.Result.PenaltyAmount)
			driver.Penalty = penalty
			driver.PenaltyAmount = penaltyAmount
		}()

		// getting rekening bank
		waitall.Add(1)
		go func() {
			defer waitall.Done()

			backAccount, err := apiclient.GetBankAccount()
			if err != nil {
				pdc_common.ReportError(err)
				return
			}

			reklist := []string{}
			for _, bank := range backAccount.Data.GetBankAccount.Data.BankAccounts {
				rekstr := fmt.Sprintf("%s(%s)", bank.AccNumber, bank.AccName)
				reklist = append(reklist, rekstr)
			}
			driver.RekeningBank = strings.Join(reklist, "#")
		}()

		waitall.Wait()
		Waitallakun.Done()
		log.Println(driver.Username, "Berhasil Checking..")

	}()

}

func runCheckAkun(cCtx *cli.Context) error {
	setupPdcLogger()
	// proxy := tokpedproxy.NewInspectProxy("localhost:8082", context.Background())
	// go proxy.RunProxy()

	// apiSession, saveSession := scenario.GetTokopediaApiClient()
	// defer saveSession()

	// hasil, _ := api.IsAutheticated()
	// log.Println(hasil)
	// time.Sleep(time.Hour)

	// driver.Proxy = proxy.Addr

	fname := cCtx.String("fname")

	akuns, save, err := report.NewCekReport(fname)
	defer save()
	if err != nil {
		pdc_common.ReportError(err)
		time.Sleep(time.Minute)
		return nil
	}

	for _, driver := range akuns {
		Waitallakun.Add(1)
		go Cekbot(driver)

	}

	Waitallakun.Wait()
	log.Println("cekbot selesai..")

	return nil
}

func CreateCheckbotCommand() *cli.Command {
	command := cli.Command{
		Name:    "cekbot",
		Aliases: []string{"cb"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "base",
				Aliases: []string{"b"},
				Value:   "../",
			},
			&cli.StringFlag{
				Name:    "fname",
				Aliases: []string{"f"},
				Value:   "cekbot.txt",
			},
		},
		Action: runCheckAkun,
	}
	return &command
}
