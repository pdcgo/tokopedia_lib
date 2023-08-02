package grabber_test

// func TestProductGrab(t *testing.T) {
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

// 	t.Run("test product keyword grabber", func(t *testing.T) {
// 		keywords := []string{
// 			"mousepad",
// 			"keyboard gaming",
// 		}
// 		grabber := grabber.NewProductListGrabber(baseGrabber, keywords)
// 		params := grabber.GenerateProductSearchParams()
// 		params.Query = keywords[0]
// 		products, err := grabber.GetProducts(params)
// 		assert.Nil(t, err)
// 		assert.NotEqual(t, len(products), 0)
// 	})
// 	t.Run("test product category grabber first level", func(t *testing.T) {
// 		// 1759, Fashion Pria
// 		baseGrabber.GrabTasker.TokpedCateg = []string{"1759"}
// 		grabber := grabber.NewCategoryGrabber(baseGrabber)
// 		params := grabber.GenerateProductSearchParams()
// 		params.CategoryId = 1759

// 		products, err := grabber.GetProducts(params)
// 		assert.Nil(t, err)
// 		assert.NotEqual(t, len(products), 0)
// 		assert.Equal(t, products[0].CategoryID, 1759)
// 	})
// 	t.Run("test product category grabber last level", func(t *testing.T) {
// 		// 297, "Komputer & Laptop"
// 		// 338, "Aksesoris Komputer & Laptop",
// 		// 340, "Keyboard"
// 		baseGrabber.GrabTasker.TokpedCateg = []string{"340"}
// 		grabber := grabber.NewCategoryGrabber(baseGrabber)
// 		params := grabber.GenerateProductSearchParams()
// 		params.CategoryId = 340

// 		products, err := grabber.GetProducts(params)
// 		assert.Nil(t, err)
// 		assert.NotEqual(t, len(products), 0)
// 		assert.Equal(t, products[0].Category, 340)
// 	})
// }
