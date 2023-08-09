package filter

import (
	"strconv"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

func CreateSoldFilter(base *legacy_source.BaseConfig) FilterHandler {

	grabBasic := legacy.NewGrabBasic(base)

	return func(layout *model_public.PdpGetlayoutQueryResp, pdp *model_public.PdpGetDataP2Resp) (cek bool, reason string, err error) {

		sold, err := strconv.Atoi(layout.Data.PdpGetLayout.BasicInfo.TxStats.CountSold)
		if err != nil {
			return true, "filter sold", err
		}

		if grabBasic.Penjualan > sold {
			return true, "filter sold", nil
		}

		return false, "", nil
	}
}
