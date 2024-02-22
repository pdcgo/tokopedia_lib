package autochat

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"os"

	"github.com/pdcgo/common_conf/pdc_application"
)

type Range struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

func (i *Range) Get() int {
	return rand.Intn(i.Max-i.Min) + i.Min
}

type AutochatConfig struct {
	AkunLoc          string `json:"akun_loc"`
	ShopLoc          string `json:"shop_loc"`
	Concurrent       int    `json:"concurrent"`
	SendInterval     *Range `json:"send_interval"`
	TypingInterval   *Range `json:"typing_interval"`
	LimitMessageSend *Range `json:"limit_message_send"`
}

var AutoConfigName = "autoconfig.json"

func NewAutochatConfig(base pdc_application.BaseApplication) (*AutochatConfig, error) {

	autoconfig := AutochatConfig{
		AkunLoc:    "akun.txt",
		ShopLoc:    "shop.txt",
		Concurrent: 5,
		SendInterval: &Range{
			Min: 3,
			Max: 5,
		},
		TypingInterval: &Range{
			Min: 1,
			Max: 2,
		},
		LimitMessageSend: &Range{
			Min: 4,
			Max: 6,
		},
	}

	fname := base.Path(AutoConfigName)
	// create if not exist
	if _, err := os.Stat(fname); errors.Is(err, os.ErrNotExist) {

		log.Println("creating autoconfig.json...")
		b, err := json.Marshal(&autoconfig)
		if err != nil {
			return &autoconfig, err
		}

		var out bytes.Buffer
		err = json.Indent(&out, b, "", "  ")
		if err != nil {
			return &autoconfig, err
		}

		f, err := os.Create(fname)
		if err != nil {
			return &autoconfig, err
		}

		_, err = f.Write(out.Bytes())
		return &autoconfig, err
	}

	// load if exist
	f, err := os.ReadFile(fname)
	if err != nil {
		return &autoconfig, err
	}

	err = json.Unmarshal(f, &autoconfig)
	return &autoconfig, err
}
