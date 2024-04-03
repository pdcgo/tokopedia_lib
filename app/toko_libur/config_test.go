package toko_libur_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/toko_libur"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {

	scen := scenario.NewScenario(t)
	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {

		t.Run("test init config", func(t *testing.T) {
			config := toko_libur.NewApplicationConfig(scen)
			assert.False(t, config.Libur)
			assert.Equal(t, config.Report, "tokopedia_toko_libur.csv")
			assert.Equal(t, config.Limit, 3)
		})

		t.Run("test save config", func(t *testing.T) {
			config := toko_libur.NewApplicationConfig(scen)
			config.Libur = true
			config.Report = "report.csv"
			config.CloseStart = 1712111901
			config.CloseEnd = 1712716701

			err := toko_libur.SaveApplicationConfig(scen, config)
			assert.Nil(t, err)

			t.Run("test config changed", func(t *testing.T) {
				cconfig := toko_libur.NewApplicationConfig(scen)
				assert.True(t, cconfig.Libur)
				assert.Equal(t, cconfig.Report, "report.csv")
				assert.Equal(t, cconfig.Limit, 3)
				assert.Equal(t, cconfig.CloseStart, 1712111901)
				assert.Equal(t, cconfig.CloseEnd, 1712716701)
			})
		})
	})
}
