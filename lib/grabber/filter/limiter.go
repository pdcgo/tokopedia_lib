package filter

import (
	"errors"
	"sync/atomic"

	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type AddCount func() (limitReached bool)

var ErrLimiterReached = errors.New("limit was reached")

func CreateLimiter(grabBasic *legacy.GrabBasic) (FilterHandler, AddCount) {

	count := int32(grabBasic.LimitGrab)

	addCount := func() bool {
		atomic.AddInt32(&count, -1)
		c := atomic.LoadInt32(&count)

		return c <= 0
	}

	return func(layout *model_public.PdpGetlayoutQueryResp, pdp *model_public.PdpGetDataP2Resp) (cek bool, reason string, err error) {

		if count <= 0 {
			return true, "limit reached", ErrLimiterReached
		}

		return false, "", nil
	}, addCount
}
