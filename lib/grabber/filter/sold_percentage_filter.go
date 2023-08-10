package filter

import (
	"strconv"

	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

func CreateSoldPercentageFilter(grabBasic *legacy.GrabBasic) FilterHandler {

	return func(layout *model_public.PdpGetlayoutQueryResp, pdp *model_public.PdpGetDataP2Resp) (cek bool, reason string, err error) {

		productSold, err := strconv.Atoi(layout.Data.PdpGetLayout.BasicInfo.TxStats.CountSold)
		if err != nil {
			return true, "filter prosentase", err
		}
		productSuccessSold, err := strconv.Atoi(layout.Data.PdpGetLayout.BasicInfo.TxStats.TransactionSuccess)
		if err != nil {
			return true, "filter prosentase", err
		}
		soldPercentage := (float64(productSuccessSold) / float64(productSold)) * 100
		if grabBasic.Prosentase > int(soldPercentage) {
			return true, "filter prosentase", nil
		}

		return false, "", nil
	}
}
