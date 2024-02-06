package main

import (
	"context"
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strings"
	"sync/atomic"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/app/app_config"
	"github.com/pdcgo/tokopedia_lib/app/withdraw"
	"github.com/pdcgo/tokopedia_lib/lib/helper"
)

type WD_Status string

const (
	SUCCESS WD_Status = "SUKSES"
	FAILDED WD_Status = "GAGAL"
)

type WithdrawResult struct {
	Email      string    `csv:"email"`
	Status     WD_Status `csv:"status"`
	Keterangan string    `csv:"keterangan"`
}

func (wd *WithdrawResult) Headers() []string {
	return []string{"email", "status", "keterangan"}
}

func (wd *WithdrawResult) Values() []string {
	return []string{wd.Email, string(wd.Status), wd.Keterangan}
}

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
			wd := &WithdrawResult{}

			os.Create(wdReport)
			SaveCsvData(wdReport, wd.Headers())
		}
	}
}

func main() {
	prox, cancel := app_config.GetProxy(app_config.WithdrawProxyKey)
	defer cancel()

	akuns, err := helper.FileLoadLineString(akunfilename)
	if err != nil {
		pdc_common.ReportError(err)
		return
	}

	results := make(chan *WithdrawResult, 3)

	go func() {
		defer close(results)

		for _, akun := range akuns {
			data := strings.Split(akun, "|")
			secret := strings.ReplaceAll(data[2], " ", "")

			driver, err := tokopedia_lib.NewDriverAccount(data[0], data[1], secret)
			if err != nil {
				pdc_common.ReportError(err)
				return
			}
			driver.SetPIN(data[3])
			driver.Proxy = prox.Addr

			result := &WithdrawResult{
				Email:  driver.Username,
				Status: SUCCESS,
			}

			wd, err := withdraw.NewWithdraw(driver)
			if err != nil {
				if errors.Is(err, withdraw.ErrSaldoKosong) {
					result.Status = FAILDED
					result.Keterangan = err.Error()
					results <- result
					continue
				}
				pdc_common.ReportError(err)
				return
			}
			err = wd.Run()
			if err != nil {
				result.Status = FAILDED
				result.Keterangan = err.Error()
				if errors.Is(err, context.Canceled) {
					result.Keterangan = "coba manual"
				}
			}

			results <- result
		}
	}()

	counter := int32(0)
	for result := range results {
		count := atomic.AddInt32(&counter, 1)
		if result.Status == SUCCESS {
			log.Printf("[ STATUS ] : %d Akun %s %s", count, result.Email, result.Status)
		} else {
			log.Printf("[ STATUS ] : %d Akun %s %s %s", count, result.Email, result.Status, result.Keterangan)
		}

		SaveCsvData(wdReport, result.Values())
	}
}
