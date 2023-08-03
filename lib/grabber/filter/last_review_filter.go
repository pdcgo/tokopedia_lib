package filter

import (
	"strconv"
	"time"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

func CreateLastReviewFilter(api *api_public.TokopediaApiPublic, base *legacy_source.BaseConfig) FilterHandler {
	return func(layout *model_public.PdpGetlayoutQueryResp, pdp *model_public.PdpGetDataP2Resp) (cek bool, reason string, err error) {
		grabConfig := legacy.NewGrabBasic(base)
		if !grabConfig.LastReviewActive {
			return false, "", nil
		}
		variable := model_public.ProductReviewListVar{
			ProductID: layout.Data.PdpGetLayout.BasicInfo.ID,
			Page:      1,
			Limit:     15,
			SortBy:    "create_time desc",
		}
		resp, err := api.ProductReviewList(&variable)
		if err != nil {
			return true, "filter last review", err
		}

		productReviews := resp.Data.ProductrevGetProductReviewList.List
		if len(productReviews) == 0 {
			return true, "filter last review", nil
		}

		lastProductReview, err := strconv.Atoi(productReviews[0].ReviewCreateTime)
		if err != nil {
			return true, "filter last review", err
		}

		t := time.Now()
		filterLastReview := t.AddDate(0, 0, -grabConfig.LastReviewDays)
		if int(filterLastReview.Unix()) > lastProductReview {
			return true, "filter last review", nil
		}

		return false, "", nil
	}
}
