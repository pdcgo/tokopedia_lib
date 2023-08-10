package filter

import (
	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
)

func NewGrabFilterBundle(
	api *api_public.TokopediaApiPublic,
	base *legacy_source.BaseConfig,
	filterText *legacy_source.FilterText,
	grabBasic *legacy.GrabBasic,
	grabTokopedia *legacy.GrabTokopedia,
	markupConfig *legacy.LegacyMarkupConfig,
) []FilterHandler {

	return []FilterHandler{
		CreateTitleFilter(filterText),
		CreateSoldFilter(grabBasic),
		CreateSoldPercentageFilter(grabBasic),
		CreateStockFilter(grabBasic),
		CreateFilterDiscount(markupConfig),
		CreatePointFilter(api, grabTokopedia),
		CreateBlacklistUsernameFilter(base, grabBasic),
		CreateLastLoginFilter(grabTokopedia),
		CreateLastReviewFilter(api, grabBasic),
	}
}
