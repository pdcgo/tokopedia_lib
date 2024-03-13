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

		limit := 5
		config, err := autochat.NewAutochatConfig(scen)
		assert.Nil(t, err)

		t.Run("get akuns file not exist", func(t *testing.T) {

			akundata, err := autochat.NewAkunData(scen, config)
			assert.Nil(t, err)
			assert.Zero(t, len(akundata.Data))
		})

		t.Run("save akuns", func(t *testing.T) {

			akundata, err := autochat.NewAkunData(scen, config)
			assert.Nil(t, err)

			for i := 0; i < limit; i++ {
				akundata.Data = append(akundata.Data, &autochat.Akun{
					Username: "pdcthoni@gmail.com",
					Password: "SilentIsMyMantra",
					Secret:   "IULIWGH6TIK3CZBKHGE27DBRLQ5LR5WQ",
				})
			}
			err = akundata.Save()
			assert.Nil(t, err)
		})

		t.Run("iterate akun sender", func(t *testing.T) {

			akundata, err := autochat.NewAkunData(scen, config)
			assert.Nil(t, err)

			count := 0
			err = akundata.IterateAkunSender(nil, func(akun *autochat.Akun, sender *autochat.AutochatSender) error {
				count++
				return nil
			})
			assert.Nil(t, err)
			assert.Equal(t, count, limit)
		})
	})
}
