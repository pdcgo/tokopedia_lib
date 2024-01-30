package main

import (
	"strings"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/app/withdraw"
	"github.com/pdcgo/tokopedia_lib/lib/helper"
)

var akunfilename = "akun.txt"

func main() {
	akuns, err := helper.FileLoadLineString(akunfilename)
	if err != nil {
		pdc_common.ReportError(err)
		return
	}

	drivers := []*tokopedia_lib.DriverAccount{}
	for _, akun := range akuns {
		data := strings.Split(akun, "|")
		secret := strings.ReplaceAll(data[2], " ", "")

		driver, err := tokopedia_lib.NewDriverAccount(data[0], data[1], secret)
		if err != nil {
			pdc_common.ReportError(err)
			return
		}
		driver.SetPIN(data[3])

		drivers = append(drivers, driver)
	}

	err = withdraw.RunWithdraw(drivers)
	if err != nil {
		pdc_common.ReportError(err)
		return
	}
}
