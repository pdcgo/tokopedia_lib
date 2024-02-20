package deleter_product

import (
	"errors"
	"log"

	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/pdcgo/tokopedia_lib/lib/model"
)

func DeleteProducts(sapi *api.TokopediaApi, datas []*model.SellerProductItem) error {

	lendata := len(datas)
	if lendata == 0 {
		return nil
	}

	log.Println(sapi.AuthenticatedData.User.Email, "deleting", lendata, "products")
	deletes := make([]*model.BulkProductEditV3Input, lendata)

	for ind, data := range datas {
		deletes[ind] = &model.BulkProductEditV3Input{
			ProductID: data.ID,
			Status:    model.DeletedStatus,
			Shop: model.BulkProductEditShop{
				ID: data.Shop.ID,
			},
		}
	}

	payload := model.BulkProductEditV3Var{
		Input: deletes,
	}

	hasils, err := sapi.BulkProductEditV3(&payload)
	if err != nil {
		return err
	}

	if len(hasils.Data.BulkProductEditV3) == 0 {
		return errors.New("product not deleted")
	}

	for _, hasil := range hasils.Data.BulkProductEditV3 {
		if !hasil.Result.IsSuccess {
			return errors.New(hasil.Result.Header.Reason)
		}
	}

	return nil
}
