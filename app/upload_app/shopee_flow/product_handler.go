package shopee_flow

import (
	"errors"
	"log"
	"strings"

	"github.com/pdcgo/common_conf/common_concept"
	libmongo "github.com/pdcgo/go_v2_shopeelib/lib/mongo"
	shopeeuploader "github.com/pdcgo/go_v2_shopeelib/lib/uploader"
	"github.com/pdcgo/tokopedia_lib/lib/repo"
	"github.com/pdcgo/tokopedia_lib/lib/uploader"
)

func (flow *ShopeeToTopedFlow) createProductHandler(akun *repo.AkunItem, spin shopeeuploader.SpinFunc) uploader.UploadHandler {
	return func(eventcore uploader.EmitFunc, tokpedup *uploader.TokopediaUploader, payload *uploader.PayloadUpload, sub *common_concept.Subscriber) error {

		sub.Cancel()
		product, err := flow.productRepo.Get(libmongo.MP_SHOPEE, akun.Collection, true)
		if err != nil {
			if strings.Contains(err.Error(), "cannot decode") {
				return errors.New(product.Name + ", " + err.Error() + ", silahkan grab baru")
			}
			return err
		}
		log.Println("getting from database", product.Name, product.Shop.Shopid, product.Id)

		source := product.PulicSource
		if source == nil {
			return errors.New(product.Name + " source not found grab cache expired.. silahkan grab baru")
		}
		distance := product.Distance
		if distance == nil {
			return errors.New(product.Name + " distance not found grab cache expired.. silahkan grab baru")
		}

		eventcore(source)
		eventcore(distance)

		return nil
	}
}
