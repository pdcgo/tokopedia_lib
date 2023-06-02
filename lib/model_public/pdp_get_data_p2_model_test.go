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

	var foundmedia bool
	for _, comp := range hasil.Components {
		switch component := comp.(type) {
		case *model_public.MediaComponent:
			t.Log(component)
			foundmedia = true
		}
	}
	assert.True(t, foundmedia)

}
