package grab

import (
	"log"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/grab_handler"
	"github.com/pdcgo/tokopedia_lib/lib/grabber"
)

type GrabApp struct {
	api          *api_public.TokopediaApiPublic
	base         *legacy_source.BaseConfig
	cacheHandler *grab_handler.CacheProductHandler
}

func NewGrabApp(
	api *api_public.TokopediaApiPublic,
	base *legacy_source.BaseConfig,
	repo *mongorepo.ProductRepo,
) *GrabApp {

	cacheHandler := grab_handler.NewCacheProductHandler(repo)
	return &GrabApp{
		api:          api,
		base:         base,
		cacheHandler: cacheHandler,
	}
}

func (a *GrabApp) Run() error {
	return legacy.IterateGrabTaskers(a.base, func(tasker *legacy.GrabTasker) error {

		if tasker.Marketplace != legacy.MARKETPLACE_TASKER_TOKOPEDIA {
			return nil
		}

		var grab grabber.Grabber
		baseGrab := grabber.NewBaseGrabber(a.api, a.base, tasker, a.cacheHandler)

		switch tasker.Mode {

		case legacy.GRAB_MODE_CATEGORY:
			grab = grabber.NewCategoryGrabber(baseGrab)

		case legacy.GRAB_MODE_CATEGORY_CSV:
			grab = grabber.NewCategoryCsvGrabber(baseGrab)

		case legacy.GRAB_MODE_KEYWORD:
			grab = grabber.NewKeywordGrabber(baseGrab)

		case legacy.GRAB_MODE_PRODUCT_URL:
			grab = grabber.NewUrlGrabber(baseGrab)

		case legacy.GRAB_MODE_TOKO_USERNAME:
			grab = grabber.NewShopGrabber(baseGrab)
		}

		if grab != nil {
			return grab.Run()
		}

		log.Printf("[ not supported ] %s - %s", tasker.Marketplace, tasker.Mode)
		return nil
	})
}
