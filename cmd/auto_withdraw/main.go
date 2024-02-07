package main

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/app/withdraw"
	"github.com/pdcgo/tokopedia_lib/lib/helper"
)

var akunfilename = "akun.txt"
var wdReport = "wd_report.csv"

func SaveCsvData(path string, data []string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(data); err != nil {
		return err
	}
	return nil
}

func init() {
	_, err := os.Stat(wdReport)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			wd := &withdraw.WithdrawReport{}

			os.Create(wdReport)
			SaveCsvData(wdReport, wd.Headers())
		}
	}
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

	err = withdraw.RunWithdrawWithReport(drivers, func(reports []*withdraw.WithdrawReport) error {
		log.Printf("[ STATUS ]: %s %s", reports[0].Email, reports[0].Status)
		for _, report := range reports {
			err := SaveCsvData(wdReport, report.Values())
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		pdc_common.ReportError(err)
		return
	}

}
