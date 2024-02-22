package autochat_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/autochat"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {

	scen := scenario.NewScenario(t)
	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {

		t.Run("test patterns n replies empty", func(t *testing.T) {

			automessage, err := autochat.NewAutochatMessage(scen)
			assert.Nil(t, err)
			assert.Empty(t, automessage.Patterns)
			assert.Empty(t, automessage.Messages)
		})

		t.Run("test patterns n replies exist", func(t *testing.T) {
			patterns := []string{}
			replies := []string{}
			for i := 0; i < 10; i++ {
				patterns = append(patterns, fmt.Sprintf("test pattern %d", i))
				replies = append(replies, fmt.Sprintf("test reply %d", i))
			}

			patternData := strings.Join(patterns, "\n")
			patternFname := scen.Path(autochat.PatternDir, "test.txt")
			patternRemove := scen.CreateFile([]byte(patternData), patternFname)
			defer patternRemove()

			replyData := strings.Join(replies, "\n")
			replyFname := scen.Path(autochat.ReplyDir, "test.txt")
			replyRemove := scen.CreateFile([]byte(replyData), replyFname)
			defer replyRemove()

			automessage, err := autochat.NewAutochatMessage(scen)
			assert.Nil(t, err)
			assert.NotEmpty(t, automessage.Patterns)
			assert.NotEmpty(t, automessage.Messages)

			t.Run("test in pattern", func(t *testing.T) {

				t.Run("test in pattern true", func(t *testing.T) {
					assert.True(t, automessage.InPattern("test pattern 1"))
				})

				t.Run("test in pattern false", func(t *testing.T) {
					assert.True(t, automessage.InPattern("test pattern 11"))
				})
			})

			t.Run("test get message", func(t *testing.T) {
				assert.Contains(t, automessage.GetMessage(), "test reply")
			})
		})
	})
}
