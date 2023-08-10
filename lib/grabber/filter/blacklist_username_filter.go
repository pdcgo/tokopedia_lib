package filter

import (
	"errors"
	"strings"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/helper"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

var ErrBlacklistUsername = errors.New("filter blacklist username")

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
		for _, username := range blUsername {
			if strings.Contains(username, shopDomain) {
				return true, "filter title", nil
			}
		}

		return false, "", nil
	}
}
