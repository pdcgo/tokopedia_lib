package main

import (
	"log"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/app/withdraw"
	"github.com/pdcgo/tokopedia_lib/lib/helper"
)

var akunfilename = "akun.txt"
var wdReport = "wd_report.csv"

func SaveReport(reports []*withdraw.WithdrawReport) error {
	file, err := os.OpenFile(wdReport, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	err = gocsv.MarshalFile(reports, file)

	return err

}

func main() {
	log.Println("[ WD BOT ]: Mulai withdrawl")

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

	results, err := withdraw.RunWithdraw(drivers)
	if err != nil {
		pdc_common.ReportError(err)
		return
	}

	reports := []*withdraw.WithdrawReport{}
	for result := range results {
		reports = append(reports, result)
	}

	SaveReport(reports)
}
