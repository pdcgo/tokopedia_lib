package filter

import (
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type FilterHandler func(layout *model_public.PdpGetlayoutQueryResp, pdp *model_public.PdpGetDataP2Resp) (cek bool, reason string, err error)

func NewFilterItem(filters ...FilterHandler) FilterHandler {
	return func(layout *model_public.PdpGetlayoutQueryResp, pdp *model_public.PdpGetDataP2Resp) (cek bool, reason string, err error) {

		for _, filter := range filters {

			cek, reason, err = filter(layout, pdp)

			if err != nil {
				return true, reason, err
			}

			if cek {
				return cek, reason, err
			}
		}

		return false, "", nil
	}
}
