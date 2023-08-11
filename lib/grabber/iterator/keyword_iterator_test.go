package iterator_test

import (
	"strings"
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/iterator"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestKeywordIterator(t *testing.T) {

	scen := scenario.NewScenario(t)

	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithBaseConfig(func(base *legacy_source.BaseConfig) {

			t.Run("test keyword iterator no file", func(t *testing.T) {

				keywords := []string{}
				err := iterator.IterateKeywords("nofile.txt", func(item string) error {
					keywords = append(keywords, item)
					return nil
				})

				assert.Nil(t, err)
				assert.Equal(t, 0, len(keywords))
			})

			t.Run("test keyword iterator with file", func(t *testing.T) {

				// create file keyword
				fname := base.Path("keywords.txt")
				fdata := []string{
					"kuningan",
					"ayam bakar",
					"cinderamata",
					"cidera mata",
				}
				fbdata := []byte(strings.Join(fdata, "\n"))

				scen.CreateFile(fbdata, fname)

				keywords := []string{}
				err := iterator.IterateKeywords(fname, func(item string) error {
					keywords = append(keywords, item)
					return nil
				})

				assert.Nil(t, err)
				assert.Equal(t, fdata, keywords)
			})
		})
	})
}
