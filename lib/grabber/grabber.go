package grabber

import (
	"net/url"
	"strings"
	"sync"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/grab_handler"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type Grabber interface {
	Run() error
}

type BaseGrabber struct {
	Api          *api_public.TokopediaApiPublic
	Base         *legacy_source.BaseConfig
	GrabTasker   *legacy.GrabTasker
	CacheHandler *grab_handler.CacheProductHandler

	wg         sync.WaitGroup
	limitGuard chan int
}

func NewBaseGrabber(
	api *api_public.TokopediaApiPublic,
	base *legacy_source.BaseConfig,
	tasker *legacy.GrabTasker,
	cacheHandler *grab_handler.CacheProductHandler,
) *BaseGrabber {

	grabBasic := legacy.NewGrabBasic(base)

	return &BaseGrabber{
		Api:          api,
		Base:         base,
		CacheHandler: cacheHandler,
		GrabTasker:   tasker,
		limitGuard:   make(chan int, grabBasic.Concurrent),
	}
}

func ParseProductDetailParamsFromUrl(uri string) (*model_public.PdpGetlayoutQueryVar, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	path := u.EscapedPath()
	query := u.Query()

	splitPath := strings.Split(path, "/")
	shopDomain := splitPath[len(splitPath)-2]
	productKey := splitPath[len(splitPath)-1]

	payload := &model_public.PdpGetlayoutQueryVar{
		ShopDomain: shopDomain,
		ProductKey: productKey,
		APIVersion: 1,
		ExtParam:   url.QueryEscape(query.Get("extParam")),
	}
	return payload, nil
}

func GenerateShopProductVar() *model_public.ShopProductVar {
	params := &model_public.ShopProductVar{
		Page:           1,
		PerPage:        100,
		EtalaseID:      "etalase",
		Sort:           1,
		Sid:            "",
		UserDistrictID: "176",
		UserCityID:     "2274",
		UserLat:        "",
		UserLong:       "",
	}
	return params
}

// func (g *Grabber) GetKeywords() ([]string, error) {
// 	path := g.Base.Path(g.GrabTasker.ProductURL)
// 	file, err := os.ReadFile(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	keywords := strings.Split(string(file), "\n")
// 	return keywords, nil
// }

// func (grab *Grabber) AppliedFilterShop(shopId int, shopDomain string) bool {
// 	shopFilter := filter.CreateShopFilter(*grab.Filter, filter.Shop{
// 		Id:     shopId,
// 		Domain: shopDomain,
// 	})

// 	return shopFilter.ApplyFilter()
// }

// func (grab *Grabber) AppliedFilterProduct(prodId int, prodName string, prodUrl string) bool {
// 	productFilter := filter.CreateProductFilter(*grab.Filter, filter.ProductFilterModel{
// 		ProductId:   prodId,
// 		ProductName: prodName,
// 		ProductUrl:  prodUrl,
// 	})
// 	return productFilter.ApplyFilter()
// }

// func (grab *Grabber) GetPublicProductLayout(url string) (*model_public.PdpGetlayoutQueryResp, error) {
// 	prodVar, _ := ParseProductDetailParamsFromUrl(url)
// 	variable := &model_public.PdpGetlayoutQueryVar{
// 		ShopDomain: prodVar.ShopDomain,
// 		ProductKey: prodVar.ProductKey,
// 		APIVersion: 1,
// 	}
// 	product_detail, err := grab.Api.PdpGetlayoutQuery(variable)
// 	if err != nil {
// 		fmt.Printf("error [ produk ] : error  mendapatkan produk [ %s ]\n", prodVar.ProductKey)
// 		return nil, err
// 	}

// 	if product_detail.Data.PdpGetLayout.BasicInfo.Alias == "" {
// 		fmt.Printf("error [ produk ] : produk [ %s ] tidak mempunyai data yang lengkap\n", prodVar.ProductKey)
// 		return nil, errors.New("")
// 	}

// 	return product_detail, nil
// }

// func (grab *Grabber) GetProductDataP2(pdpSession string, shopId string) (*model_public.PdpGetDataP2Resp, error) {

// 	payload := &model_public.PdpGetDataP2Var{
// 		PdpSession: pdpSession,
// 		ProductID:  shopId,
// 	}
// 	product_p2, err := grab.Api.PdpGetDataP2(payload)
// 	if err != nil {
// 		fmt.Printf("error [ produk ] : error mendapatkan produk\n")
// 	}
// 	return product_p2, err
// }

// func (grab *Grabber) GenerateProductSearchParams() *model_public.SearchProductVar {
// 	locs := arrayConverter(grab.Filter.GrabTokopedia.Query.Fcity)
// 	shippings := arrayConverter(grab.Filter.GrabTokopedia.Query.Shipping)

// 	conditions := strings.Split(grab.Filter.GrabTokopedia.Query.Condition, ",")
// 	shopRating := strings.Split(grab.Filter.GrabTokopedia.Query.Rt, ",")

// 	shopTier := []string{}
// 	if grab.Filter.GrabTokopedia.Query.Official {
// 		shopTier = append(shopTier, "2")
// 	}

// 	if grab.Filter.GrabTokopedia.Query.Goldmerchant {
// 		shopTier = append(shopTier, "3")
// 	}

// 	params := model_public.SearchProductVar{
// 		Device:         "desktop",
// 		Sort:           grab.Filter.GrabTokopedia.Query.Ob,
// 		Page:           1,
// 		Rows:           100,
// 		UserDistrictID: "176",
// 		UserCityID:     "2274",
// 		Related:        true,
// 		Scheme:         "https",
// 		SafeSearch:     false,
// 		TopadsBucket:   true,
// 		Source:         "search",
// 		PriceMin:       grab.Filter.GrabTokopedia.Query.Pmin,
// 		PriceMax:       grab.Filter.GrabTokopedia.Query.Pmax,
// 		PreOrder:       grab.Filter.GrabTokopedia.Query.Preorder,
// 		Locations:      url.QueryEscape(strings.Join(locs, ",")),
// 		Rate:           url.QueryEscape(strings.Join(shopRating, "#")),
// 		Condition:      url.QueryEscape(strings.Join(conditions, "#")),
// 		Shipping:       url.QueryEscape(strings.Join(shippings, "#")),
// 		ShopTier:       url.QueryEscape(strings.Join(shopTier, "#")),
// 	}
// 	return &params
// }
