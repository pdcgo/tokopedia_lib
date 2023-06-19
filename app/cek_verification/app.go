package cek_verification

import (
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/urfave/cli/v2"
)

// func reportFile() (func(data string), func()) {

// }

func runCheckKtp(cCtx *cli.Context) error {
	fname := cCtx.String("fname")
	akuns, save, _ := getakunFromFile(fname)

	for _, driver := range akuns {
		if driver.Status == "success" || driver.Status == "gagal" {
			continue
		}
		driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {
			driver.MitraLogin(dctx.Ctx)

			timeout := time.After(time.Second * 30)
			success := make(chan int, 1)
			gagal := make(chan string, 1)

			go func() {
				pathverif := `//*/div[contains(text(), "Terverifikasi")]`
				chromedp.Run(dctx.Ctx,
					chromedp.Navigate("https://mitra.tokopedia.com/user/akun-saya"),
					chromedp.WaitVisible(pathverif, chromedp.BySearch),
				)

				success <- 1
			}()

			go func() {
				var hasil string
				title := `//*/h5[@data-unify="Typography"]`
				chromedp.Run(dctx.Ctx,
					chromedp.WaitVisible(title, chromedp.BySearch),
					chromedp.Text(`//*/div[@aria-label="unf-ticker-item"]/div/span/p[@data-unify="Typography"]`, &hasil, chromedp.NodeVisible),
				)

				gagal <- hasil
			}()

			select {
			case <-timeout:
				driver.Status = "unknown"
				log.Println(driver.Username, " unknown")
			case <-success:
				driver.Status = "success"
				log.Println(driver.Username, " success")
			case pesangagal := <-gagal:
				driver.Status = "gagal"
				driver.Pesan = pesangagal
				log.Println(driver.Username, " gagal ", pesangagal)
			}

			save()

			return nil
		})

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
