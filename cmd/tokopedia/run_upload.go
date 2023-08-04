package main

import (
	"context"

	"github.com/pdcgo/common_conf/pdc_application"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	appcfg "github.com/pdcgo/tokopedia_lib/app/config"
	"github.com/pdcgo/tokopedia_lib/app/shopee/shopee_repo"
	"github.com/pdcgo/tokopedia_lib/app/upload_app/config"
	"github.com/pdcgo/tokopedia_lib/app/upload_app/shopee_flow"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/datasource"
	"github.com/urfave/cli/v2"
)

func runUploadShopeeToped(ctx *cli.Context) error {

	app := pdc_application.PdcApplication{
		Base: &legacy_source.BaseConfig{
			BaseData: ctx.String("base"),
		},
		Credential:    Cred,
		Version:       appcfg.Version,
		ReplaceLogger: true,
		AppID:         AppID,
	}

	app.RunWithLicenseFile("data/config.json", "golang_tokopedia_upload", func(app *pdc_application.PdcApplication) {

		cfg := config.NewUploadConfigBase(app.Base.Path())
		sqlitedb := datasource.NewSqliteDatabase(app.Base.Path("tokopedia_data.db"))

		concurent := shopee_flow.CreateConfigConcurencyFromCmd()

		mdb := mongorepo.NewDatabase(context.Background(), cfg.Database.DbURI, cfg.Database.DbName)

		publicapi, err := api_public.NewTokopediaApiPublic()
		if err != nil {
			pdc_common.ReportError(err)
			return
		}

		shopeeagg := shopee_repo.NewProductAggregate(mdb.Collection("item"))

		flow := shopee_flow.NewShopeeToTopedFlow(
			app.Base.Path(),
			context.Background(),
			mdb,
			sqlitedb,
			concurent,
			publicapi,
			shopeeagg,
		)

		flow.Run()

	})

	// ----------------------

	return nil

}
