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

type AutosendConfig struct {
	CustomLastMessage string `json:"custom_last_message"`
}

type AutochatConfig struct {
	AkunLoc          string          `json:"akun_loc"`
	ShopLoc          string          `json:"shop_loc"`
	Concurrent       int             `json:"concurrent"`
	SendInterval     *Range          `json:"send_interval"`
	TypingInterval   *Range          `json:"typing_interval"`
	LimitMessageSend *Range          `json:"limit_message_send"`
	Autosend         *AutosendConfig `json:"autosend"`
}

var AutoConfigName = "autoconfig.json"

func writeConfig(fname string, autoconfig *AutochatConfig) error {
	b, err := json.Marshal(&autoconfig)
	if err != nil {
		return err
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return err
	}

	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(out.Bytes())
	return err
}

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
		Autosend: &AutosendConfig{},
	}

	fname := base.Path(AutoConfigName)
	// create if not exist
	if _, err := os.Stat(fname); errors.Is(err, os.ErrNotExist) {

		log.Println("creating autoconfig.json...")
		err = writeConfig(fname, &autoconfig)
		return &autoconfig, err
	}

	// load if exist
	f, err := os.ReadFile(fname)
	if err != nil {
		return &autoconfig, err
	}

	err = json.Unmarshal(f, &autoconfig)
	if err != nil {
		return &autoconfig, err
	}

	// rewrite
	checkAutoconfig := AutochatConfig{}
	err = json.Unmarshal(f, &checkAutoconfig)
	if err != nil {
		return &autoconfig, err
	}
	if checkAutoconfig.Autosend == nil {
		err = writeConfig(fname, &autoconfig)
		if err != nil {
			return &autoconfig, err
		}
	}

	return &autoconfig, nil
}
