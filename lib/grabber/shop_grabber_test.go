package grabber_test

// func TestShopGrabber(t *testing.T) {
// 	api, err := api_public.NewTokopediaApiPublic()
// 	assert.Nil(t, err)
// 	ctx := context.Background()
// 	ctx, cancel := context.WithCancel(ctx)
// 	defer cancel()

// 	baseConfig := &legacy_source.BaseConfig{
// 		BaseData: "../..",
// 	}
// 	database := scenario.GetMongoDatabase(t)

// 	productRepo := mongorepo.NewProductRepo(ctx, database)
// 	cacheHandler := grab_handler.NewCacheProductHandler(productRepo)
// 	tasker := legacy.NewGrabTasker(baseConfig.Path("data/tasker.json"))
// 	baseGrabber := grabber.NewBaseGrabber(api, baseConfig, tasker, cacheHandler)
// 	baseGrabber.Filter.GrabBasic.LimitGrab = 1

// 	t.Run("test grab shop", func(t *testing.T) {
// 		grab := grabber.NewShopListGrabber(baseGrabber)
// 		grab.Run()
// 	})
// }
