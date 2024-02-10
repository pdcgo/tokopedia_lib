package withdraw

import (
	"errors"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib"
)

type WDStatus string

const (
	SUCCESS WDStatus = "SUKSES"
	FAILDED WDStatus = "GAGAL"
)

type WithdrawReport struct {
	Email        string   `csv:"email"`
	ShopName     string   `csv:"shop_name"`
	Type         int      `csv:"type"`
	Deskripsi    string   `csv:"deskripsi"`
	Invoice      string   `csv:"invoice"`
	Jumlah       string   `csv:"jumlah"`
	SisaSaldo    string   `csv:"sisa_saldo"`
	Status       WDStatus `csv:"status"`
	Keterangan   string   `csv:"keterangan"`
	ErrorMessage string   `csv:"error_message"`
}

func RunWithdraw(payload []*tokopedia_lib.DriverAccount) (chan []*WithdrawReport, error) {
	reports := make(chan []*WithdrawReport)

	go func() {
		defer close(reports)

		for _, driver := range payload {
			tApi, _, err := driver.CreateApi()
			if err != nil {
				pdc_common.ReportError(err)
				continue
			}
			defer func() {
				driver.Session.SaveSession()
			}()

			items, err := GetUnwithdrawTransaction(tApi)
			if err != nil {
				pdc_common.ReportError(err)
				continue
			}

			item := &WithdrawReport{
				Email:      driver.Username,
				ShopName:   tApi.AuthenticatedData.UserShopInfo.Info.ShopName,
				Type:       7001,
				Jumlah:     "Rp0",
				SisaSaldo:  "Rp0",
				Status:     SUCCESS,
				Keterangan: "Withdrawal",
			}

			withdraw := NewWithdraw(tApi)

			err = driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {
				err := withdraw.Run(dctx, driver.PIN, item)
				if err != nil {
					item.Status = FAILDED
					item.Keterangan = err.Error()
					if errors.Is(err, ErrSaldoKosong) {
						return nil
					}
				}
				return err
			})
			if err != nil {
				pdc_common.ReportError(err)
			}

			items = append(items, item)
			reports <- items
		}
	}()

	return reports, nil
}
