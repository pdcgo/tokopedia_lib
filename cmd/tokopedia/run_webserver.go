package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/controller"
	mongolib "github.com/pdcgo/go_v2_shopeelib/lib/mongo"
	"github.com/pdcgo/tokopedia_lib/app/upload_app"
	"github.com/pdcgo/tokopedia_lib/app/upload_app/config"
	"github.com/pdcgo/tokopedia_lib/app/web"
	"github.com/pdcgo/tokopedia_lib/app/web/api"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/app_builder"
	"github.com/pdcgo/tokopedia_lib/lib/category_mapper"
	"github.com/pdcgo/tokopedia_lib/lib/datasource"
	"github.com/pdcgo/tokopedia_lib/lib/repo"
	"github.com/pdcgo/v2_gots_sdk"
	"github.com/urfave/cli/v2"
)

type TokopediaWebServer struct {
	Base    string
	DevMode bool
}

func (webtoped *TokopediaWebServer) SetupRouter(r *gin.Engine, prefix string) error {

	dbpath := filepath.Join(webtoped.Base, "tokopedia_data.db")
	db := datasource.NewSqliteDatabase(dbpath)

	app := upload_app.NewUploadApp(db)
	repo := repo.NewAkunRepo(db)

	sdk := v2_gots_sdk.NewApiSdk(r)

	save := func() {}
	if webtoped.DevMode {
		BuildExeOnDev(webtoped.Base)
		save = sdk.GenerateSdkFunc("frontend/src/client/sdk_types_test.ts", true)
	}

	defer save()

	g := sdk.Group(prefix)
	RegisterAkunApi(g, db, repo)
	RegisterCommand(g, app, webtoped.Base)

	// bagian hendra
	validate := validator.New()
	baseData := &legacy_source.BaseConfig{
		BaseData: webtoped.Base,
	}

	cfg := config.NewUploadConfigBase(webtoped.Base)
	mdb := mongolib.NewDatabase(context.Background(), cfg.Database.DbURI, cfg.Database.DbName)
	base := controller.NewBaseController(validate, baseData, nil, mdb)
	controller.RegisterSpinController(sdk, base)
	controller.RegisterMarkupController(sdk, base)
	controller.RegisterProductController(sdk, base)

	productRepo := mongolib.NewProductRepo(context.TODO(), mdb)
	pubapi, err := api_public.NewTokopediaApiPublic()

	if err != nil {
		pdc_common.ReportError(err)
	}

	mapperdata := category_mapper.NewMapper(pubapi)
	api.RegisterShopeeTopedMap(g, db, productRepo, mapperdata)
	api.RegisterCategoryApi(g, baseData)

	web.RegisterTokopediaFrontend(r, prefix)

	return nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func BuildExeOnDev(base string) {
	basebin := filepath.Join(base, "./bin")
	base, err := app_builder.BuildBynaryCmd(basebin, "./cmd/tokopedia", "tokopedia.exe")
	log.Println("bin created on ", base)
	if err != nil {
		panic(err)
	}
}

func runWebServer(ctx *cli.Context) error {
	var devmode = os.Getenv("DEV_MODE") != ""

	roopAplicationPath := ctx.String("b")
	log.Println("using root aplication path..", roopAplicationPath)

	if !devmode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(CORSMiddleware())

	server := TokopediaWebServer{
		Base:    roopAplicationPath,
		DevMode: devmode,
	}

	server.SetupRouter(r, "tokopedia")

	r.Run("localhost:8080")

	return nil
}