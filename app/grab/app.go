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
	base     *legacy_source.BaseConfig
	baseGrab *grabber.BaseGrabber
}

func NewGrabApp(
	api *api_public.TokopediaApiPublic,
	base *legacy_source.BaseConfig,
	repo *mongorepo.ProductRepo,
) *GrabApp {

	cacheHandler := grab_handler.NewCacheProductHandler(repo)
	baseGrab := grabber.NewBaseGrabber(api, base, &legacy.GrabTasker{}, cacheHandler)

	return &GrabApp{
		base:     base,
		baseGrab: baseGrab,
	}
}

func (runner *GrabApp) Run() error {
	return legacy.IterateGrabTaskers(runner.base, func(tasker *legacy.GrabTasker) error {
		var grab grabber.Grabber
		runner.baseGrab.GrabTasker = tasker

		mode := tasker.Mode
		switch mode {

		case legacy.GRAB_MODE_CATEGORY:
			grab = grabber.NewCategoryGrabber(runner.baseGrab)

		case legacy.GRAB_MODE_CATEGORY_CSV:
			grab = grabber.NewCategoryCsvGrabber(runner.baseGrab)

		case legacy.GRAB_MODE_KEYWORD:
			grab = grabber.NewKeywordGrabber(runner.baseGrab)

		case legacy.GRAB_MODE_PRODUCT_URL:
			grab = grabber.NewUrlGrabber(runner.baseGrab)

		case legacy.GRAB_MODE_TOKO_USERNAME:
			grab = grabber.NewShopGrabber(runner.baseGrab)
		}

		if grab != nil {
			return grab.Run()
		}

		log.Printf("[ not supported ] %s - %s", tasker.Marketplace, tasker.Mode)
		return nil
	})
}
