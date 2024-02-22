package autochat_test

import (
	"strings"
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/autochat"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestAkun(t *testing.T) {

	scen := scenario.NewScenario(t)
	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {

		message, err := autochat.NewAutochatMessage(scen)
		assert.Nil(t, err)

		config, err := autochat.NewAutochatConfig(scen)
		assert.Nil(t, err)

		app := autochat.NewApplication(scen, message, config)

		t.Run("iterate akun sender file not exist", func(t *testing.T) {

			akunchan, err := app.IterateAkunSender()
			assert.Nil(t, err)

			count := 0
			for range akunchan {
				count++
			}
			assert.Zero(t, count)
		})

		t.Run("iterate akun sender file exist", func(t *testing.T) {

			akuns := []string{}
			for i := 0; i < 5; i++ {
				akuns = append(akuns, "pdcthoni@gmail.com|SilentIsMyMantra|IULIWGH6TIK3CZBKHGE27DBRLQ5LR5WQ")
			}

			akundata := strings.Join(akuns, "\n")
			remove := scen.CreateFile([]byte(akundata), scen.Path(config.AkunLoc))
			defer remove()

			akunchan, err := app.IterateAkunSender()
			assert.Nil(t, err)

			count := 0
			for range akunchan {
				count++
			}
			assert.Equal(t, count, len(akuns))
		})
	})
}
