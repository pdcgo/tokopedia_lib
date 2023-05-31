package model_public_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestParsingPDPLayout(t *testing.T) {
	fname := scenario.GetBaseTestAsset("api_response", "pdplayout.json")
	data, err := os.ReadFile(fname)
	assert.Nil(t, err)

	var hasil model_public.PdpGetLayout

	err = json.Unmarshal(data, &hasil)
	assert.Nil(t, err)
	t.Log(hasil.Components)
}
