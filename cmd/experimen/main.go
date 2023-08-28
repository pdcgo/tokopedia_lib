package main

import (
	"context"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	"github.com/pdcgo/go_v2_shopeelib/lib/upload_config"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/app/grab"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
)

func RunGrab(ctx context.Context) error {
	api, err := api_public.NewTokopediaApiPublic()
	if err != nil {
		return err
	}

	base := &legacy_source.BaseConfig{
		BaseData: "./../../pdcgo/pdc_launcher/test_base",
	}

	botConfig := upload_config.NewBotConfig(base)

	db := mongorepo.NewDatabase(ctx, botConfig.Database.DBUri, botConfig.Database.DBName)
	productRepo := mongorepo.NewProductRepo(db)

	app := grab.NewGrabApp(api, base, productRepo)

	return app.Run()
}

func main() {
	driver, _ := tokopedia_lib.NewDriverAccount("farikahmad84@gmail.com", "Jatim123", "5OWGYNU7TM3C34SDDDJ7XVC3M4SAI6H4")

	driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {
		driver.SellerLogin(dctx)

		// time.Sleep(time.Hour)

		return nil
	})

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// if err := RunGrab(ctx); err != nil {
	// 	log.Fatal(err)
	// }
}
