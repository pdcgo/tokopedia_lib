package iterator

import (
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
)

type CategoryHandler func(catId string) error

func IterateCategory(
	tasker *legacy.GrabTasker,
	handler CategoryHandler,
) error {
	categoryIds := tasker.TokpedCateg

	for _, categoryId := range categoryIds {
		err := handler(categoryId)
		if err != nil {
			pdc_common.ReportError(err)
		}

	}
	return nil
}
