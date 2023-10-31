package main

import (
	"context"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/go-vgo/robotgo"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib"
)

var baseKtp = "./ktp"

func createBaseKtp(nik string) (string, string, error) {
	nikfoto := filepath.Join(baseKtp, nik+".jpg")
	nikfoto, _ = filepath.Abs(nikfoto)
	selfiefoto := filepath.Join(baseKtp, nik+"_2.jpg")
	selfiefoto, _ = filepath.Abs(selfiefoto)

	if _, err := os.Stat(nikfoto); errors.Is(err, os.ErrNotExist) {
		return "", "", errors.New(nikfoto + " tidak ada")
	}
	if _, err := os.Stat(selfiefoto); errors.Is(err, os.ErrNotExist) {
		return "", "", errors.New(selfiefoto + " selfie tidak ada")
	}

	return nikfoto, selfiefoto, nil
}

func runVerification(akun *tokopedia_lib.Account) error {
	driver, _ := tokopedia_lib.NewDriverAccount(akun.Username, akun.Pass, akun.Secret)

	nikfoto, selfiefoto, err := createBaseKtp(akun.Ktp)
	if err != nil {
		akun.Status = err.Error()
		return nil
	}

	done := make(chan int, 1)
	errChan := make(chan error)
	defer close(errChan)

	return driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {
		driver.SellerLogin(dctx)

		submitCtx, cancel := context.WithCancel(dctx.Ctx)
		defer cancel()

		button1 := `//*/label[@for="pic1"]`
		finish := `//*/button/div/span[contains(text(), "Upload")]`
		finishTitle := `//*/div[contains(text(), "Fotomu sedang kami proses")]`
		toast := `//*/div[@data-unify="Toaster"]`

		go func() {
		Parent:
			for {
				select {
				case <-dctx.Ctx.Done():
					break Parent
				default:
					var pesan string
					chromedp.Run(dctx.Ctx,
						chromedp.WaitReady(toast, chromedp.BySearch),
						chromedp.Text(toast, &pesan, chromedp.BySearch),
					)
					if pesan != "" {
						log.Println("error", pesan)
						pesan = strings.ReplaceAll(pesan, "\n", "")
						pesan = strings.ReplaceAll(pesan, "OK", "")
						pesan = strings.ReplaceAll(pesan, "\t", "")
						pesan = strings.ReplaceAll(pesan, "'", "")
						// cancel()
						errChan <- errors.New(pesan)
						time.Sleep(time.Second * 3)
					}

				}
			}
		}()

		go func() {
			for {
				select {
				case <-submitCtx.Done():
					return
				default:
					chromedp.Run(dctx.Ctx,
						chromedp.WaitReady(finish, chromedp.BySearch),
						chromedp.Click(finish, chromedp.BySearch),
					)
				}
			}
		}()

		go chromedp.Run(dctx.Ctx,
			chromedp.WaitReady(finishTitle, chromedp.BySearch),
			chromedp.ActionFunc(func(ctx context.Context) error {
				done <- 1

				return nil
			}),
		)

		isiktp := func() {
			chromedp.Run(dctx.Ctx,
				chromedp.Navigate("https://mitra.tokopedia.com/kyc"),
				chromedp.WaitReady(button1, chromedp.BySearch),
			)

			title := `//*/div[contains(text(), "Foto Bagian Depan KTP")]`
			chromedp.Run(submitCtx,
				chromedp.WaitReady(button1, chromedp.BySearch),
				chromedp.Sleep(time.Second*4),
				chromedp.Click(button1, chromedp.BySearch),
				chromedp.Sleep(time.Second*2),
				chromedp.ActionFunc(func(ctx context.Context) error {
					robotgo.TypeStr(nikfoto)
					robotgo.KeyTap("enter")
					return nil
				}),
				chromedp.WaitReady(title, chromedp.BySearch),
				chromedp.WaitReady(button1, chromedp.BySearch),
				chromedp.Sleep(time.Second*4),
				chromedp.Click(button1, chromedp.BySearch),
				chromedp.Sleep(time.Second*2),
				chromedp.ActionFunc(func(ctx context.Context) error {
					robotgo.TypeStr(selfiefoto)
					robotgo.KeyTap("enter")
					return nil
				}),
			)

		}

		go isiktp()

		// jika ktp gagal
		go func() {
			tick := time.NewTicker(time.Second * 2)
			defer tick.Stop()

			titlerr := `//*/h5[@data-unify="Typography"]`
			cek := "Foto KTP gagal diproses"

			retry := 5

			for {
				select {
				case <-submitCtx.Done():
					return
				case <-tick.C:
					title := ""
					chromedp.Run(submitCtx,
						chromedp.WaitVisible(titlerr, chromedp.BySearch),
						chromedp.Text(titlerr, &title, chromedp.BySearch),
					)
					log.Println(title)
					if strings.Contains(title, cek) {
						go isiktp()
						retry -= 1
					}

					if retry <= 0 {
						return
					}
				}
			}

		}()

		timeout := time.After(time.Minute)

		select {
		case <-done:
			log.Println(driver.Username, "Submit Done")
			akun.Status = "done"
		case err := <-errChan:
			log.Println(driver.Username, "error", err.Error())
			pdc_common.ReportError(err)
			akun.Status = err.Error()
			time.Sleep(time.Minute)

		case <-timeout:
			akun.Status = "timeout 60 second"
		}

		return nil
	})

}

func main() {
	akuns, save, err := tokopedia_lib.GetVerificationAkun("akun_verification.txt")
	if err != nil {
		pdc_common.ReportError(err)
		time.Sleep(time.Minute)
		return
	}

	for _, akun := range akuns {
		runVerification(akun)
		save()
	}

	save()
}
