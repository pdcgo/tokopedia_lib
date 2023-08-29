package csv_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/csv"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestCategoryCsv(t *testing.T) {

	scen := scenario.NewScenario(t)

	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {

		t.Run("test load category csv deprecated", func(t *testing.T) {

			fname := scen.Path("tokopedia_list_category.csv")
			fdata := `type,parent_name,name,url,status
category,Rumah Tangga,Cover Kipas Angin,https://www.tokopedia.com/p/rumah-tangga/dekorasi/cover-kipas-angin,`

			remove := scen.CreateFile([]byte(fdata), fname)
			defer remove()

			categories, err := csv.LoadCategoryCsv(scen)
			assert.Equal(t, err, csv.ErrDeprecatedCategoryCsv)
			assert.Empty(t, categories)
		})

		t.Run("test load category csv ok", func(t *testing.T) {

			fname := scen.Path("tokopedia_list_category.csv")
			fdata := `type,cat_1,cat_2,cat_3,url,status
category,Rumah Tangga,,,https://www.tokopedia.com/p/rumah-tangga,
category,Rumah Tangga,Dekorasi,,https://www.tokopedia.com/p/rumah-tangga/dekorasi,
category,Rumah Tangga,Dekorasi,Cover Kipas Angin,https://www.tokopedia.com/p/rumah-tangga/dekorasi/cover-kipas-angin,`

			remove := scen.CreateFile([]byte(fdata), fname)
			defer remove()

			categories, err := csv.LoadCategoryCsv(scen)
			assert.Nil(t, err)
			assert.Equal(t, len(categories), 3)
		})
	})
}
