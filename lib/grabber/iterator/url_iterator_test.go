package iterator_test

import (
	"strings"
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/iterator"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestUrlIterator(t *testing.T) {

	scen := scenario.NewScenario(t)

	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithBaseConfig(func(base *legacy_source.BaseConfig) {

			t.Run("test url iterator no file", func(t *testing.T) {

				urls := []string{}
				err := iterator.IterateUrls("nofile.txt", func(items []string) error {
					urls = append(urls, items...)
					return nil
				})

				assert.Nil(t, err)
				assert.Equal(t, 0, len(urls))
			})

			t.Run("test url iterator with file", func(t *testing.T) {

				// create file keyword
				fname := base.Path("urls.txt")
				fdata := []string{}
				for i := 0; i < 15; i++ {
					fdata = append(fdata, "https://www.tokopedia.com/alazkafrozenmeat/jantung-ayam-chicken-hearts-jeroan-jantung-ayam-beku-1-kg-500-gram")
				}
				fbdata := []byte(strings.Join(fdata, "\n"))

				scen.CreateFile(fbdata, fname)

				urls := []string{}
				err := iterator.IterateUrls(fname, func(items []string) error {

					urls = append(urls, items...)
					assert.LessOrEqual(t, len(items), 10)

					return nil
				})

				assert.Nil(t, err)
				assert.Equal(t, len(fdata), len(urls))
			})
		})
	})
}
