package main

import (
	"context"
	"log"
	"path/filepath"

	"github.com/pdcgo/common_conf/pdc_common"
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

	cfgname := "data/config.json"
	pdc_common.SetConfig(cfgname, appcfg.Version, "golang_tokopedia_upload", appcfg.Cred)
	pdc_common.InitializeLogger()

	rootBase := ctx.String("b")

	cfg := config.NewUploadConfigBase(rootBase)

	log.Println("running on ", rootBase)

	sqlpath := filepath.Join(rootBase, "tokopedia_data.db")
	sqlitedb := datasource.NewSqliteDatabase(sqlpath)

	concurent := shopee_flow.CreateConfigConcurencyFromCmd()

	mdb := mongorepo.NewDatabase(context.Background(), cfg.Database.DbURI, cfg.Database.DbName)

	publicapi, err := api_public.NewTokopediaApiPublic()
	if err != nil {
		return err
	}

	shopeeagg := shopee_repo.NewProductAggregate(mdb.Collection("item"))

	flow := shopee_flow.NewShopeeToTopedFlow(
		rootBase,
		context.Background(),
		mdb,
		sqlitedb,
		concurent,
		publicapi,
		shopeeagg,
	)

	return flow.Run()

}
