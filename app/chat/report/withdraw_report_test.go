package report_test

import (
	"os"
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/chat/report"
	"github.com/pdcgo/tokopedia_lib/app/withdraw"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestWithdrawReport(t *testing.T) {

	scen := scenario.NewScenario(t)
	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {

		fname := scen.Path("report.csv")
		report := report.NewWitdrawReport(fname)

		report.Add(&withdraw.WithdrawReport{
			Email:    "test@mail.com",
			ShopName: "testshop",
		})
		report.Save()

		b, err := os.ReadFile(fname)
		assert.Nil(t, err)
		assert.Contains(t, string(b), "test@mail.com")
		assert.Contains(t, string(b), "testshop")
	})
}
