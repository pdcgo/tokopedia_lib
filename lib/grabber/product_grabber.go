package grabber

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"sync"

	"github.com/pdcgo/go_v2_shopeelib/helper"
	"github.com/pdcgo/tokopedia_lib/lib/filter"
	"github.com/pdcgo/tokopedia_lib/lib/grab_handler"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type ProductGrabber struct {
	*Grabber
}

func arrayConverter(datas []interface{}) []string {
	results := make([]string, len(datas))
	for _, data := range datas {
		switch value := data.(type) {
		case int:
			results = append(results, fmt.Sprint(value))
		case string:
			results = append(results, value)
		}
	}
	return results
}

func (grab *ProductGrabber) getProducts(params *model_public.SearchProductVar) ([]*model_public.ProductSearch, error) {

	rawParams, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	stringParams := string(rawParams)
	replacer := strings.NewReplacer(`"`, "", ":", "=", ",", "&", " ", "", "[", "", "]", "")
	validParams := replacer.Replace(stringParams)

	variable := &model_public.ParamsVar{
		Params: validParams[1 : len(validParams)-1],
	}
	resp, err := grab.Api.SearchProductQueryV4(variable)
	if err != nil {
		return nil, err
	}
	products := resp.Data.AceSearchProductV4.Data.Products

	return products, nil
}

func (grab *ProductGrabber) AppliedFilterShop(shopId int, shopDomain string) bool {
	shopFilter := filter.CreateShopFilter(*grab.Filter, filter.Shop{
		Id:     shopId,
		Domain: shopDomain,
	})

	return shopFilter.ApplyFilter()
}

func (grab *ProductGrabber) AppliedFilterProduct(prodId int, prodName string, prodUrl string) bool {
	productFilter := filter.CreateProductFilter(*grab.Filter, filter.ProductFilterModel{
		ProductId:   prodId,
		ProductName: prodName,
		ProductUrl:  prodUrl,
	})
	return productFilter.ApplyFilter()
}

func (grab *ProductGrabber) ProcessProduct(product *model_public.ProductSearch) error {
	prodVar, _ := ParseProductDetailParamsFromUrl(product.URL)

	if grab.AppliedFilterShop(product.Shop.ShopID, prodVar.ShopDomain) {
		return errors.New("terkena filter toko")
	}

	if grab.AppliedFilterProduct(int(product.ID), product.Name, product.URL) {
		return errors.New("terkena filter toko")
	}

	product_detail, err := grab.GetPublicProductLayout(product.URL)
	if err != nil {
		fmt.Printf("error [ produk ] : error  mendapatkan produk [ %s ]\n", product.Name)
		return err
	}
	grab.Save("", &grab_handler.ProductListGrabberResp{Product: product, ProductDetail: product_detail})
	return nil
}

// - single responsebility x
// - filter
// Stepes:
//  get Products
// parse params to string
// request
// get data
//  filter shop
// create shop filter
// apply filter
// check shop tier
// check last login
//  filter product
// create product filter
// apply filter
// check stock
// check persentase
// check penjualan
// iterasi next page

// get produk
//

// - unit test x
// - flexible v

func (grab *ProductGrabber) PageIterate(params *model_public.SearchProductVar, handle func(*model_public.ProductSearch) error) error {
	var errResp error
	fmt.Printf("grab [ keyword ] : memulai grab keyword [ %s ]\n", params.Query)

Parent:
	for {
		products, err := grab.getProducts(params)
		if err != nil {
			errResp = err
			break Parent
		}

		if len(products) == 0 {
			fmt.Printf("finish [ produk ] : halaman ini tidak mempunyai produk\n")
			break Parent
		}

		for _, product := range products {
			handle(product)
		}

		params.Page += 1
		params.Start = params.Page * params.Rows
	}
	return errResp
}

func (grab *ProductGrabber) IterateProductPages(params *model_public.SearchProductVar) (<-chan *model_public.ProductSearch, *helper.ChannelError) {
	res := make(chan *model_public.ProductSearch)
	errChan := helper.NewChannelError()

	go func() {
		defer close(res)
		defer errChan.Raise()

	Parent:
		for {
			products, err := grab.getProducts(params)
			if err != nil {
				fmt.Println(err)
				break Parent
			}
			for _, product := range products {
				res <- product
			}
			params.Page += 1
			params.Start = params.Page * params.Rows
		}
	}()
	return res, errChan
}

func (grab *ProductGrabber) Save(namespace string, product *grab_handler.ProductListGrabberResp) {
	grab.CacheHandler.AddItemProductSearch(namespace, product)
}

// Product List
type ProductListGrabber struct {
	ProductGrabber
	Keywords []string
	wg       *sync.WaitGroup
}

func NewProductListGrabber(
	grabber *Grabber,
	keywords []string) *ProductListGrabber {
	return &ProductListGrabber{
		ProductGrabber: ProductGrabber{
			Grabber: grabber,
		},
		Keywords: keywords,
		wg:       &sync.WaitGroup{},
	}
}

func (grab *ProductListGrabber) Run() {

Keywords:
	for _, keyword := range grab.Keywords {
		fmt.Printf("grab [ keyword ] : memulai grab keyword [ %s ]\n", keyword)
		params := grab.GenerateProductSearchParams()
		params.Query = url.QueryEscape(keyword)

		limit := int32(grab.Filter.GrabBasic.LimitGrab)
		limiter := helper.NewLimiter(limit)
		counter := helper.Counter{}

		products, errChan := grab.IterateProductPages(params)
		for product := range products {
			if limiter.LimitReached() {
				continue Keywords
			}

			err := grab.ProcessProduct(product)
			if err != nil {
				continue
			}
			fmt.Printf("grab [ keyword ] : mendapatkan produk [ %s ] [ %d ]\n", product.Name, counter.Total)
			limiter.Add()
			counter.Add()
		}

		err := errChan.GetError()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

type CategoryGrabber struct {
	ProductGrabber
	CatId int
}

func NewCategoryGrabber(grabber *Grabber, catId int) *CategoryGrabber {
	return &CategoryGrabber{
		ProductGrabber: ProductGrabber{
			Grabber: grabber,
		},
		CatId: catId,
	}
}

func (grab *CategoryGrabber) Run() {
	params := grab.GenerateProductSearchParams()
	params.CategoryId = grab.CatId

	limit := int32(grab.Filter.GrabBasic.LimitGrab)
	limiter := helper.NewLimiter(limit)

	products, errChan := grab.IterateProductPages(params)
	for product := range products {
		if limiter.LimitReached() {
			fmt.Println("filter [ produk ] : telah mencapai batas grab")
			return
		}

		err := grab.ProcessProduct(product)
		if err != nil {
			continue
		}
		fmt.Printf("grab [ kategori ] : mendapatkan produk [ %s ]\n", product.Name)
		limiter.Add()
	}

	err := errChan.GetError()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
