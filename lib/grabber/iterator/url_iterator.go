package iterator

import (
	"log"
	"math"
	"os"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/helper"
)

type UrlHandler func(items []string) error

func IterateUrls(fname string, handler UrlHandler) error {
	urls, err := helper.FileLoadLineString(fname)
	if err != nil {

		if os.IsNotExist(err) {
			log.Printf("[ warning ] file %s not found", fname)
			return nil
		}

		pdc_common.ReportError(err)
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
