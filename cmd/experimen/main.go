package main

import (
	"time"

	"github.com/pdcgo/tokopedia_lib"
)

func main() {
	driver, _ := tokopedia_lib.NewDriverAccount("mrwicaksono690@gmail.com", "Semogaberkah", "KQ2YZZJDS2DAC7Y6K4HCTTXSS5B7N4IN")

	driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {
		driver.SellerLogin(dctx)

		time.Sleep(time.Hour)

		return nil
	})

}
