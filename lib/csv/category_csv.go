package csv

import (
	"encoding/csv"
	"io"
	"os"

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
	Type       string                `csv:"type" json:"type"`
	ParentName string                `csv:"parent_name" json:"parent_name"`
	Name       string                `csv:"name" json:"name"`
	Url        string                `csv:"url" json:"url"`
	Status     StatusGrabCategoryCsv `csv:"status" json:"status"`
}

func NewCategoryCsv(parent, category *model_public.Categories) *CategoryCsv {

	return &CategoryCsv{
		Type:       "category",
		ParentName: parent.Name,
		Name:       category.Name,
		Url:        category.URL,
		Status:     "",
	}
}

func LoadCategoryCsv(base BaseConfig) ([]*CategoryCsv, error) {

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.FieldsPerRecord = -1
		r.LazyQuotes = true
		return r
	})

	fname := base.Path("tokopedia_list_category.csv")
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	results := []*CategoryCsv{}
	err = gocsv.UnmarshalFile(file, &results)
	if err != nil {
		return nil, err
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

func GetCategoryByUrl(categories []*model_public.Categories, url string) <-chan *model_public.Categories {

	category := make(chan *model_public.Categories)
	go func() {
		defer close(category)

	Loop:
		for _, categ := range categories {
			if categ.URL == url {
				category <- categ
				break Loop
			}
			if len(categ.Children) != 0 {
				childsGroup := GetCategoryByUrl(categ.Children, url)
				for child := range childsGroup {
					category <- child
					break Loop
				}
			}
		}
	}()
	return category
}
