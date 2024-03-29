package grab_handler

import (
	"context"

	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type UrlGrabberResp struct {
	Product   *model_public.PdpGetlayoutQueryResp
	ProductP2 *model_public.PdpGetDataP2Resp
}

type ShopGrabberResp struct {
	Shop    model_public.ShopCoreInfoResp
	Product model_public.PdpGetlayoutQueryResp
}

type ProductListGrabberResp struct {
	Product       model_public.ProductSearch
	ProductDetail model_public.PdpGetlayoutQueryResp
}

type ProductCategoryGrabResp struct {
	ProductCategory model_public.CategoryProduct
	ProductDetail   model_public.PdpGetlayoutQueryResp
}

type CacheProductHandler struct {
	repo *mongorepo.ProductRepo
}

func NewCacheProductHandler(repo *mongorepo.ProductRepo) *CacheProductHandler {
	return &CacheProductHandler{
		repo: repo,
	}
}

func (handler *CacheProductHandler) addItem(cache mongorepo.CacheProduct) error {
	r := handler.repo
	_, err := r.Collection.InsertOne(context.TODO(), cache)
	if err != nil {
		return err
	}
	return nil
}

func (handler *CacheProductHandler) AddItemProductUrl(namespace string, source *UrlGrabberResp) error {
	cache := CreateCacheProductUrl(namespace, source)

	return handler.addItem(cache)
}

func (handler *CacheProductHandler) AddItemProductCategory(namespace string, source *ProductCategoryGrabResp) error {
	cache := CreateCacheProductCategory(namespace, source)

	return handler.addItem(cache)
}

func (handler *CacheProductHandler) AddItemProductShop(namespace string, source *ShopGrabberResp) error {
	cache := CreateCacheProductShop(namespace, source)

	return handler.addItem(cache)
}

func (handler *CacheProductHandler) AddItemProductSearch(namespace string, source *ProductListGrabberResp) error {
	cache := CreateCacheProductSearch(namespace, source)
	return handler.addItem(cache)
}
