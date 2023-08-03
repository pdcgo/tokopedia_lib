package filter

import (
	"strconv"
	"time"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

func CreateLastLoginFilter(base *legacy_source.BaseConfig) FilterHandler {
	return func(layout *model_public.PdpGetlayoutQueryResp, pdp *model_public.PdpGetDataP2Resp) (cek bool, reason string, err error) {
		grabTokped := legacy.NewGrabTokopedia(base)
		if !grabTokped.LastLoginActive {
			return false, "", nil
		}
		lastLogin, err := strconv.Atoi(pdp.Data.PdpGetData.ShopInfo.ShopLastActive)
		if err != nil {
			return true, "filter last login", err
		}

		t := time.Now()
		filterLastLogin := t.AddDate(0, 0, -grabTokped.LastLoginDays)
		if filterLastLogin.Unix() > int64(lastLogin) {
			return true, "filter last review", nil
		}
		return false, "", nil
	}
}
