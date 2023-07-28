package grabber_test

import (
	"context"
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/grabber"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestUrlGrabber(t *testing.T) {
	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	baseConfig := &legacy_source.BaseConfig{
		BaseData: "../..",
	}
	database := scenario.GetMongoDatabase(t)
	productRepo := mongorepo.NewProductRepo(ctx, database)
	baseGrabber := grabber.NewBaseGrabber(api, baseConfig, productRepo)

	t.Run("test url grabber", func(t *testing.T) {
		urls := []string{
			"https://www.tokopedia.com/wellborn/wellborn-fantasy-black-t-shirt-s",
		}
		grabber := grabber.NewUrlGrabber(baseGrabber, urls)
		grabber.Run()
	})
}
