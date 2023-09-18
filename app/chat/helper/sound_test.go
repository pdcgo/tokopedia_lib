package helper_test

import (
	"testing"
	"time"

	"github.com/pdcgo/tokopedia_lib/app/chat/config"
	"github.com/pdcgo/tokopedia_lib/app/chat/helper"
	"github.com/pdcgo/tokopedia_lib/scenario"
)

func TestSound(t *testing.T) {
	sound := helper.NewSoundPlayer(&config.AppConfig{
		Base: scenario.GetBaseTestAsset("chat_base"),
	})

	sound.Play()

	time.Sleep(time.Second * 6)
	sound.Pause()

	time.Sleep(time.Second * 3)

	sound.Play()
	time.Sleep(time.Second * 6)
}
