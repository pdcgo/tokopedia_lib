package grabber

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/pdcgo/go_v2_shopeelib/helper"
	"github.com/pdcgo/tokopedia_lib/lib/grab_handler"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type UrlGrabber struct {
	Grabber *Grabber
	Urls    []string
}

func NewUrlGrabber(
	grabber *Grabber,
	urls []string) *UrlGrabber {

	return &UrlGrabber{
		Grabber: grabber,
		Urls:    urls,
	}
}

func (grab *UrlGrabber) ProcessProduct(product *model_public.PdpGetlayoutQueryResp) error {
	prodVar, _ := ParseProductDetailParamsFromUrl(product.Data.PdpGetLayout.BasicInfo.URL)
	shopId, _ := strconv.Atoi(product.Data.PdpGetLayout.BasicInfo.ShopID)
	prodId, _ := strconv.Atoi(product.Data.PdpGetLayout.BasicInfo.ID)
	prodName := product.Data.PdpGetLayout.BasicInfo.Alias

	if grab.Grabber.AppliedFilterShop(shopId, prodVar.ShopDomain) {
		return errors.New("terkena filter toko")
	}

	if grab.Grabber.AppliedFilterProduct(prodId, prodName, product.Data.PdpGetLayout.BasicInfo.URL) {
		return errors.New("terkena filter produk")
	}

	productP2, err := grab.Grabber.GetProductDataP2(product.Data.PdpGetLayout.PdpSession, product.Data.PdpGetLayout.BasicInfo.ID)
	if err != nil {
		return err
	}
	grab.Save("", &grab_handler.UrlGrabberResp{Product: product, ProductP2: productP2})
	fmt.Printf("grab [ url ] : mendapatkan produk [ %s ]\n", prodName)
	return nil
}

func (grab *UrlGrabber) Save(namespace string, product *grab_handler.UrlGrabberResp) error {
	return grab.Grabber.CacheHandler.AddItemProductUrl("", product)
}

func (grab *UrlGrabber) Run() {

	limit := int32(grab.Grabber.Filter.GrabBasic.LimitGrab)
	limiter := helper.NewLimiter(limit)

	for _, url := range grab.Urls {
		if limiter.LimitReached() {
			continue
		}

		product, err := grab.Grabber.GetPublicProductLayout(url)
		if err != nil {
			continue
		}
		grab.ProcessProduct(product)
		limiter.Add()
	}
}
