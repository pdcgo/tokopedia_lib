package grabber_test

// import (
// 	"context"
// 	"testing"

// 	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
// 	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
// 	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
// 	"github.com/pdcgo/tokopedia_lib/lib/api_public"
// 	"github.com/pdcgo/tokopedia_lib/lib/filter"
// 	"github.com/pdcgo/tokopedia_lib/lib/grabber"
// 	"github.com/pdcgo/tokopedia_lib/scenario"
// 	"github.com/stretchr/testify/assert"
// )

// func TestProductCategoryGrabber(t *testing.T) {

// 	api, err := api_public.NewTokopediaApiPublic()
// 	assert.Nil(t, err)
// 	ctx := context.Background()
// 	ctx, cancel := context.WithCancel(ctx)
// 	defer cancel()

// 	baseConfig := legacy_source.BaseConfig{
// 		BaseData: "../..",
// 	}
// 	database := scenario.GetMongoDatabase(t)

// 	productRepo := mongorepo.NewProductRepo(ctx, database)
// 	filter := &filter.BaseFilter{
// 		GrabBasic:     legacy.NewGrabBasic(&baseConfig),
// 		GrabTokopedia: legacy.NewGrabTokopedia(&baseConfig),
// 	}
// 	filter.GrabBasic.LimitGrab = 1

// 	t.Run("test grab by first level category", func(t *testing.T) {
// 		category := grabber.Category{
// 			1759, "Fashion Pria", "https://www.tokopedia.com/p/fashion-pria",
// 		}
// 		grab := grabber.CreateProductCategoryGrabber(api, filter, productRepo, category)
// 		assert.Nil(t, err)

// 		grab.Run()
// 	})

// 	t.Run("test grab by last level category", func(t *testing.T) {
// 		category := grabber.Category{
// 			1873, "Jam Tangan Analog Pria", "https://www.tokopedia.com/p/fashion-pria/jam-tangan-pria/jam-tangan-analog-pria",
// 		}
// 		grab := grabber.CreateProductCategoryGrabber(api, filter, productRepo, category)
// 		assert.Nil(t, err)

// 		grab.Run()
// 	})
// }
