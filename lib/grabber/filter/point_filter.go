package filter

import (
	"strconv"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

func CreatePointFilter(api *api_public.TokopediaApiPublic, base *legacy_source.BaseConfig) FilterHandler {
	return func(layout *model_public.PdpGetlayoutQueryResp, pdp *model_public.PdpGetDataP2Resp) (cek bool, reason string, err error) {

		shopId := layout.Data.PdpGetLayout.BasicInfo.ShopID
		shopIdInt, err := strconv.Atoi(shopId)
		if err != nil {
			return true, "filter point", err
		}

		variable := model_public.ShopStatisticQueryVar{
			ShopID:    shopIdInt,
			ShopIDStr: shopId,
		}

		stats, err := api.ShopStatisticQuery(&variable)
		if err != nil {
			return true, "filter point", err
		}

		point, err := strconv.Atoi(stats.Data.ShopReputation[0].ScoreMap)
		if err != nil {
			return true, "filter point", err
		}

		grabTokped := legacy.NewGrabTokopedia(base)
		isLwPointMin := point < grabTokped.Point[0]
		isGtPointMax := point > grabTokped.Point[1]

		if isLwPointMin || isGtPointMax {
			return true, "filter point", nil
		}

		return false, "", nil
	}
}
