package filter

import (
	"errors"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/helper"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

var ErrLimiterReached = errors.New("limit was reached")

func CreateLimiter(base *legacy_source.BaseConfig) (FilterHandler, *helper.Limiter) {

	grabBasic := legacy.NewGrabBasic(base)
	limiter := helper.NewLimiter(int32(grabBasic.LimitGrab))
	// count := int32(grabBasic.LimitGrab)

	// addCount := func() {
	// 	atomic.AddInt32(&count, -1)
	// }

	return func(layout *model_public.PdpGetlayoutQueryResp, pdp *model_public.PdpGetDataP2Resp) (cek bool, reason string, err error) {

		if limiter.LimitReached() {
			return true, "limit reached", ErrLimiterReached
		}

		return false, "", nil
	}, limiter
}
