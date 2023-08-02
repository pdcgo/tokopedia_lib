package grabber

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/pdcgo/go_v2_shopeelib/helper"
	"github.com/pdcgo/tokopedia_lib/lib/grab_handler"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type ShopGrabber struct {
	*Grabber
}

func (grab *ShopGrabber) getShopCoreInfo(domain string) (*model_public.ShopCoreInfoResp, error) {
	variable, err := GenerateShopCoreInfoParamsFormUrl(domain)
	if err != nil {
		return nil, err
	}
	fmt.Printf("grab [ shop ] : memulai grab shop [ %s ]\n", variable.Domain)
	shopCoreInfo, err := grab.Api.ShopCoreInfo(variable)
	if err != nil {
		return nil, err
	}
	return shopCoreInfo, nil
}

func (grab *ShopGrabber) IterateProductPages(params *model_public.ShopProductVar) (<-chan *model_public.ShopProductData, *helper.ChannelError) {
	res := make(chan *model_public.ShopProductData)
	errChan := helper.NewChannelError()
	go func() {
		defer close(res)
		defer errChan.Raise()

	Parent:
		for {
			resp, err := grab.Api.ShopProducts(params)
			if err != nil {
				errChan.SetError(err)
				break Parent
			}

			products := resp.Data.GetShopProduct.Data
			if len(products) == 0 {
				errChan.SetError(errors.New("halaman kosong"))
				break Parent
			}

			for _, product := range products {
				res <- &product
			}
			if resp.Data.GetShopProduct.Links.Next == "" {
				break Parent
			}
			params.Page += 1
		}
	}()
	return res, errChan
}

func (grab *ShopGrabber) Save(product *grab_handler.ShopGrabberResp) error {
	return grab.CacheHandler.AddItemProductShop(grab.GrabTasker.Namespace, product)
}

type ShopListGrabber struct {
	*ShopGrabber
}

func NewShopListGrabber(
	grabber *Grabber) *ShopListGrabber {

	return &ShopListGrabber{
		ShopGrabber: &ShopGrabber{
			Grabber: grabber,
		},
	}
}

func (grab *ShopListGrabber) ProcessProduct(shopCore *model_public.ShopCoreInfoResp, product *model_public.ShopProductData) error {
	prodId, _ := strconv.Atoi(product.ProductID)
	if grab.GrabTasker.UseFilter {
		if grab.AppliedFilterProduct(prodId, product.Name, product.ProductURL) {
			return errors.New("terkena filter produk")
		}
	}

	pubProduct, err := grab.GetPublicProductLayout(product.ProductURL)
	if err != nil {
		return err
	}

	err = grab.Save(&grab_handler.ShopGrabberResp{
		Shop:    shopCore,
		Product: pubProduct,
	})
	if err != nil {
		return err
	}

	fmt.Printf("grab [ shop ] : mendapatkan produk [ %s ]\n", product.Name)
	return nil
}

func (grab *ShopListGrabber) Run() error {
	path := grab.Base.Path(grab.GrabTasker.TokoUsername)
	shops, err := helper.FileLineStringLoad(path)
	if err != nil {
		return err
	}

Shop:
	for _, shopUrl := range shops {

		limit := int32(grab.Grabber.Filter.GrabBasic.LimitGrab)
		limiter := helper.NewLimiter(limit)

		shopCoreInfo, err := grab.getShopCoreInfo(shopUrl)
		if err != nil {
			continue Shop
		}

		shopId, _ := strconv.Atoi(shopCoreInfo.Data.Result[0].ShopCore.ShopID)
		if grab.GrabTasker.UseFilter {
			if grab.AppliedFilterShop(shopId, shopCoreInfo.Data.Result[0].ShopCore.Domain) {
				continue Shop
			}
		}

		params := GenerateShopProductVar()
		params.Sid = fmt.Sprintf("%d", shopId)
		products, errChan := grab.IterateProductPages(params)
		for product := range products {
			if limiter.LimitReached() {
				fmt.Println("filter [ produk ] : telah mencapai batas grab")
				continue Shop
			}

			err := grab.ProcessProduct(shopCoreInfo, product)
			if err != nil {
				continue
			}

			limiter.Add()
		}

		err = errChan.GetError()
		if err != nil {
			fmt.Println(err)
			continue Shop
		}
	}
	return nil
}
