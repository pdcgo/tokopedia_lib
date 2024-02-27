package report_test

import (
	"os"
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/autochat/report"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestAutosendReport(t *testing.T) {

	scen := scenario.NewScenario(t)
	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {

		autoreport := report.NewAutosendReport(scen)
		fname := scen.Path(report.AutosendReportName)

		t.Run("test save autosend report", func(t *testing.T) {
			autoreport.Items = append(autoreport.Items, &report.AutosendReportItem{
				Username:        "test_save",
				SellerChatCount: 10,
			})

			err := report.SaveAutosendReport(autoreport)
			assert.Nil(t, err)

			b, err := os.ReadFile(fname)
			assert.Nil(t, err)
			assert.Contains(t, string(b), "test_save")
		})

		t.Run("test create n update autosend report item", func(t *testing.T) {

			_, updateItem := autoreport.CreateItem("test_create", 10)

			b, err := os.ReadFile(fname)
			assert.Nil(t, err)
			assert.Contains(t, string(b), "test_create")

			t.Run("test update report item", func(t *testing.T) {
				updateItem(func(item *report.AutosendReportItem) error {
					item.Username = "test_update"
					return nil
				})

				b, err := os.ReadFile(fname)
				assert.Nil(t, err)
				assert.Contains(t, string(b), "test_update")
			})
		})
	})
}
