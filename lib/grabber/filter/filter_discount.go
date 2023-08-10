package filter

import (
	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/tokopedia_lib/lib/helper"
	"github.com/pdcgo/tokopedia_lib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

func CreateFilterDiscount(base *legacy_source.BaseConfig) FilterHandler {
	grabBasic := legacy.NewGrabBasic(base)

	return func(layout *model_public.PdpGetlayoutQueryResp, pdp *model_public.PdpGetDataP2Resp) (bool, string, error) {
		components := helper.ParseProductLayoutComponents(layout.Data.PdpGetLayout.Components)

		if grabBasic.UsePriceDiscount {
			discountPercentage := components.ProductContent.Data[len(components.ProductContent.Data)-1].Campaign.DiscountPercentage
			cek := discountPercentage != 0
			if cek {
				return cek, "filter discount", nil
			}
		}

		return false, "", nil
	}
}
