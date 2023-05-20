package main

import (
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/pdcgo/tokopedia_lib"
)

// func reportFile() (func(data string), func()) {

// }

func main() {
	akuns, save, _ := getakunFromFile("akun.txt")

	for _, driver := range akuns {
		if driver.Status == "success" || driver.Status == "gagal" {
			continue
		}
		driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {
			driver.ExecLogin(dctx)

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

}
