package grab_handler

import (
	"math/rand"
	"strconv"

	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type ProductParser struct {
	ShopId         int
	ProductId      int
	CatId          int
	ProductSold    int
	Stock          int
	ProductContent model_public.ProductContentComponent
	ProductDetail  model_public.ProductDetailComponent
	ProductMedia   model_public.MediaComponent
	ProductDesc    model_public.ProductDetailContent
	Images         []string
	Category       []int64
	Categories     []mongorepo.ProductCategory
}

func createProductCategories(category model_public.Category) ([]mongorepo.ProductCategory, error) {

	categories := []mongorepo.ProductCategory{}

	for _, cat := range *category.Detail {

		id, err := strconv.ParseInt(cat.ID, 10, 64)
		if err != nil {
			return categories, err
		}

		categ := mongorepo.ProductCategory{
			Catid:       id,
			DisplayName: cat.Name,
		}
		categories = append(categories, categ)
	}

	return categories, nil
}

func CreateCacheProduct(
	namespace string,
	layout *model_public.PdpGetlayoutQueryResp,
	pdp *model_public.PdpGetDataP2Resp,
) (mongorepo.CacheProduct, error) {

	pdpLayout := layout.Data.PdpGetLayout

	catname := layout.Data.PdpGetLayout.BasicInfo.Category.Name
	shopLocation := pdp.Data.PdpGetData.ShopInfo.Location
	url := layout.Data.PdpGetLayout.BasicInfo.URL

	shopid, err := strconv.ParseInt(pdpLayout.BasicInfo.ShopID, 10, 64)
	if err != nil {
		return mongorepo.CacheProduct{}, err
	}
	shop := mongorepo.ProductShop{
		Shopid:   shopid,
		Location: shopLocation,
	}

	productid, err := strconv.ParseInt(pdpLayout.BasicInfo.ID, 10, 64)
	if err != nil {
		return mongorepo.CacheProduct{}, err
	}

	catid := int(pdpLayout.BasicInfo.Category.ID)
	name, err := pdpLayout.GetProductName()
	if err != nil {
		return mongorepo.CacheProduct{}, err
	}

	price, err := pdpLayout.GetPrice()
	if err != nil {
		return mongorepo.CacheProduct{}, err
	}

	priceBeforeDiscount, err := pdpLayout.GetPriceBeforeDiscount()
	if err != nil {
		return mongorepo.CacheProduct{}, err
	}

	images, err := pdpLayout.GetImages()
	if err != nil {
		return mongorepo.CacheProduct{}, err
	}

	stock, err := pdpLayout.GetStock()
	if err != nil {
		return mongorepo.CacheProduct{}, err
	}

	sold := pdpLayout.BasicInfo.TxStats.CountSold
	if err != nil {
		return mongorepo.CacheProduct{}, err
	}

	desc, err := pdpLayout.GetDescription()
	if err != nil {
		return mongorepo.CacheProduct{}, err
	}

	categories, err := createProductCategories(layout.Data.PdpGetLayout.BasicInfo.Category)
	if err != nil {
		return mongorepo.CacheProduct{}, err
	}

	catIds := []int64{}
	for _, c := range categories {
		catIds = append(catIds, c.Catid)
	}

	res := mongorepo.CacheProduct{
		Shop:                shop,
		Marketplace:         mongorepo.MP_TOKOPEDIA,
		Id:                  productid,
		Productid:           productid,
		Namespace:           namespace,
		Rnd:                 rand.Float64(),
		Name:                name,
		Price:               int64(price),
		PriceAfterDiscount:  int64(price),
		PriceBeforeDiscount: int64(priceBeforeDiscount),
		Image:               images[0],
		Images:              images,
		Sold:                int32(sold),
		ShopLocation:        shopLocation,
		Catid:               int32(catid),
		CategoryId:          int64(catid),
		CatName:             catname,
		Stock:               int32(stock),
		Desc:                desc,
		Url:                 url,
		Category:            catIds,
		Categories:          categories,
	}

	return res, nil
}
