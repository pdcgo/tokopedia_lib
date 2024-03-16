package helper

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/pdcgo/tokopedia_lib/app/chat/config"
)

type SoundPlayer struct {
	ctrl     *beep.Ctrl
	initOnce sync.Once
	format   beep.Format
}

func (s *SoundPlayer) Play() {
	s.initOnce.Do(func() {
		speaker.Init(s.format.SampleRate, s.format.SampleRate.N(time.Second/10))
		speaker.Play(s.ctrl)
	})

	speaker.Lock()
	s.ctrl.Paused = false
	speaker.Unlock()
}
func (s *SoundPlayer) Pause() {
	speaker.Lock()
	s.ctrl.Paused = true
	speaker.Unlock()
}

func NewSoundPlayer(config *config.AppConfig) *SoundPlayer {

	f, err := os.Open(config.Path("./data/sound/alert-chat.mp3"))
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}

	sound := SoundPlayer{
		ctrl:   ctrl,
		format: format,
	}
	return &sound
}
