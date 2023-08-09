package iterator

import (
	"fmt"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/csv"
	"github.com/pdcgo/tokopedia_lib/lib/helper"
)

type KeywordHandler func(item string) error

func IterateKeywords(base *legacy_source.BaseConfig, tasker *legacy.GrabTasker, handler KeywordHandler) error {
	fname := base.Path(tasker.Keyword)
	keywords, err := helper.FileLoadLineString(fname)
	if err != nil {
		fmt.Println(err, "errors")
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
