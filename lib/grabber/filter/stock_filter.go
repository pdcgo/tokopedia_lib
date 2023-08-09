package filter

import (
	"strconv"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/helper"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

func CreateStockFilter(base *legacy_source.BaseConfig) FilterHandler {

	grabBasic := legacy.NewGrabBasic(base)

	return func(layout *model_public.PdpGetlayoutQueryResp, pdp *model_public.PdpGetDataP2Resp) (cek bool, reason string, err error) {

		productLayout := helper.ParseProductLayoutComponents(layout.Data.PdpGetLayout.Components)
		stock := productLayout.ProductContent.Data[0].Campaign.OriginalStock
		if stock == 0 {
			stock, err = strconv.Atoi(productLayout.ProductContent.Data[0].Stock.Value)
			if err != nil {
				return true, "filter stock", err
			}
		}

		if grabBasic.Stock > stock {
			return true, "filter stock", nil
		}

		return false, "", nil
	}
}
