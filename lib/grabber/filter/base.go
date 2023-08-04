package filter

import (
	"context"
	"errors"

	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type FilterHandler func(layout *model_public.PdpGetlayoutQueryResp, pdp *model_public.PdpGetDataP2Resp) (cek bool, reason string, err error)

var ErrFilterCancel = errors.New("filter canceled")

func NewFilterItem(ctx context.Context, filters ...FilterHandler) FilterHandler {

	return func(layout *model_public.PdpGetlayoutQueryResp, pdp *model_public.PdpGetDataP2Resp) (cek bool, reason string, err error) {

		for _, filter := range filters {
			select {
			case <-ctx.Done():
				return true, "", ErrFilterCancel
			default:
				cek, reason, err = filter(layout, pdp)

				if err != nil {
					return true, reason, err
				}

				if cek {
					return cek, reason, err
				}
			}

		}

		return false, "", nil
	}
}
