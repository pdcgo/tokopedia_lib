package config_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/config"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

type UploadConfig struct {
	UseMapper bool `json:"use_mapper"`
}

func (cfg *UploadConfig) KeyConfig() string {
	return "UPLOAD_CONFIG"
}

func TestConfigRepo(t *testing.T) {

	db := scenario.GetDb()
	configrepo := config.NewConfigRepo(db)

	var config UploadConfig

	err := configrepo.GetConfig(&config)
	assert.Nil(t, err)

	config.UseMapper = true
	err = configrepo.Save(&config)
	assert.Nil(t, err)
}
