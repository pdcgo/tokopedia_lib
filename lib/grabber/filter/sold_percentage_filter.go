package filter

import (
	"strconv"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

func CreateSoldPercentageFilter(base *legacy_source.BaseConfig) FilterHandler {

	grabBasic := legacy.NewGrabBasic(base)

	return func(layout *model_public.PdpGetlayoutQueryResp, pdp *model_public.PdpGetDataP2Resp) (cek bool, reason string, err error) {

		productSold, _ := strconv.Atoi(layout.Data.PdpGetLayout.BasicInfo.TxStats.CountSold)
		productSuccessSold, _ := strconv.Atoi(layout.Data.PdpGetLayout.BasicInfo.TxStats.TransactionSuccess)
		soldPercentage := (float64(productSuccessSold) / float64(productSold)) * 100
		if grabBasic.Prosentase > int(soldPercentage) {
			return true, "filter prosentase", nil
		}

		return false, "", nil
	}
}
