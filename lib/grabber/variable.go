package grabber

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

func arrayConverter(datas []interface{}) []string {
	results := make([]string, len(datas))
	for _, data := range datas {
		switch value := data.(type) {
		case int:
			results = append(results, fmt.Sprint(value))
		case string:
			results = append(results, value)
		}
	}
	return results
}

func CreateGrabSearchVar(base *legacy_source.BaseConfig) *model_public.SearchProductVar {

	grabConfig := legacy.NewGrabTokopedia(base)

	locs := arrayConverter(grabConfig.Query.Fcity)
	shippings := arrayConverter(grabConfig.Query.Shipping)

	conditions := strings.Split(grabConfig.Query.Condition, ",")
	shopRating := strings.Split(grabConfig.Query.Rt, ",")

	shopTier := []string{}
	if grabConfig.Query.Official {
		shopTier = append(shopTier, "2")
	}

	if grabConfig.Query.Goldmerchant {
		shopTier = append(shopTier, "3")
	}

	params := model_public.SearchProductVar{
		Device:         "desktop",
		Sort:           grabConfig.Query.Ob,
		Page:           1,
		Rows:           100,
		UserDistrictID: "176",
		UserCityID:     "2274",
		Related:        true,
		Scheme:         "https",
		SafeSearch:     false,
		TopadsBucket:   true,
		Source:         "search",
		PriceMin:       grabConfig.Query.Pmin,
		PriceMax:       grabConfig.Query.Pmax,
		PreOrder:       grabConfig.Query.Preorder,
		Locations:      url.QueryEscape(strings.Join(locs, ",")),
		Rate:           url.QueryEscape(strings.Join(shopRating, "#")),
		Condition:      url.QueryEscape(strings.Join(conditions, "#")),
		Shipping:       url.QueryEscape(strings.Join(shippings, "#")),
		ShopTier:       url.QueryEscape(strings.Join(shopTier, "#")),
	}
	return &params
}
