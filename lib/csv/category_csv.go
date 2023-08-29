package csv

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type StatusGrabCategoryCsv string
type BaseConfig interface {
	Path(...string) string
}

const (
	STATUS_GRAB_CATEGORY_GRABBED   StatusGrabCategoryCsv = "grabbed"
	STATUS_GRAB_CATEGORY_UNGRABBED StatusGrabCategoryCsv = ""
)

type CategoryCsv struct {
	Type   string                `csv:"type" json:"type"`
	Cat1   string                `csv:"cat_1" json:"cat_1"`
	Cat2   string                `csv:"cat_2" json:"cat_2"`
	Cat3   string                `csv:"cat_3" json:"cat_3"`
	Url    string                `csv:"url" json:"url"`
	Status StatusGrabCategoryCsv `csv:"status" json:"status"`
}

func (c *CategoryCsv) GetName() string {
	if c.Cat3 != "" {
		return c.Cat3
	}

	if c.Cat2 != "" {
		return c.Cat2
	}

	return c.Cat1
}

var ErrNoCategories = errors.New("no categories")

func NewCategoryCsv(categories []*model_public.Categories) (*CategoryCsv, error) {

	category := &CategoryCsv{
		Type:   "category",
		Status: "",
	}

	switch catlen := len(categories); catlen {

	case 1:
		category.Cat1 = categories[0].Name
		category.Url = categories[0].URL
		return category, nil

	case 2:
		category.Cat1 = categories[0].Name
		category.Cat2 = categories[1].Name
		category.Url = categories[1].URL
		return category, nil

	case 3:
		category.Cat1 = categories[0].Name
		category.Cat2 = categories[1].Name
		category.Cat3 = categories[2].Name
		category.Url = categories[2].URL
		return category, nil

	default:
		return category, ErrNoCategories
	}
}

var ErrDeprecatedCategoryCsv = errors.New("category csv kadaluarsa silahkan generate ulang")

func LoadCategoryCsv(base BaseConfig) ([]*CategoryCsv, error) {

	results := []*CategoryCsv{}
	fname := base.Path("tokopedia_list_category.csv")

	file, err := os.ReadFile(fname)
	if err != nil {
		return results, err
	}

	header := strings.Split(string(file), "\n")[0]
	if !strings.Contains(header, "cat_1") {
		return results, ErrDeprecatedCategoryCsv
	}

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.FieldsPerRecord = -1
		r.LazyQuotes = true
		return r
	})
	err = gocsv.UnmarshalBytes(file, &results)
	if err != nil {
		return results, err
	}
	return results, nil

}

func SaveCategoryCsv(base BaseConfig, categories []*CategoryCsv) error {
	fname := base.Path("tokopedia_list_category.csv")

	file, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	err = gocsv.MarshalFile(categories, file)
	if err != nil {
		return err
	}
	return nil
}
