package dumper

import (
	"sync"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/csv"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type CategoryCsvDumper struct {
	Api  *api_public.TokopediaApiPublic
	base *legacy_source.BaseConfig
	Data []*csv.CategoryCsv
}

func NewCategoryCsvDumper(api *api_public.TokopediaApiPublic, base *legacy_source.BaseConfig) *CategoryCsvDumper {
	return &CategoryCsvDumper{
		Api:  api,
		base: base,
	}
}

var getCategoriesOne sync.Once
var Categories []*model_public.Categories

func (d *CategoryCsvDumper) GetCategories() []*model_public.Categories {
	getCategoriesOne.Do(func() {
		categories, err := d.Api.HeaderMainData()
		if err != nil {
			panic(err)
		}
		Categories = categories.Data.CategoryAllListLite.Categories
	})
	return Categories

}

func (d *CategoryCsvDumper) IterateGroup(categories []*model_public.Categories, groups []*csv.CategoryCsv) <-chan []*csv.CategoryCsv {
	if categories == nil {
		categories = d.GetCategories()
	}

	cGroupChan := make(chan []*csv.CategoryCsv)

	go func() {
		defer close(cGroupChan)

		for _, categ := range categories {
			if groups == nil {
				groups = []*csv.CategoryCsv{}
			}

			parentName := ""
			if len(groups) > 0 {
				parentName = groups[0].Name
			}

			groupsClone := append(groups, &csv.CategoryCsv{
				Type:       "category",
				ParentName: parentName,
				Name:       categ.Name,
				Url:        categ.URL,
				Status:     "",
			})

			if len(categ.Children) != 0 {
				childsGroup := d.IterateGroup(categ.Children, groupsClone)
				for child := range childsGroup {
					cGroupChan <- child
				}
			}
			cGroupChan <- groupsClone
		}
	}()
	return cGroupChan
}

func (d *CategoryCsvDumper) DumpCategory() error {
	categories := d.IterateGroup(nil, nil)
	for c := range categories {
		category := &csv.CategoryCsv{
			Type:       "category",
			ParentName: c[len(c)-1].ParentName,
			Name:       c[len(c)-1].Name,
			Url:        c[len(c)-1].Url,
			Status:     "",
		}
		d.Data = append(d.Data, category)
	}
	err := d.Save()
	return err
}

func (d *CategoryCsvDumper) Save() error {
	return csv.SaveCategoryCsv(d.base, d.Data)
}

func (d *CategoryCsvDumper) Load() ([]*csv.CategoryCsv, error) {
	return csv.LoadCategoryCsv(d.base)
}
