package main

import (
	"context"
	"encoding/csv"
	"errors"
	"log"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/app/withdraw"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/pdcgo/tokopedia_lib/lib/helper"
)

type WDStatus string

const (
	SUCCESS WDStatus = "SUKSES"
	FAILDED WDStatus = "GAGAL"
)

type WithdrawReport struct {
	Email      string   `csv:"email"`
	Transaksi  string   `csv:"transaksi"`
	Invoice    string   `csv:"invoice"`
	Jumlah     string   `csv:"jumlah"`
	SisaSaldo  string   `csv:"sisa_saldo"`
	Status     WDStatus `csv:"status"`
	Keterangan string   `csv:"keterangan"`
}

func (wd *WithdrawReport) Headers() []string {
	return []string{"email", "transaksi", "invoice", "jumlah", "sisa_saldo", "status", "keterangan"}
}

func (wd *WithdrawReport) Values() []string {
	return []string{wd.Email, wd.Transaksi, wd.Invoice, wd.Jumlah, wd.SisaSaldo, string(wd.Status), wd.Keterangan}
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
			wd := &WithdrawReport{}

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

	results := make(chan []*WithdrawReport, 3)

	go func() {
		defer close(results)

		for _, akun := range akuns {
			data := strings.Split(akun, "|")

			log.Printf("[ WD BOT ]: %s  mulai withdraw", data[0])

			secret := strings.ReplaceAll(data[2], " ", "")
			driver, err := tokopedia_lib.NewDriverAccount(data[0], data[1], secret)
			if err != nil {
				pdc_common.ReportError(err)
				return
			}
			driver.SetPIN(data[3])

			tApi, _, err := driver.CreateApi()
			if err != nil {
				pdc_common.ReportError(err)
				return
			}
			defer func() {
				driver.Session.SaveSession()
			}()

			reports, err := GetTransaction(driver, tApi)
			if err != nil {
				pdc_common.ReportError(err)
				return
			}

			result, err := RunWithdraw(driver)
			if err != nil {
				pdc_common.ReportError(err)
				return
			}

			log.Printf("[ STATUS ] : %s %s", result.Email, result.Status)

			if result.Status == FAILDED {
				if result.Keterangan == withdraw.ErrSaldoKosong.Error() {
					reports := []*WithdrawReport{result}
					results <- reports
					continue
				}
			}

			for _, report := range reports {
				report.Status = result.Status
			}

			reports = append(reports, result)
			results <- reports
		}
	}()

	for result := range results {
		for _, report := range result {
			SaveCsvData(wdReport, report.Values())
		}

	}
}

func GetTransaction(driver *tokopedia_lib.DriverAccount, tApi *api.TokopediaApi) ([]*WithdrawReport, error) {
	reports := []*WithdrawReport{}

	now := time.Now().UTC()
	lastMonth := now.AddDate(0, 0, -31)

	haveNextPage := true
Parent:
	for haveNextPage {
		variable := api.NewDepositHistoryVariable()
		variable.MaxRows = 100
		variable.DateFrom = lastMonth.Format(api.YYYYMMDD)
		depositHistories, err := tApi.MidasGetDepositHistory(variable)
		if err != nil {
			return reports, err
		}

		depositContent := depositHistories.GetContent()
		var unWithdrawTransaction []*api.DepositContent
		for ind, history := range depositContent {
			if history.Type == 7001 || history.Class == "Withdrawal" {
				if ind == 0 {
					break Parent
				}

				unWithdrawTransaction = append(unWithdrawTransaction, depositContent[:ind]...)
				break
			}
		}

		if unWithdrawTransaction == nil {
			unWithdrawTransaction = depositContent
		}

		for _, transaction := range unWithdrawTransaction {
			inv := strings.Split(transaction.Note, "-")

			report := &WithdrawReport{
				Email:     driver.Username,
				Transaksi: transaction.TypeDescription,
				Invoice:   inv[len(inv)-1],
				Jumlah:    transaction.AmountFmt,
				SisaSaldo: transaction.SaldoFmt,
			}

			reports = append(reports, report)
		}

		haveNextPage = depositHistories.Data.MidasGetDepositHistory.HaveNextPage
	}

	slices.Reverse(reports)
	return reports, nil
}

func RunWithdraw(driver *tokopedia_lib.DriverAccount) (*WithdrawReport, error) {
	result := &WithdrawReport{
		Email:     driver.Username,
		Transaksi: "Penarikan Saldo Penghasilan",
		Jumlah:    "Rp0",
		SisaSaldo: "Rp0",
		Status:    SUCCESS,
	}

	wd, err := withdraw.NewWithdraw(driver)
	if err != nil {
		if errors.Is(err, withdraw.ErrSaldoKosong) {
			result.Status = FAILDED
			result.Keterangan = err.Error()
			return result, nil
		}
		return nil, err
	}
	result.Jumlah = wd.Balance.Data.MidasGetAllDepositAmount.SellerAllFmt

	err = wd.Run()
	if err != nil {
		result.Status = FAILDED
		result.Keterangan = err.Error()
		if errors.Is(err, context.Canceled) {
			result.Keterangan = "coba manual"
		}
	}

	return result, nil
}
