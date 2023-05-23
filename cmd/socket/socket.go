package main

import (
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/socket"
)

func main() {
	driver, _ := tokopedia_lib.NewDriverAccount("082329471203", "220599", "IULI WGH6 TIK3 CZBK HGE2 7DBR LQ5L R5WQ")
	// dCtx, cancel := driver.CreateContext(false)
	// defer cancel()
	// driver.SellerLogin(dCtx)
	sClient := socket.CreateSocketClient(driver.Session)
	sClient.ConnectWebsocket()
}
