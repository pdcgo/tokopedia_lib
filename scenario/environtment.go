package scenario

import (
	"encoding/json"
	"errors"
	"net"
	"net/url"
	"os"
	"sync"
	"testing"
	"time"
)

func raw_connect(host string, port string) bool {
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		return false
	}
	if conn != nil {
		defer conn.Close()
		return true
	}

	return true
}

var isMongoRunning bool
var chekcMongo sync.Once
var CfgTest *ConfigTest

type ConfigTest struct {
	MongoUri string `json:"mongo_uri"`
}

func (cfg *ConfigTest) GetHostPort() (string, string) {
	u, _ := url.Parse(cfg.MongoUri)

	return u.Hostname(), u.Port()
}

func CheckConfig(t *testing.T) *ConfigTest {
	var hasil ConfigTest
	configpath := GetBaseTestAsset("database_test_config.json")
	if _, err := os.Stat(configpath); errors.Is(err, os.ErrNotExist) {

		t.Log("untuk testing berkaitan mongo harus config di ", configpath)
		return &hasil
	}
	data, _ := os.ReadFile(configpath)
	json.Unmarshal(data, &hasil)

	CfgTest = &hasil
	return &hasil
}

func CheckMongoActive(t *testing.T) {
	chekcMongo.Do(func() {
		cfg := CheckConfig(t)
		if cfg == nil {
			return
		}

		host, port := cfg.GetHostPort()
		isMongoRunning = raw_connect(host, port)
	})
	if !isMongoRunning {
		t.Skip("skipping database test because no database open")
	}
}
