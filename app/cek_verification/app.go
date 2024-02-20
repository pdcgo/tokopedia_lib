package cek_verification

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/chromedp/cdproto/browser"
	"github.com/chromedp/chromedp"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/pdcgo/tokopedia_lib/lib/report"
	"github.com/sethvargo/go-retry"
	"github.com/urfave/cli/v2"
)

func CheckVerif(driver *report.CekVerifReport) error {
	return driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {

		driver.SellerLogin(dctx)

		timeout := time.After(time.Second * 30)
		success := make(chan int, 1)
		gagal := make(chan string, 1)
		process := make(chan string, 1)
		errorcek := make(chan error, 1)

		go func() {
			driver.RLock()
			defer driver.RUnlock()

			b := retry.NewFibonacci(time.Second)
			pathverif := `//*/div[contains(text(), "Terverifikasi")]`

			err := retry.Do(dctx.Ctx, retry.WithMaxRetries(3, b), func(ctx context.Context) error {
				return chromedp.Run(dctx.Ctx,
					browser.SetDownloadBehavior(browser.SetDownloadBehaviorBehaviorAllowAndName),
					chromedp.Navigate("https://mitra.tokopedia.com/user/akun-saya"),
					chromedp.WaitVisible(pathverif, chromedp.BySearch),
				)
			})

			if err != nil {
				errorcek <- err
			} else {
				success <- 1
			}
		}()

		go func() {
			var hasil string
			title := `//*/h5[@data-unify="Typography"]`
			chromedp.Run(dctx.Ctx,
				chromedp.WaitVisible(title, chromedp.BySearch),
				chromedp.Text(`//*/div[@aria-label="Gagal Verifikasi"]/div/span/p[@data-unify="Typography"]`, &hasil, chromedp.NodeVisible),
			)

			gagal <- hasil
		}()

		go func() {
			var hasil string
			contenttext := `//*/div[@class="content"]/p[@data-unify="Typography"]`
			chromedp.Run(dctx.Ctx,
				chromedp.WaitVisible(contenttext, chromedp.BySearch),
				chromedp.Text(contenttext, &hasil, chromedp.NodeVisible),
			)

			if strings.Contains(hasil, "dalam proses verifikasi") {
				process <- hasil
			}
		}()

		select {
		case <-timeout:
			driver.Status = "unknown"
			log.Println(driver.Username, "unknown")

		case <-success:
			driver.Status = "success"
			log.Println(driver.Username, "success")

		case pesanprocess := <-process:
			driver.Status = "process"
			driver.Pesan = pesanprocess
			log.Println(driver.Username, "process")

		case pesangagal := <-gagal:
			driver.Status = "gagal"
			driver.Pesan = pesangagal
			log.Println(driver.Username, "gagal ", pesangagal)
			time.Sleep(time.Second * 10)

		case err := <-errorcek:
			driver.Status = "error"
			driver.Pesan = err.Error()
			log.Println(driver.Username, "error", err)
		}

		return nil
	})
}

var KycStatus = map[int]string{
	-1: "gagal",
	0:  "process",
	1:  "success",
}

func CheckVerifV2(driver *report.CekVerifReport) error {

	sellerApi := api.NewTokopediaApi(driver.Session)
	b := retry.NewFibonacci(time.Second)

	return retry.Do(context.Background(), retry.WithMaxRetries(3, b), func(ctx context.Context) error {

		kyc, err := sellerApi.GetInfoKyc()
		if err != nil {
			if errors.Is(err, api.ErrKycInfoNotFound) {
				driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {
					return driver.SellerLogin(dctx)
				})
				return retry.RetryableError(err)
			}
			return err
		}

		driver.Status = KycStatus[kyc.Status]
		if driver.Status == "" {
			driver.Status = "unknown"
		}

		if len(kyc.Reason) > 0 {
			driver.Pesan = kyc.Reason[0]
		} else {
			driver.Pesan = kyc.StatusName
		}

		return nil
	})
}

func runCheckKtp(cCtx *cli.Context) error {
	// cfgname := "data/config.json"
	// pdc_common.SetConfig(cfgname, config.Version, "golang_tokopedia_check_ktp", config.Cred)
	// pdc_common.InitializeLogger()

	fname := cCtx.String("fname")
	akuns, save, _ := report.NewCekVerifReport(fname)
	defer save()
	log.Println("running checkverif...")

	for _, driver := range akuns {
		if driver.Status == "success" || driver.Status == "gagal" {
			continue
		}

		CheckVerif(driver)
	}

	return nil

}

func CreateCheckVerifCommand() *cli.Command {
	command := cli.Command{
		Name:    "cekverif",
		Aliases: []string{"cv"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "base",
				Aliases: []string{"b"},
				Value:   "./",
			},
			&cli.StringFlag{
				Name:    "fname",
				Aliases: []string{"f"},
				Value:   "akun_verification.txt",
			},
		},
		Action: runCheckKtp,
	}
	return &command
}
