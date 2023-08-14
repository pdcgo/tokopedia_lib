package filter

import (
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

func CreateFilterDiscount(markupConfig *legacy.LegacyMarkupConfig) FilterHandler {

	return func(layout *model_public.PdpGetlayoutQueryResp, pdp *model_public.PdpGetDataP2Resp) (bool, string, error) {

		if markupConfig.UsePriceDiscount {

			percentageAmount, err := layout.Data.PdpGetLayout.GetPercentageAmount()
			if err != nil {
				return true, "filter discount", err
			}

			cek := percentageAmount != 0
			if cek {
				return cek, "filter discount", nil
			}
		}

		return false, "", nil
	}
}
