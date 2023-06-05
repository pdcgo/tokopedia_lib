package main

import (
	"time"

	"github.com/pdcgo/tokopedia_lib"
)

func main() {
	driver, _ := tokopedia_lib.NewDriverAccount("pdcthoni@gmail.com", "SilentIsMyMantra", "IULIWGH6TIK3CZBKHGE27DBRLQ5LR5WQ")

	driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {
		driver.SellerLogin(dctx)

		time.Sleep(time.Hour)

		return nil
	})

}
