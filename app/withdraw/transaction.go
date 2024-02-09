package withdraw

import (
	"strings"
	"time"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/api"
)

func GetUnwithdrawTransaction(driver *tokopedia_lib.DriverAccount, tApi *api.TokopediaApi) ([]*WithdrawReport, error) {
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
				ShopName:  tApi.AuthenticatedData.UserShopInfo.Info.ShopName,
				Transaksi: transaction.TypeDescription,
				Invoice:   inv[len(inv)-1],
				Jumlah:    transaction.AmountFmt,
				SisaSaldo: transaction.SaldoFmt,
			}

			reports = append(reports, report)
		}

		haveNextPage = depositHistories.Data.MidasGetDepositHistory.HaveNextPage
	}

	return reports, nil
}
