package grabber

import (
	"net/url"
	"strings"

	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

func CreateGrabSearchVar(grabTokopedia *legacy.GrabTokopedia) *model_public.SearchProductVar {

	locs := grabTokopedia.Query.Fcity
	shippings := grabTokopedia.Query.Shipping
	conditions := strings.Split(grabTokopedia.Query.Condition, ",")
	shopRating := strings.Split(grabTokopedia.Query.Rt, ",")

	shopTier := []string{}
	if grabTokopedia.Query.Official {
		shopTier = append(shopTier, "2")
	}

	if grabTokopedia.Query.Goldmerchant {
		shopTier = append(shopTier, "3")
	}

	searchVar := model_public.NewSearchProductVar()
	searchVar.Sort = grabTokopedia.Query.Ob
	searchVar.PriceMin = grabTokopedia.Query.Pmin
	searchVar.PriceMax = grabTokopedia.Query.Pmax
	searchVar.PreOrder = grabTokopedia.Query.Preorder
	searchVar.Locations = url.QueryEscape(strings.Join(locs, ","))
	searchVar.Rate = url.QueryEscape(strings.Join(shopRating, "#"))
	searchVar.Condition = url.QueryEscape(strings.Join(conditions, "#"))
	searchVar.Shipping = url.QueryEscape(strings.Join(shippings, "#"))
	searchVar.ShopTier = url.QueryEscape(strings.Join(shopTier, "#"))

	return searchVar
}
