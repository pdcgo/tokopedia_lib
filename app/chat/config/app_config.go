package config

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Base                string     `yaml:"-"`
	Port                string     `yaml:"-"`
	Host                string     `yaml:"-"`
	DebugMode           bool       `yaml:"debug_mode"`
	OpenBrowser         bool       `yaml:"open_browser"`
	GsheetUrl           string     `yaml:"google_sheet_url"`
	SyncCommandInterval [2]float32 `yaml:"sync_command_interval"`
}

func (c *AppConfig) GetSync() time.Duration {
	min := c.SyncCommandInterval[0]
	max := c.SyncCommandInterval[1]
	r := min + rand.Float32()*(max-min)
	return time.Second * time.Duration(r)
}

func (app *AppConfig) Path(elem ...string) string {
	elem = append([]string{app.Base}, elem...)
	return filepath.Join(elem...)
}

func NewAppConfig(base string) *AppConfig {
	fconfig := filepath.Join(base, "chat_config.yml")

	config := AppConfig{
		Base:                base,
		OpenBrowser:         true,
		DebugMode:           false,
		Port:                "5003",
		Host:                "localhost",
		SyncCommandInterval: [2]float32{5, 30},
	}

	if _, err := os.Stat(fconfig); os.IsNotExist(err) {
		f, err := os.OpenFile(fconfig, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
		if err != nil {
			log.Panicln(err)
		}
		defer f.Close()
		yaml.NewEncoder(f).Encode(&config)
	} else {
		data, _ := os.ReadFile(fconfig)
		yaml.Unmarshal(data, &config)
	}

	return &config
}
