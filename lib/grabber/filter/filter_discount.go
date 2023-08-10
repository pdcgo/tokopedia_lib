package filter

import (
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

func CreateFilterDiscount(markupConfig *legacy.LegacyMarkupConfig) FilterHandler {

	return func(layout *model_public.PdpGetlayoutQueryResp, pdp *model_public.PdpGetDataP2Resp) (bool, string, error) {

		if markupConfig.UsePriceDiscount {

			percentageAmount := layout.Data.PdpGetLayout.GetPercentageAmount()
			cek := percentageAmount != 0

			if cek {
				return cek, "filter discount", nil
			}
		}

		return false, "", nil
	}
}
