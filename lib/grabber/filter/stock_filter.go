package filter

import (
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

func CreateStockFilter(grabBasic *legacy.GrabBasic) FilterHandler {

	return func(layout *model_public.PdpGetlayoutQueryResp, pdp *model_public.PdpGetDataP2Resp) (cek bool, reason string, err error) {

		stock, err := layout.Data.PdpGetLayout.GetStock()
		if err != nil {
			return true, "filter stock", err
		}

		if grabBasic.Stock > stock {
			return true, "filter stock", nil
		}

		return false, "", nil
	}
}
