package filter

import (
	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/helper"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"golang.org/x/exp/slices"
)

func CreateBlacklistUsernameFilter(base *legacy_source.BaseConfig) FilterHandler {
	return func(layout *model_public.PdpGetlayoutQueryResp, pdp *model_public.PdpGetDataP2Resp) (cek bool, reason string, err error) {
		grabBasic := legacy.NewGrabBasic(base)
		if !grabBasic.BlacklistUsername.Active {
			return false, "", nil
		}

		var blUsername []string

		fname := grabBasic.BlacklistUsername.Tokopedia.Filename
		if fname != "" {
			pathfile := base.Path(fname)
			results, err := helper.FileLoadLineString(pathfile)
			if err != nil {
				return true, "filter blacklist username", err
			}
			blUsername = results
		} else {
			blUsername = grabBasic.BlacklistUsername.Tokopedia.Data
		}

		shopDomain := pdp.Data.PdpGetData.ShopInfo.ShopCore.Domain
		if slices.Contains(blUsername, shopDomain) {
			return true, "filter blacklist username", nil
		}

		return false, "", nil
	}
}
