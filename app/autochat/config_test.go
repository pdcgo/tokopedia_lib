package autochat_test

import (
	"encoding/json"
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/autochat"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {

	scen := scenario.NewScenario(t)
	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {

		t.Run("test config not exist", func(t *testing.T) {

			config, err := autochat.NewAutochatConfig(scen)
			assert.Nil(t, err)
			assert.NotEmpty(t, config)
		})

		t.Run("test config exist", func(t *testing.T) {

			fname := scen.Path(autochat.AutoConfigName)
			newconfig := autochat.AutochatConfig{
				Concurrent: 100,
			}

			b, err := json.Marshal(&newconfig)
			assert.Nil(t, err)
			remove := scen.CreateFile(b, fname)
			defer remove()

			config, err := autochat.NewAutochatConfig(scen)
			assert.Nil(t, err)
			assert.NotEmpty(t, config)
			assert.Equal(t, config.Concurrent, 100)
		})
	})
}
