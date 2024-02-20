package model_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestModelShopInfoById(t *testing.T) {
	fname := scenario.GetBaseTestAsset("api_response", "shopinfobyid_err.json")
	sourcebytes, err := os.ReadFile(fname)
	assert.Nil(t, err)

	var result model.ShopInfoByIDRes
	err = json.Unmarshal(sourcebytes, &result)
	assert.Nil(t, err)
	assert.NotEmpty(t, result)
	assert.NotEmpty(t, result.Errors)
	assert.True(t, result.Errors.IsNotAuthorized())
}
