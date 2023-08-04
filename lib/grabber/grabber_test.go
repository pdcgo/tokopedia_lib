package grabber_test

import (
	"context"
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	"github.com/pdcgo/go_v2_shopeelib/lib/upload_config"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/grab_handler"
	"github.com/pdcgo/tokopedia_lib/lib/grabber"
	"github.com/zeebo/assert"
)

func TestGrabAndInsertDB(t *testing.T) {
	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	baseConfig := legacy_source.BaseConfig{
		BaseData: "../..",
	}
	botConfig := upload_config.NewBotConfig(&baseConfig)
	db := botConfig.CreateDB(ctx)

	productRepo := mongorepo.NewProductRepo(ctx, db)

	cacheHandler := grab_handler.NewCacheProductHandler(productRepo)
	// Grab Category
	t.Run("test grab by category", func(t *testing.T) {
		category := grabber.Category{
			1759, "Fashion Pria", "https://www.tokopedia.com/p/fashion-pria",
		}
		grab, err := grabber.CreateProductCategoryGrabber(api, &baseConfig, productRepo, category)
		assert.Nil(t, err)

		resp := make(chan grab_handler.ProductCategoryGrabResp)
		go grab.Run(resp)
		for res := range resp {
			grab.Save(productRepo.Collection.Name(), &res)
		}
	})

	// Grab Urls
	t.Run("test grab by Urls", func(t *testing.T) {
		urls := []string{
			"https://www.tokopedia.com/rurushopping/orico-mouse-pad-cork-small-200x300mm-two-sides-merah-muda?extParam=ivf%3Dfalse%26src%3Dsearch",
			"https://www.tokopedia.com/acomeindone/acome-keyboard-mouse-wireless-portable-1600dpi-silikon-akm2000-keyboard-only?extParam=cmp%3D1%26ivf%3Dfalse&src=topads",
			"https://www.tokopedia.com/gojeteindonesia/jete-x-mouse-gaming-msx1-rgb-wired-6-programmable-buttons-original?extParam=ivf%3Dfalse%26src%3Dsearch%26whid%3D3474367",
			"https://www.tokopedia.com/razer/razer-deathadder-essential-white-essential-gaming-mouse?extParam=ivf%3Dtrue%26src%3Dsearch%26whid%3D1778422",
			"https://www.tokopedia.com/erigo/kaos-erigo-t-shirt-oversize-amery-cotton-combed-light-olive-xl",
			"https://www.tokopedia.com/lawakboy/kaos-pria-distro-anime-one-piece-luffy-luffy-monkey-m?src=topads",
			"https://www.tokopedia.com/tokokaosdistropremium/t-shirt-baju-kaos-distro-anime-luffy-one-piece-original-jumbo-big-size-army-m",
			"https://www.tokopedia.com/tokokaosdistropremium/baju-kaos-t-shirt-atasan-laki-laki-dewasa-distro-premium-elmo-keren-putih-m",
			"https://www.tokopedia.com/coolnerds/kaos-lengan-pendek-minions-kaos-karakter-kartun-l?src=topads",
			"https://www.tokopedia.com/wellborn/wellborn-fantasy-black-t-shirt-s",
			"https://www.tokopedia.com/kizaru/kaos-pria-kizaru-t-shirt-anime-one-piece-luffy-eat-well-s?src=topads",
			"https://www.tokopedia.com/dmchousemerchandise/kaos-baju-anime-monkey-d-luffy-2-one-piece-lengan-panjang-navy-s?src=topads",
			"https://www.tokopedia.com/clothingpedia-printing/kaos-despicable-me-minions-stuart-kevin-bob-minion-face-cute-t-shirt-stuart-anak-xs",
		}
		grab, err := grabber.CreateUrlGrabber(urls)
		assert.Nil(t, err)

		products, err := grab.Run()
		assert.Nil(t, err)

		for _, product := range products {
			cacheHandler.AddItemProductUrl(productRepo.Collection.Name(), product)
		}
	})

	// Grab Shop
	t.Run("test grab shop", func(t *testing.T) {
		shops := []string{
			"logitech-g",
			"https://www.tokopedia.com/jbl-official",
			"https://www.tokopedia.com/ellipses",
			"https://www.tokopedia.com/erigo",
			"https://www.tokopedia.com/wellborn",
			"https://www.tokopedia.com/clothingpedia-printing",
			"https://www.tokopedia.com/coolnerds",
			"https://www.tokopedia.com/tokokaosdistropremium",
			"https://www.tokopedia.com/lawakboy",
			"https://www.tokopedia.com/dmchousemerchandise",
			"https://www.tokopedia.com/kizaru",
		}
		grab, err := grabber.CreateShopListGrabber(shops, "")
		assert.Nil(t, err)

		products := make(chan grab_handler.ShopGrabberResp)
		go grab.Run(products)
		for product := range products {
			cacheHandler.AddItemProductShop(productRepo.Collection.Name(), &product)
		}
	})

	// Grab Keyword
	t.Run("test grab search", func(t *testing.T) {
		keywords := []string{
			"mousepad",
			"keyboard gaming",
		}
		grab, err := grabber.NewProductListGrabber(keywords)
		assert.Nil(t, err)

		products := make(chan grab_handler.ProductListGrabberResp)
		go grab.Run(products)
		for product := range products {
			cacheHandler.AddItemProductSearch(productRepo.Collection.Name(), &product)
		}
	})
}
