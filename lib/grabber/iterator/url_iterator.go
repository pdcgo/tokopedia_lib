package iterator

import (
	"fmt"
	"math"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/helper"
)

type UrlHandler func(items []string) error

func IterateUrls(base *legacy_source.BaseConfig, tasker *legacy.GrabTasker, handler UrlHandler) error {
	fname := base.Path(tasker.ProductURL)
	urls, err := helper.FileLoadLineString(fname)
	if err != nil {
		fmt.Println(err, "errors")
		return err
	}

	setUrl := func(urls []string) error {
		return helper.FileSaveLineString(fname, urls)
	}

	prodLength := len(urls)
	maxArray := math.Ceil(float64(prodLength) / 10)
	for i := 1; i <= int(maxArray); i++ {
		startIndex := i*10 - 10
		endIndex := i * 10
		if endIndex > prodLength {
			endIndex = prodLength
		}

		err := handler(urls[startIndex:endIndex])
		if err != nil {
			return err
		}

		var saveUrls []string
		if endIndex == prodLength {
			saveUrls = urls[endIndex:]
		} else {
			saveUrls = urls[endIndex+1:]
		}
		err = setUrl(saveUrls)
		if err != nil {
			pdc_common.ReportError(err)
		}
	}

	return nil
}
