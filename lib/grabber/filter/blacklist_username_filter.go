package filter

import (
	"errors"
	"os"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/helper"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

func CreateBlacklistUsernameFilter(base *legacy_source.BaseConfig, grabBasic *legacy.GrabBasic) FilterHandler {

	var loadErr error
	blacklist := grabBasic.BlacklistUsername

	if blacklist.Active {
		fname := blacklist.Tokopedia.Filename

		if fname != "" {
			fname = base.Path(fname)
			blacklist.Tokopedia.Data, loadErr = helper.FileLoadLineString(fname)

			if os.IsNotExist(loadErr) {
				loadErr = errors.New("blacklist toko " + fname + "tidak ditemukan")
			}
		}
	}

	return func(layout *model_public.PdpGetlayoutQueryResp, pdp *model_public.PdpGetDataP2Resp) (cek bool, reason string, err error) {
		if loadErr != nil {
			return true, "", loadErr
		}

		shopDomain := pdp.Data.PdpGetData.ShopInfo.ShopCore.Domain
		if len(blacklist.Tokopedia.Data) > 0 {
			for _, name := range blacklist.Tokopedia.Data {
				if name == shopDomain {
					return true, "filter shopname", nil
				}
			}
		}

		return false, "", nil
	}
}
