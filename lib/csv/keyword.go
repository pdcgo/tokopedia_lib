package csv

import (
	"github.com/pdcgo/tokopedia_lib/lib/helper"
)

func LoadKeyword(fname string) ([]string, error) {
	results, err := helper.FileLoadLineString(fname)
	if err != nil {
		return nil, err
	}
	return results, err
}

func SaveKeyword(fname string, keywords []string) error {
	return helper.FileSaveLineString(fname, keywords)
}
