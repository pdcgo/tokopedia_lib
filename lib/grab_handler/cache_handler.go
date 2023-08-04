package grab_handler

import (
	"context"

	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type CacheProductHandler struct {
	repo *mongorepo.ProductRepo
}

func NewCacheProductHandler(repo *mongorepo.ProductRepo) *CacheProductHandler {
	return &CacheProductHandler{
		repo: repo,
	}
}

func (handler *CacheProductHandler) addItem(cache mongorepo.CacheProduct) error {
	if cache.Namespace == "" {
		cache.Namespace = handler.repo.Collection.Name()
	}
	r := handler.repo
	_, err := r.Collection.InsertOne(context.TODO(), cache)
	if err != nil {
		return err
	}
	return nil
}

func (h *CacheProductHandler) AddProductItem(
	namespace string,
	layout *model_public.PdpGetlayoutQueryResp,
	pdpDataP2 *model_public.PdpGetDataP2Resp,
) error {
	cache := createCacheProduct(layout)
	cache.Namespace = namespace
	cache.Shop.Location = pdpDataP2.Data.PdpGetData.ShopInfo.Location
	cache.ShopLocation = pdpDataP2.Data.PdpGetData.ShopInfo.Location
	return h.addItem(cache)
}
