package config_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/chat/config"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {

	scen := scenario.NewScenario(t)
	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		appconfig := config.NewAppConfig(dirbase)
		initconfig := config.NewInitConfig(appconfig)

		t.Run("test save", func(t *testing.T) {

			initconfig.ActiveGroup = "test save"
			err := initconfig.Save()
			assert.Nil(t, err)

			newinitconfig := config.NewInitConfig(appconfig)
			assert.Equal(t, newinitconfig.ActiveGroup, "test save")
		})

		t.Run("test set group", func(t *testing.T) {

			err := initconfig.SetGroup("test set group")
			assert.Nil(t, err)

			newinitconfig := config.NewInitConfig(appconfig)
			assert.Equal(t, newinitconfig.ActiveGroup, "test set group")
		})
	})
}
