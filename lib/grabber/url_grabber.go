package grabber

// import (
// 	"errors"
// 	"fmt"
// 	"strconv"

// 	"github.com/pdcgo/go_v2_shopeelib/helper"
// 	"github.com/pdcgo/tokopedia_lib/lib/grab_handler"
// 	"github.com/pdcgo/tokopedia_lib/lib/model_public"
// )

// type UrlGrabber struct {
// 	*Grabber
// }

// func NewUrlGrabber(
// 	grabber *Grabber) *UrlGrabber {
// 	return &UrlGrabber{
// 		Grabber: grabber,
// 	}
// }

// func (grab *UrlGrabber) ProcessProduct(product *model_public.PdpGetlayoutQueryResp) error {
// 	prodName := product.Data.PdpGetLayout.BasicInfo.Alias
// 	if grab.GrabTasker.UseFilter {
// 		prodVar, _ := ParseProductDetailParamsFromUrl(product.Data.PdpGetLayout.BasicInfo.URL)
// 		shopId, _ := strconv.Atoi(product.Data.PdpGetLayout.BasicInfo.ShopID)
// 		prodId, _ := strconv.Atoi(product.Data.PdpGetLayout.BasicInfo.ID)
// 		if grab.Grabber.AppliedFilterShop(shopId, prodVar.ShopDomain) {
// 			return errors.New("terkena filter toko")
// 		}

// 		if grab.Grabber.AppliedFilterProduct(prodId, prodName, product.Data.PdpGetLayout.BasicInfo.URL) {
// 			return errors.New("terkena filter produk")
// 		}
// 	}

// 	productP2, err := grab.Grabber.GetProductDataP2(product.Data.PdpGetLayout.PdpSession, product.Data.PdpGetLayout.BasicInfo.ID)
// 	if err != nil {
// 		return err
// 	}
// 	err = grab.Save(&grab_handler.UrlGrabberResp{Product: product, ProductP2: productP2})
// 	fmt.Printf("grab [ url ] : mendapatkan produk [ %s ]\n", prodName)
// 	return err
// }

// func (grab *UrlGrabber) Save(product *grab_handler.UrlGrabberResp) error {
// 	return grab.Grabber.CacheHandler.AddItemProductUrl(grab.GrabTasker.Namespace, product)
// }

// func (grab *UrlGrabber) Run() error {
// 	path := grab.Base.Path(grab.GrabTasker.ProductURL)
// 	urls, err := helper.FileLineStringLoad(path)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	for _, url := range urls {
// 		product, err := grab.Grabber.GetPublicProductLayout(url)
// 		if err != nil {
// 			continue
// 		}
// 		grab.ProcessProduct(product)
// 	}

// 	return nil
// }
