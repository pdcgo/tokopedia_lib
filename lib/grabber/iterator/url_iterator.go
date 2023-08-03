package iterator

import (
	"fmt"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/helper"
)

type UrlHandler func(item string) error

func IterateUrls(base *legacy_source.BaseConfig, tasker *legacy.GrabTasker, handler KeywordHandler) error {
	fname := base.Path(tasker.ProductURL)
	urls, err := helper.FileLoadLineString(fname)
	if err != nil {
		fmt.Println(err, "errors")
		return err
	}

	setUrl := func(urls []string) error {
		return helper.FileSaveLineString(fname, urls)
	}

	for index, keyword := range urls {

		err := handler(keyword)
		if err != nil {
			pdc_common.ReportError(err)
		}

		err = setUrl(urls[index+1:])
		if err != nil {
			pdc_common.ReportError(err)
		}
	}

	return nil
}
