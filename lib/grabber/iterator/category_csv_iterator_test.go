package iterator_test

import (
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/tokopedia_lib/lib/csv"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/iterator"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestCategoryCsvIterator(t *testing.T) {

	scen := scenario.NewScenario(t)

	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithBaseConfig(func(base *legacy_source.BaseConfig) {

			t.Run("test category csv iterator no file", func(t *testing.T) {

				items := []*csv.CategoryCsv{}
				err := iterator.IterateCategoryCsv(base, func(item *csv.CategoryCsv) error {
					items = append(items, item)
					return nil
				})

				assert.Nil(t, err)
				assert.Equal(t, 0, len(items))
			})

			t.Run("test category csv iterator with file deprecated", func(t *testing.T) {

				// create file category csv
				fname := base.Path("tokopedia_list_category.csv")
				fdata := `type,parent_name,name,url,status
category,Rumah Tangga,Cover Kipas Angin,https://www.tokopedia.com/p/rumah-tangga/dekorasi/cover-kipas-angin,`

				remove := scen.CreateFile([]byte(fdata), fname)
				defer remove()

				items := []*csv.CategoryCsv{}
				err := iterator.IterateCategoryCsv(base, func(item *csv.CategoryCsv) error {
					items = append(items, item)
					return nil
				})

				assert.Nil(t, err)
				assert.Empty(t, len(items))
			})

			t.Run("test category csv iterator with file", func(t *testing.T) {

				// create file category csv
				fname := base.Path("tokopedia_list_category.csv")
				fdata := `type,cat_1,cat_2,cat_3,url,status
category,Rumah Tangga,,,https://www.tokopedia.com/p/rumah-tangga,
category,Rumah Tangga,Dekorasi,,https://www.tokopedia.com/p/rumah-tangga/dekorasi,
category,Rumah Tangga,Dekorasi,Cover Kipas Angin,https://www.tokopedia.com/p/rumah-tangga/dekorasi/cover-kipas-angin,`

				remove := scen.CreateFile([]byte(fdata), fname)
				defer remove()

				items := []*csv.CategoryCsv{}
				err := iterator.IterateCategoryCsv(base, func(item *csv.CategoryCsv) error {
					items = append(items, item)
					return nil
				})

				assert.Nil(t, err)
				assert.Equal(t, 3, len(items))
			})

		})
	})
}
