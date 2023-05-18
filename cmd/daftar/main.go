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
	"github.com/pdcgo/tokopedia_lib"
)

var baseKtp = "./ktp"

func createBaseKtp(nik string) (string, string, error) {
	nikfoto := filepath.Join(baseKtp, nik+".jpg")
	nikfoto, _ = filepath.Abs(nikfoto)
	selfiefoto := filepath.Join(baseKtp, nik+"_2.jpg")
	selfiefoto, _ = filepath.Abs(selfiefoto)

	if _, err := os.Stat(nikfoto); errors.Is(err, os.ErrNotExist) {
		return "", "", errors.New(nikfoto + "tidak ada")
	}
	if _, err := os.Stat(selfiefoto); errors.Is(err, os.ErrNotExist) {
		return "", "", errors.New(nikfoto + "tidak ada")
	}

	return nikfoto, selfiefoto, nil
}

func runVerification(akun *tokopedia_lib.Account) error {
	driver, _ := tokopedia_lib.NewDriverAccount(akun.Username, akun.Pass, akun.Secret, "", "")

	nikfoto, selfiefoto, err := createBaseKtp(akun.Ktp)
	if err != nil {
		akun.Status = err.Error()
		return nil
	}

	done := make(chan int, 1)
	errChan := make(chan error, 1)

	driver.Run(false, func(dctx *tokopedia_lib.DriverContext) {
		driver.ExecLogin(dctx)

		submitCtx, cancel := context.WithCancel(dctx.Ctx)

		button1 := `//*/label[@for="pic1"]`
		finish := `//*/button/div/span[contains(text(), "Upload")]`
		finishTitle := `//*/div[contains(text(), "Fotomu sedang kami proses")]`
		toast := `//*/div[@data-unify="Toaster"]`
		chromedp.Run(dctx.Ctx,
			chromedp.Navigate("https://mitra.tokopedia.com/kyc"),
			chromedp.WaitReady(button1, chromedp.BySearch),
		)
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
						pesan = strings.ReplaceAll(pesan, "\n", "")
						pesan = strings.ReplaceAll(pesan, "OK", "")
						pesan = strings.ReplaceAll(pesan, "\t", "")
						pesan = strings.ReplaceAll(pesan, "'", "")
						cancel()
						errChan <- errors.New(pesan)
						time.Sleep(time.Second * 3)
					}

				}
			}
		}()

		go chromedp.Run(dctx.Ctx,
			chromedp.WaitReady(finish, chromedp.BySearch),
			chromedp.Click(finish, chromedp.BySearch),
		)
		go chromedp.Run(dctx.Ctx,
			chromedp.WaitReady(finishTitle, chromedp.BySearch),
			chromedp.ActionFunc(func(ctx context.Context) error {
				done <- 1

				return nil
			}),
		)

		title := `//*/div[contains(text(), "Foto Bagian Depan KTP")]`
		go func() {
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

		}()

	})

	select {
	case <-done:
		log.Println(driver.Username, "Submit Done")
		akun.Status = "done"
	case err := <-errChan:
		log.Println(driver.Username, "error", err.Error())
		akun.Status = err.Error()
	}

	return nil
}

func main() {
	akuns, save, err := tokopedia_lib.GetVerificationAkun("akun_verification.txt")
	if err != nil {
		panic(err)
	}

	for _, akun := range akuns {
		runVerification(akun)
	}

	save()
}
