package grab

import (
	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/grabber"
)

type GrabApp struct {
	base     *legacy_source.BaseConfig
	baseGrab *grabber.Grabber
}

func NewGrabApp(
	api *api_public.TokopediaApiPublic,
	base *legacy_source.BaseConfig,
	repo *mongorepo.ProductRepo,
) *GrabApp {

	baseGrab := grabber.CreateBaseGrabber(api, base, repo)

	return &GrabApp{
		base:     base,
		baseGrab: baseGrab,
	}
}

func (runner *GrabApp) Run() {

}
