package iterator

import (
	"log"
	"os"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/tokopedia_lib/lib/csv"
)

type CategoryCsvHandler func(item *csv.CategoryCsv) error

func IterateCategoryCsv(base *legacy_source.BaseConfig, handler CategoryCsvHandler) error {
	categories, err := csv.LoadCategoryCsv(base)
	if err != nil {

		if os.IsNotExist(err) {
			log.Println("[ warning ] file tokopedia_list_category.csv tidak ditemukan")
			return nil
		}

		pdc_common.ReportError(err)
		return err
	}
	setGrabbed := func(category *csv.CategoryCsv) error {
		category.Status = csv.STATUS_GRAB_CATEGORY_GRABBED
		err := csv.SaveCategoryCsv(base, categories)
		return err
	}

	for _, category := range categories {
		if category.Status == csv.STATUS_GRAB_CATEGORY_GRABBED {
			continue
		}

		err := handler(category)
		if err != nil {
			pdc_common.ReportError(err)
		}

		err = setGrabbed(category)
		if err != nil {
			pdc_common.ReportError(err)
		}
	}

	return nil
}
