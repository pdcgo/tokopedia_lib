package toko_libur_test

import (
	"os"
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/toko_libur"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestReport(t *testing.T) {

	scen := scenario.NewScenario(t)
	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {

		config := toko_libur.TokopediaTokoLiburConfig{
			BaseConfig: scen,
			Report:     "toko_libur.csv",
		}
		report, err := toko_libur.NewReport(&config)
		assert.Nil(t, err)

		t.Run("test save report", func(t *testing.T) {

			report.Data = append(report.Data, &toko_libur.ReportItem{
				AkunTokoLibur: &toko_libur.AkunTokoLibur{
					Username: "test save",
				},
			})
			err := report.Save()
			assert.Nil(t, err)

			b, err := os.ReadFile(config.GetReportPath())
			assert.Nil(t, err)
			assert.Contains(t, string(b), "test save")
		})

		t.Run("test iterate n update item", func(t *testing.T) {

			err := report.Iterate(func(item *toko_libur.ReportItem, updateitem func(handler func(item *toko_libur.ReportItem) error) error) error {
				return updateitem(func(item *toko_libur.ReportItem) error {
					item.Username = "test update"
					return nil
				})
			})
			assert.Nil(t, err)

			b, err := os.ReadFile(config.GetReportPath())
			assert.Nil(t, err)
			assert.Contains(t, string(b), "test update")
		})
	})
}
