package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type InitConfig struct {
	Loc         string `json:"-"`
	ActiveGroup string `json:"active_group"`
}

func (cfg *InitConfig) Save() error {
	f, err := os.OpenFile(cfg.Loc, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)

	if err != nil {
		log.Panicln(err)
	}
	defer f.Close()
	err = yaml.NewEncoder(f).Encode(cfg)
	if err != nil {
		log.Panicln(err)
	}
	return nil
}

func NewInitConfig(config *AppConfig) *InitConfig {
	loc := config.FullPath("chat_init_config.yml")

	initcfg := InitConfig{
		Loc:         loc,
		ActiveGroup: "",
	}

	data, _ := os.ReadFile(loc)
	yaml.Unmarshal(data, &initcfg)

	return &initcfg
}
