package report_test

import (
	"os"
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/autochat/report"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestAutoreplyReport(t *testing.T) {

	scen := scenario.NewScenario(t)
	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {

		autoreport := report.NewAutoreplyReport(scen)
		fname := scen.Path(report.AutoreplyReportName)

		t.Run("test save autorely report", func(t *testing.T) {
			autoreport.Items = append(autoreport.Items, &report.AutoreplyReportItem{
				Username: "test_save",
			})

			err := report.SaveAutoreplyReport(autoreport)
			assert.Nil(t, err)

			b, err := os.ReadFile(fname)
			assert.Nil(t, err)
			assert.Contains(t, string(b), "test_save")
		})

		t.Run("test create n update autorely report item", func(t *testing.T) {

			_, updateItem := autoreport.CreateItem("test_create", "")

			b, err := os.ReadFile(fname)
			assert.Nil(t, err)
			assert.Contains(t, string(b), "test_create")

			t.Run("test update report item", func(t *testing.T) {
				updateItem(func(item *report.AutoreplyReportItem) error {
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
