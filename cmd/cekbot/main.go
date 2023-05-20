package main

import (
	"log"
	"os"
	"time"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/report"
)

var devmode = os.Getenv("DEV_MODE") != ""

func main() {

	// proxy := tokpedproxy.NewInspectProxy("localhost:8082", context.Background())
	// go proxy.RunProxy()

	// driver, _ := tokopedia_lib.NewDriverAccount("bethdunn892@outlook.com", "MZT2Zk8U", "FSR3 CTR2 5ZJX XIL5 TVK6 E72R HSRA U5GW")
	// driver.Proxy = proxy.Addr

	akuns, save, err := report.NewCekReport("cekbot.txt")
	defer save()
	if err != nil {
		pdc_common.ReportError(err)
		time.Sleep(time.Minute)
		return
	}

	for _, driver := range akuns {
		driver.DevMode = devmode

		driver.CreateApi()

	}

	log.Println("cekbot selesai..")
}
