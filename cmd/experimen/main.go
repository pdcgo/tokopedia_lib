package main

import (
	"time"

	"github.com/pdcgo/tokopedia_lib"
)

func main() {
	driver, _ := tokopedia_lib.NewDriverAccount("delaccozz12@yahoo.com", "M@ret2022", "QFY76RQYBXJ6NPBWVF24HF7QUT7IAG74")

	driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {
		driver.SellerLogin(dctx)

		time.Sleep(time.Hour)

		return nil
	})

}
