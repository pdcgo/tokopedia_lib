package autosubmit

import (
	"context"
	"encoding/json"
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
	"github.com/pdcgo/tokopedia_lib/app/config"
	"github.com/urfave/cli/v2"
)

func CreateSubmitCommand(cred []byte) *cli.Command {
	command := cli.Command{
		Name:    "submit_ktp",
		Aliases: []string{"submit"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "base",
				Aliases: []string{"b"},
				Value:   "./",
			},
		},
		Action: func(ctx *cli.Context) error {
			cfgname := "data/config.json"
			pdc_common.SetConfig(cfgname, config.Version, "golang_tokopedia_submit", cred)
			pdc_common.InitializeLogger()

			base := ctx.String("base")

			submit, err := LoadTokopediaAutoSubmit(base)
			submit.RunAutoSubmit()

			return err
		},
	}
	return &command
}

type AutoSubmit struct {
	BaseKtp  string `json:"base_ktp"` // "akun_verification.txt"
	Filename string `json:"filename"` // "./ktp",
}

var nameConfig = "data/tokopedia_submit_config"

func SaveTokopediaAutoSubmit(base string, payload *AutoSubmit) error {
	fname := filepath.Join(base, nameConfig)

	f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(payload)

}

func LoadTokopediaAutoSubmit(base string) (*AutoSubmit, error) {
	fname := filepath.Join(base, nameConfig)
	submit := AutoSubmit{
		BaseKtp:  "akun_verification.txt",
		Filename: "./ktp",
	}

	if _, err := os.Stat(fname); errors.Is(err, os.ErrNotExist) {
		return &submit, nil
	}

	data, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &submit)
	if err != nil {
		return nil, err
	}

	return &submit, nil

}

func (auto *AutoSubmit) createBaseKtp(nik string) (string, string, error) {
	baseKtp := auto.BaseKtp

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

func (auto *AutoSubmit) RunAutoSubmit() {
	akuns, save, err := tokopedia_lib.GetVerificationAkun(auto.Filename)
	if err != nil {
		pdc_common.ReportError(err)
		time.Sleep(time.Minute)
		return
	}

	for _, akun := range akuns {
		auto.RunVerification(akun)
		save()
	}

	save()
}

func (auto *AutoSubmit) RunVerification(akun *tokopedia_lib.Account) error {
	driver, _ := tokopedia_lib.NewDriverAccount(akun.Username, akun.Pass, akun.Secret)

	nikfoto, selfiefoto, err := auto.createBaseKtp(akun.Ktp)
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
