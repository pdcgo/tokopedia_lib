package shopee_flow

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/pdcgo/common_conf/pdc_common"
	shopee_upapp "github.com/pdcgo/go_v2_shopeelib/app/upload_app"
	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/spin"
	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	shopeeuploader "github.com/pdcgo/go_v2_shopeelib/lib/uploader"
	"github.com/pdcgo/tokopedia_lib/app/config"
	"github.com/pdcgo/tokopedia_lib/app/services"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/repo"
	"github.com/pdcgo/tokopedia_lib/lib/uploader"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func CreateConfigConcurencyFromCmd() *shopee_upapp.UploadConcurencyConfig {
	// AccountConcurency: browserC,
	// ProductPerAccount: productC,

	fmt.Print("Tentukan Jumlah Concurent Akun: ")
	var browserC int
	_, err := fmt.Scanln(&browserC)
	if err != nil {
		log.Fatal(err)
	}

	return &shopee_upapp.UploadConcurencyConfig{
		AccountConcurency: browserC,
		ProductPerAccount: 100,
	}
}

type ShopeeToTopedFlow struct {
	configrepo     *config.ConfigRepo
	mapper         *config.ShopeeMapper
	Ctx            context.Context
	CancelCtx      func()
	limitGuard     chan int
	productRepo    *mongorepo.ProductRepo
	ConfigFlow     *shopee_upapp.ConfigUploadFlow
	TopedPublicApi *api_public.TokopediaApiPublic
	AkunIterator   *repo.AkunUploadIterator
	CacheApi       *CacheApiDriver
	etalasemap     *services.EtalaseMapService
}

func NewShopeeToTopedFlow(rootBase string, ctx context.Context, db *mongo.Database, sqlitedb *gorm.DB, concurent *shopee_upapp.UploadConcurencyConfig, publicapi *api_public.TokopediaApiPublic) *ShopeeToTopedFlow {
	cctx, cancel := context.WithCancel(ctx)

	configFlow := shopee_upapp.InitUploadFlowConfig(&legacy_source.BaseConfig{
		BaseData: rootBase,
	}, db, concurent)

	productRepo := mongorepo.NewProductRepo(ctx, db)

	iterator := repo.NewAkunUploadIterator(sqlitedb)
	shopeemapper := config.NewShopeeMapper(sqlitedb)

	configrepo := config.NewConfigRepo(sqlitedb)

	return &ShopeeToTopedFlow{
		configrepo:     configrepo,
		mapper:         shopeemapper,
		limitGuard:     make(chan int, configFlow.UploadConcurencyConfig.AccountConcurency),
		ConfigFlow:     configFlow,
		productRepo:    productRepo,
		TopedPublicApi: publicapi,
		AkunIterator:   iterator,
		Ctx:            cctx,
		CacheApi:       NewCacheApiDriver(),
		CancelCtx:      cancel,
		etalasemap:     services.NewEtalaseMapService(sqlitedb),
	}
}

func (flow *ShopeeToTopedFlow) RunTask() {
	defer func() {
		<-flow.limitGuard
	}()

	akun, updateinc, _, err := flow.AkunIterator.Get()
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {

			count, err := flow.AkunIterator.InProcessCount()
			if err != nil {
				pdc_common.ReportError(err)
				return
			}

			if count == 0 {
				log.Println("akun yang diproses habis...")
				flow.CancelCtx()
				return
			}

			log.Println("semua akun masih diproses")
			time.Sleep(time.Second * 5)
			return

		} else {
			pdc_common.ReportError(err)
			return
		}

	}

	sleep := func() {
		sl := flow.ConfigFlow.GenSleep()
		log.Println(akun.Username, " delay for ", sl)
		time.Sleep(time.Second * time.Duration(sl))
	}

	api, saveApi, err := flow.CacheApi.Get(akun)
	defer saveApi()
	if err != nil {
		pdc_common.ReportError(err)
		return
	}

	ctx, cancel := context.WithTimeout(flow.Ctx, time.Minute*3)
	defer cancel()

	uploaderItem := uploader.NewTokopediaUploader(ctx, api)
	handlers := flow.GenerateHandler(akun, flow.GenerateSpinHandler(akun))
	_, err = uploaderItem.RunUploader(handlers...)

	if err != nil {
		pdc_common.ReportError(err)
		go func() {
			updateinc(0, err)
			sleep()
		}()

		return
	} else {
		go func() {
			log.Println(akun.CountUpload+1, "/", akun.LimitUpload, api.AuthenticatedData.User.Email, "uploaded...")
			sleep()
			updateinc(1, err)
		}()

	}

	// sleeping account

}

func (flow *ShopeeToTopedFlow) GenerateSpinHandler(akun *repo.AkunItem) shopeeuploader.SpinFunc {
	akunlegacy := legacy_source.Akun{
		Name:      akun.Username,
		Hastag:    akun.Hastag,
		Namespace: akun.Collection,
		Markup:    akun.Markup,
		Polatitle: akun.TitlePattern,
	}
	spinhandler := spin.NewSpinHandler(
		&akunlegacy,
		flow.ConfigFlow.FilterText,
		flow.ConfigFlow.TitleConfig,
		flow.ConfigFlow.DescConfig,
		flow.ConfigFlow.SpinDataRepo,
		flow.ConfigFlow.HastagRepo,
		flow.ConfigFlow.MarkupConfig,
	)

	return spinhandler
}

func (flow *ShopeeToTopedFlow) Run() error {
	log.Println("running Tokopedia Upload...")
	err := flow.AkunIterator.Reset()
	if err != nil {
		return err
	}

MainLoop:
	for {
		select {
		case flow.limitGuard <- 1:
			go flow.RunTask()
		case <-flow.Ctx.Done():
			break MainLoop
		}
	}

	return nil

}

func (flow *ShopeeToTopedFlow) GenerateHandler(akun *repo.AkunItem, spin shopeeuploader.SpinFunc) []uploader.UploadHandler {
	handlers := []uploader.UploadHandler{
		flow.createProductHandler(akun, spin),
		flow.createSpinHandler(akun, spin),
		flow.createAnnotationHandler(),
		flow.createImageHandler(),
		flow.createCategoryHandler(),
		flow.createEtalaseHandler(),
		flow.createVariantHandler(spin),
	}
	return handlers
}
