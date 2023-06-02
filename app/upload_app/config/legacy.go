package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type DatabaseConfig struct {
	DbURI  string `json:"DB_URI"`
	DbName string `json:"DB_NAME"`
}
type UploadBaseConfig struct {
	Database     DatabaseConfig `json:"database"`
	basePath     string
	relativeBase string
}

func NewUploadConfigBase(base string) *UploadBaseConfig {
	abspath, _ := filepath.Abs(base)
	hasil := UploadBaseConfig{
		basePath:     abspath,
		relativeBase: base,
	}

	path := filepath.Join(abspath, "data", "config.json")
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&hasil)
	if err != nil {
		panic(err)
	}

	if hasil.Database.DbURI == "" {
		hasil.Database.DbURI = "mongodb://localhost:27017"
	}

	return &hasil
}
