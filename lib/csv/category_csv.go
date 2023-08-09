package csv

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type StatusGrabCategoryCsv string

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

func LoadCategoryCsv(base *legacy_source.BaseConfig) ([]*CategoryCsv, error) {
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

func SaveCategoryCsv(base *legacy_source.BaseConfig, categories []*CategoryCsv) error {
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
