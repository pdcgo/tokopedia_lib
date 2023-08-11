package iterator_test

import (
	"context"
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/iterator"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/stretchr/testify/assert"
)

func TestBatchLayoutIterator(t *testing.T) {

	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	urls := []string{
		"https://www.tokopedia.com/milenialbook-1/novel-d-bijis",
		"https://www.tokopedia.com/indomieofficial/10-pcs-indomie-goreng-spesial",
		"https://www.tokopedia.com/baseus/baseus-true-wireless-bluetooth-earphone-mini-earbuds-tws-wm01-ungu",
	}

	t.Run("test batch layout iterator ok", func(t *testing.T) {

		ctx := context.Background()
		layouts := []*model_public.PdpGetlayoutQueryResp{}

		err := iterator.IterateBatchLayout(api, ctx, urls, func(layout *model_public.PdpGetlayoutQueryResp) error {
			layouts = append(layouts, layout)
			return nil
		})

		assert.Nil(t, err)
		assert.Equal(t, len(urls), len(layouts))

	})

	t.Run("test batch layout iterator context cancel", func(t *testing.T) {

		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		layouts := []*model_public.PdpGetlayoutQueryResp{}

		err := iterator.IterateBatchLayout(api, ctx, urls, func(layout *model_public.PdpGetlayoutQueryResp) error {
			layouts = append(layouts, layout)
			return nil
		})

		assert.Nil(t, err)
		assert.Equal(t, 0, len(layouts))

	})
}
