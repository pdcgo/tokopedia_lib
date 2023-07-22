package grabber

import (
	"net/url"
	"strings"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/filter"
	"github.com/pdcgo/tokopedia_lib/lib/grab_handler"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type Grabber struct {
	Api          *api_public.TokopediaApiPublic
	Filter       *filter.BaseFilter
	CacheHandler *grab_handler.CacheProductHandler
}

func (grab *Grabber) generateProductSearchParams() *model_public.SearchProductVar {
	locs := arrayConverter(grab.Filter.GrabTokopedia.Query.Fcity)
	shippings := arrayConverter(grab.Filter.GrabTokopedia.Query.Shipping)

	conditions := strings.Split(grab.Filter.GrabTokopedia.Query.Condition, ",")
	shopRating := strings.Split(grab.Filter.GrabTokopedia.Query.Rt, ",")

	shopTier := []string{}
	if grab.Filter.GrabTokopedia.Query.Official {
		shopTier = append(shopTier, "2")
	}
	if grab.Filter.GrabTokopedia.Query.Goldmerchant {
		shopTier = append(shopTier, "3")
	}

	params := model_public.SearchProductVar{
		Device:         "desktop",
		Sort:           grab.Filter.GrabTokopedia.Query.Ob,
		Page:           1,
		Rows:           100,
		UserDistrictID: "176",
		UserCityID:     "2274",
		Related:        true,
		Scheme:         "https",
		SafeSearch:     false,
		TopadsBucket:   true,
		Source:         "search",
		PriceMin:       grab.Filter.GrabTokopedia.Query.Pmin,
		PriceMax:       grab.Filter.GrabTokopedia.Query.Pmax,
		PreOrder:       grab.Filter.GrabTokopedia.Query.Preorder,
		Locations:      url.QueryEscape(strings.Join(locs, ",")),
		Rate:           url.QueryEscape(strings.Join(shopRating, "#")),
		Condition:      url.QueryEscape(strings.Join(conditions, "#")),
		Shipping:       url.QueryEscape(strings.Join(shippings, "#")),
		ShopTier:       url.QueryEscape(strings.Join(shopTier, "#")),
	}
	return &params
}

func CreateBaseGrabber(api *api_public.TokopediaApiPublic, base *legacy_source.BaseConfig, repo *mongorepo.ProductRepo) *Grabber {
	filter := filter.CreateBaseFilter(api, base)
	cacheHandler := grab_handler.NewCacheProductHandler(repo)
	return &Grabber{
		Api:          api,
		Filter:       filter,
		CacheHandler: cacheHandler,
	}
}
