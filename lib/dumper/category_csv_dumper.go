package dumper

import (
	"errors"
	"os"
	"sync"

	"github.com/gocarina/gocsv"
	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type Category struct {
	Type       string `csv:"type" json:"type"`
	ParentName string `csv:"parent_name" json:"parent_name"`
	Name       string `csv:"name" json:"name"`
	Url        string `csv:"url" json:"url"`
	Status     string `csv:"status" json:"status"`
}

type CategoryCsvDumper struct {
	Api      *api_public.TokopediaApiPublic
	base     *legacy_source.BaseConfig
	Pathfile string
	Data     []*Category
}

func NewCategoryCsvDumper(api *api_public.TokopediaApiPublic, base *legacy_source.BaseConfig, pathfile string) *CategoryCsvDumper {
	return &CategoryCsvDumper{
		Api:      api,
		base:     base,
		Pathfile: pathfile,
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

func (d *CategoryCsvDumper) IterateGroup(categories []*model_public.Categories, groups []*Category) <-chan []*Category {
	if categories == nil {
		categories = d.GetCategories()
	}

	cGroupChan := make(chan []*Category)

	go func() {
		defer close(cGroupChan)

		for _, categ := range categories {
			if groups == nil {
				groups = []*Category{}
			}

			parentName := ""
			if len(groups) > 0 {
				parentName = groups[0].Name
			}

			groupsClone := append(groups, &Category{
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

func (d *CategoryCsvDumper) DumpCategory() {
	categories := d.IterateGroup(nil, nil)
	for c := range categories {
		category := &Category{
			Type:       "category",
			ParentName: c[len(c)-1].ParentName,
			Name:       c[len(c)-1].Name,
			Url:        c[len(c)-1].Url,
			Status:     "",
		}
		d.Data = append(d.Data, category)
	}
	d.Save()
}

func (d *CategoryCsvDumper) Save() error {
	path := d.base.Path(d.Pathfile)
	_, err := os.Stat(path)
	if errors.Is(os.ErrNotExist, err) {
		os.Create(path)
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	err = gocsv.MarshalFile(d.Data, file)
	if err != nil {
		return err
	}
	return nil
}

func (d *CategoryCsvDumper) Load() ([]*Category, error) {
	path := d.base.Path(d.Pathfile)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = gocsv.UnmarshalFile(file, &d.Data)
	if err != nil {
		return nil, err
	}
	return d.Data, nil
}

func GetFullCategory(categs []*Category) []string {
	name := []string{}
	for _, c := range categs {
		name = append(name, c.Name)
	}
	return name
}

func (d *CategoryCsvDumper) GetCategoryByUrl(categories []*model_public.Categories, url string) <-chan *model_public.Categories {
	if categories == nil {
		categories = d.GetCategories()
	}

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
				childsGroup := d.GetCategoryByUrl(categ.Children, url)
				for child := range childsGroup {
					category <- child
					break Loop
				}
			}
		}
	}()
	return category
}
