package main

import (
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/socket"
)

func main() {
	driver, _ := tokopedia_lib.NewDriverAccount("bethdunn892@outlook.com", "MZT2Zk8U", "FSR3 CTR2 5ZJX XIL5 TVK6 E72R HSRA U5GW")
	// dCtx, cancel := driver.CreateContext(false)
	// defer cancel()
	// driver.SellerLogin(dCtx)
	sClient := socket.CreateSocketClient(driver.Session)
	sClient.ConnectWebsocket()
}
