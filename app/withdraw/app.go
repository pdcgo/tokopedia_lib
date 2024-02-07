package withdraw

import (
	"context"
	"errors"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/api"
)

type WDStatus string

const (
	SUCCESS WDStatus = "SUKSES"
	FAILDED WDStatus = "GAGAL"
)

type WithdrawReport struct {
	Email      string   `csv:"email"`
	ShopName   string   `csv:"shop_name"`
	Transaksi  string   `csv:"transaksi"`
	Invoice    string   `csv:"invoice"`
	Jumlah     string   `csv:"jumlah"`
	SisaSaldo  string   `csv:"sisa_saldo"`
	Status     WDStatus `csv:"status"`
	Keterangan string   `csv:"keterangan"`
}

func (wd *WithdrawReport) Headers() []string {
	return []string{"email", "shop_name", "transaksi", "invoice", "jumlah", "sisa_saldo", "status", "keterangan"}
}

func (wd *WithdrawReport) Values() []string {
	return []string{wd.Email, wd.ShopName, wd.Transaksi, wd.Invoice, wd.Jumlah, wd.SisaSaldo, string(wd.Status), wd.Keterangan}
}

func RunWithdraw(payload []*tokopedia_lib.DriverAccount) error {
	for _, driver := range payload {
		tApi, _, err := driver.CreateApi()
		if err != nil {
			return err
		}
		defer func() {
			driver.Session.SaveSession()
		}()

		err = run(driver, tApi)
		if err != nil {
			return err
		}
	}

	return nil
}

func run(driver *tokopedia_lib.DriverAccount, api *api.TokopediaApi) error {
	withdraw, err := NewWithdraw(driver, api)
	if err != nil {
		return err
	}

	return withdraw.Run()
}

func WithdrawWithReport(driver *tokopedia_lib.DriverAccount, api *api.TokopediaApi) (*WithdrawReport, error) {
	result := &WithdrawReport{
		Email:     driver.Username,
		Transaksi: "Penarikan Saldo Penghasilan",
		Jumlah:    "Rp0",
		SisaSaldo: "Rp0",
		Status:    SUCCESS,
	}

	wd, err := NewWithdraw(driver, api)
	if err != nil {
		if errors.Is(err, ErrSaldoKosong) {
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

type HandlerWithdrawReport func(reports []*WithdrawReport) error

func RunWithdrawWithReport(drivers []*tokopedia_lib.DriverAccount, handler HandlerWithdrawReport) error {
	for _, driver := range drivers {
		tApi, _, err := driver.CreateApi()
		if err != nil {
			return err
		}
		defer func() {
			driver.Session.SaveSession()
		}()

		reports, err := GetUnwithdrawTransaction(driver, tApi)
		if err != nil {
			return err
		}

		result, err := WithdrawWithReport(driver, tApi)
		if err != nil {
			return err
		}
		result.ShopName = tApi.AuthenticatedData.UserShopInfo.Info.ShopName

		if result.Status == FAILDED {
			if result.Keterangan == ErrSaldoKosong.Error() {
				reports := []*WithdrawReport{result}
				err := handler(reports)
				if err != nil {
					return err
				}
				continue
			}
		}

		for _, report := range reports {
			report.Status = result.Status
		}

		reports = append(reports, result)
		err = handler(reports)
		if err != nil {
			return err
		}
	}

	return nil
}
