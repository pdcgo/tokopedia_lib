package main

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/pdcgo/v2_gots_sdk"
)

func TestApi(t *testing.T) {
	scenario.RunWebSdk(t, func(sdk *v2_gots_sdk.ApiSdk, sendApi scenario.SendRequest) error {
		return nil
	})
}
