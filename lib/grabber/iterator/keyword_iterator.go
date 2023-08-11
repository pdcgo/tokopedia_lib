package iterator

import (
	"log"
	"os"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/csv"
	"github.com/pdcgo/tokopedia_lib/lib/helper"
)

type KeywordHandler func(item string) error

func IterateKeywords(fname string, handler KeywordHandler) error {
	keywords, err := helper.FileLoadLineString(fname)
	if err != nil {

		if os.IsNotExist(err) {
			log.Printf("[ warning ] file %s not found", fname)
			return nil
		}

		pdc_common.ReportError(err)
		return err
	}

	setKeyword := func(keywords []string) error {
		return csv.SaveKeyword(fname, keywords)
	}

	for index, keyword := range keywords {

		err := handler(keyword)
		if err != nil {
			pdc_common.ReportError(err)
		}

		err = setKeyword(keywords[index+1:])
		if err != nil {
			pdc_common.ReportError(err)
		}
	}

	return nil
}
