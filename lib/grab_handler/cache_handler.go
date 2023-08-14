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

func (handler *CacheProductHandler) AddItem(
	namespace string,
	layout *model_public.PdpGetlayoutQueryResp,
	pdpDataP2 *model_public.PdpGetDataP2Resp,
) error {

	cache, err := CreateCacheProduct(namespace, layout, pdpDataP2)
	if err != nil {
		return err
	}

	r := handler.repo
	_, err = r.Collection.InsertOne(context.TODO(), cache)
	if err != nil {
		return err
	}
	return nil
}
