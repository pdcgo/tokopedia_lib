package autochat_test

import (
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

		limit := 5
		fname := scen.Path(config.AkunLoc)
		app := autochat.NewApplication(scen, message, config)

		t.Run("get akuns file not exist", func(t *testing.T) {

			akuns, err := autochat.GetAkuns(fname)
			assert.Nil(t, err)
			assert.Zero(t, len(akuns))
		})

		t.Run("save akuns", func(t *testing.T) {

			akuns := []*autochat.Akun{}
			for i := 0; i < limit; i++ {
				akuns = append(akuns, &autochat.Akun{
					Username: "pdcthoni@gmail.com",
					Password: "SilentIsMyMantra",
					Secret:   "IULIWGH6TIK3CZBKHGE27DBRLQ5LR5WQ",
				})
			}

			err := autochat.SaveAkuns(fname, akuns)
			assert.Nil(t, err)
		})

		t.Run("iterate akun sender", func(t *testing.T) {
			akunchan, err := app.IterateAkunSender()
			assert.Nil(t, err)

			count := 0
			for range akunchan {
				count++
			}
			assert.Equal(t, count, limit)
		})
	})
}
