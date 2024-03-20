package config

import (
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type AutoReplyItem struct {
	Delay   int    `json:"delay" yaml:"delay"`
	IsRegex bool   `json:"is_regex" yaml:"is_regex"`
	Pola    string `json:"pola" yaml:"pola"`
	Reply   string `json:"reply" yaml:"reply"`
}

func (item *AutoReplyItem) GetDelay() time.Duration {
	r := rand.Float32() * float32(item.Delay)
	return time.Second * time.Duration(r)
}

type AutoReplyConfigData struct {
	Active  bool             `json:"active" yaml:"active"`
	Replies []*AutoReplyItem `json:"replies" yaml:"replies"`
}

type AutoReplyConfig struct {
	Loc  string               `json:"-" yaml:"-"`
	Data *AutoReplyConfigData `json:"data" yaml:"data"`
}

func NewAutoReplyConfig(config *AppConfig) *AutoReplyConfig {

	loc := config.Path("data", "auto_reply.yml")
	areplyconfig := AutoReplyConfig{
		Loc:  loc,
		Data: &AutoReplyConfigData{},
	}

	data, _ := os.ReadFile(loc)
	yaml.Unmarshal(data, &areplyconfig)

	if areplyconfig.Data.Active {
		log.Println("menggunakan auto reply...")
	}

	return &areplyconfig
}

func (cfg *AutoReplyConfig) Save() error {

	f, err := os.OpenFile(cfg.Loc, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	err = yaml.NewEncoder(f).Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}

func (cfg *AutoReplyConfig) Find(text string) *AutoReplyItem {

	for _, reply := range cfg.Data.Replies {
		if reply.IsRegex {

			rx := regexp.MustCompile(reply.Pola)
			if rx.MatchString(text) {
				return reply
			}
		} else {

			if strings.Contains(
				strings.ToLower(text),
				strings.ToLower(reply.Pola),
			) {
				return reply
			}
		}
	}

	return nil
}
