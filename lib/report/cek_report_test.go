package report_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/report"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestCekReport(t *testing.T) {

	scen := scenario.NewScenario(t)
	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {

		fname := scen.Path("test_report.csv")
		driver, err := tokopedia_lib.NewDriverAccount("test", "pass", "secret")
		assert.Nil(t, err)

		t.Run("test buka report ketika file tidak ada", func(t *testing.T) {
			items, _, err := report.NewCekReport(fname)
			assert.Nil(t, err)
			assert.Empty(t, items)
		})

		t.Run("test save report", func(t *testing.T) {
			saveItem := report.CekReport{
				DriverAccount: driver,
			}
			err = report.SaveCekReport(fname, []*report.CekReport{&saveItem})
			assert.Nil(t, err)

			t.Run("test report saved", func(t *testing.T) {
				saveditems, _, err := report.NewCekReport(fname)
				assert.Nil(t, err)
				assert.NotEmpty(t, saveditems)

				for _, item := range saveditems {
					assert.Equal(t, item.Username, driver.Username)
					assert.Equal(t, item.Password, driver.Password)
					assert.Equal(t, item.Secret, driver.Secret)
				}
			})
		})

		t.Run("test buka report ketika file ada", func(t *testing.T) {
			items, save, err := report.NewCekReport(fname)
			assert.Nil(t, err)
			assert.NotEmpty(t, items)

			t.Run("test save report", func(t *testing.T) {
				for _, item := range items {
					item.Password = "repass"
				}
				err = save()
				assert.Nil(t, err)

				t.Run("test report saved", func(t *testing.T) {
					saveditems, _, err := report.NewCekReport(fname)
					assert.Nil(t, err)
					assert.NotEmpty(t, saveditems)

					for _, item := range saveditems {
						assert.Equal(t, item.Username, driver.Username)
						assert.Equal(t, item.Password, "repass")
						assert.Equal(t, item.Secret, driver.Secret)
					}
				})
			})
		})
	})
}
