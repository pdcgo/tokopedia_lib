package config_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/chat/config"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestAppConfig(t *testing.T) {

	scen := scenario.NewScenario(t)
	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {

		appconfig := config.NewAppConfig(dirbase)
		assert.Equal(t, appconfig.Base, dirbase)
		assert.Equal(t, appconfig.Host, "localhost")
		assert.Equal(t, appconfig.Port, "5003")
		assert.True(t, appconfig.OpenBrowser)
		assert.False(t, appconfig.DebugMode)
		assert.Equal(t, appconfig.SyncCommandInterval, [2]float32{5, 30})
	})
}
