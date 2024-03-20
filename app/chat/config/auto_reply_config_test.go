package config_test

import (
	"testing"
	"time"

	"github.com/pdcgo/tokopedia_lib/app/chat/config"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestAutoReplyConfig(t *testing.T) {

	scen := scenario.NewScenario(t)
	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {

		autoconfig := config.NewAppConfig(dirbase)
		areplyconfig := config.NewAutoReplyConfig(autoconfig)

		t.Run("test save", func(t *testing.T) {

			areplyconfig.Data.Active = true
			areplyconfig.Data.Replies = []*config.AutoReplyItem{
				{
					Delay: 3,
					Pola:  "makan",
					Reply: "test makan",
				},
				{
					Delay:   3,
					Pola:    "ayam",
					IsRegex: true,
					Reply:   "test ayam",
				},
			}
			err := areplyconfig.Save()
			assert.Nil(t, err)

			newareplyconfig := config.NewAutoReplyConfig(autoconfig)
			assert.NotEmpty(t, newareplyconfig.Data.Replies)
		})

		t.Run("test find reply", func(t *testing.T) {

			t.Run("test not found", func(t *testing.T) {
				reply := areplyconfig.Find("payung")
				assert.Nil(t, reply)
			})

			t.Run("test found", func(t *testing.T) {
				reply := areplyconfig.Find("makan")
				assert.Equal(t, reply.Reply, "test makan")

				delay := reply.GetDelay()
				assert.Greater(t, delay, time.Duration(0))
				assert.LessOrEqual(t, delay, time.Second*3)
			})

			t.Run("test regex found", func(t *testing.T) {
				reply := areplyconfig.Find("ayam")
				assert.Equal(t, reply.Reply, "test ayam")

				delay := reply.GetDelay()
				assert.Greater(t, delay, time.Duration(0))
				assert.LessOrEqual(t, delay, time.Second*3)
			})
		})
	})
}
