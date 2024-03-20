package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type InitConfig struct {
	Loc         string `json:"-" yaml:"-"`
	FirstInit   bool   `json:"-" yaml:"-"`
	ActiveGroup string `json:"active_group" yaml:"active_group"`
}

func NewInitConfig(config *AppConfig) *InitConfig {
	loc := config.Path("chat_init_config.yml")

	initcfg := InitConfig{
		Loc:         loc,
		FirstInit:   true,
		ActiveGroup: "",
	}

	data, _ := os.ReadFile(loc)
	yaml.Unmarshal(data, &initcfg)

	return &initcfg
}

func (cfg *InitConfig) CheckGroup(group string) bool {

	if cfg.FirstInit {
		cfg.FirstInit = false
		return false
	}

	return group == cfg.ActiveGroup
}

func (cfg *InitConfig) SetGroup(group string) error {
	cfg.ActiveGroup = group
	return cfg.Save()
}

func (cfg *InitConfig) Save() error {

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
