package grabber

// type ProductGrabber struct {
// 	*Grabber
// }

// func arrayConverter(datas []interface{}) []string {
// 	results := make([]string, len(datas))
// 	for _, data := range datas {
// 		switch value := data.(type) {
// 		case int:
// 			results = append(results, fmt.Sprint(value))
// 		case string:
// 			results = append(results, value)
// 		}
// 	}
// 	return results
// }

// func (grab *ProductGrabber) GetProducts(params *model_public.SearchProductVar) ([]*model_public.ProductSearch, error) {

// 	rawParams, err := json.Marshal(params)
// 	if err != nil {
// 		return nil, err
// 	}
// 	stringParams := string(rawParams)
// 	replacer := strings.NewReplacer(`"`, "", ":", "=", ",", "&", " ", "", "[", "", "]", "")
// 	validParams := replacer.Replace(stringParams)

// 	variable := &model_public.ParamsVar{
// 		Params: validParams[1 : len(validParams)-1],
// 	}
// 	resp, err := grab.Api.SearchProductQueryV4(variable)
// 	if err != nil {
// 		return nil, err
// 	}
// 	products := resp.Data.AceSearchProductV4.Data.Products

// 	return products, nil
// }

// func (grab *ProductGrabber) ProcessProduct(product *model_public.ProductSearch) error {
// 	prodVar, _ := ParseProductDetailParamsFromUrl(product.URL)

// 	if grab.GrabTasker.UseFilter {
// 		if grab.AppliedFilterShop(product.Shop.ShopID, prodVar.ShopDomain) {
// 			return errors.New("terkena filter toko")
// 		}

// 		if grab.AppliedFilterProduct(int(product.ID), product.Name, product.URL) {
// 			return errors.New("terkena filter produk")
// 		}
// 	}

// 	product_detail, err := grab.GetPublicProductLayout(product.URL)
// 	if err != nil {
// 		fmt.Printf("error [ produk ] : error  mendapatkan produk [ %s ]\n", product.Name)
// 		return err
// 	}
// 	err = grab.Save(&grab_handler.ProductListGrabberResp{Product: product, ProductDetail: product_detail})
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (grab *ProductGrabber) IterateProductPages(params *model_public.SearchProductVar) (<-chan *model_public.ProductSearch, *helper.ChannelError) {
// 	res := make(chan *model_public.ProductSearch)
// 	errChan := helper.NewChannelError()

// 	go func() {
// 		defer close(res)
// 		defer errChan.Raise()

// 	Parent:
// 		for {
// 			products, err := grab.GetProducts(params)
// 			if err != nil {
// 				fmt.Println(err)
// 				break Parent
// 			}
// 			if len(products) == 0 {
// 				fmt.Printf("grab [ produk ] : produk telah habis\n")
// 				break Parent
// 			}

// 			for _, product := range products {
// 				res <- product
// 			}
// 			params.Page += 1
// 			params.Start = params.Page * params.Rows
// 		}
// 	}()
// 	return res, errChan
// }

// func (grab *ProductGrabber) Save(product *grab_handler.ProductListGrabberResp) error {
// 	return grab.CacheHandler.AddItemProductSearch(grab.GrabTasker.Namespace, product)
// }

// // Product List
// type ProductListGrabber struct {
// 	ProductGrabber
// 	Keywords []string
// }

// func NewProductListGrabber(
// 	grabber *Grabber,
// 	keywords []string) *ProductListGrabber {
// 	return &ProductListGrabber{
// 		ProductGrabber: ProductGrabber{
// 			Grabber: grabber,
// 		},
// 		Keywords: keywords,
// 	}
// }

// func (grab *ProductListGrabber) Run() {

// Keywords:
// 	for _, keyword := range grab.Keywords {
// 		fmt.Printf("grab [ keyword ] : memulai grab keyword [ %s ]\n", keyword)
// 		params := grab.GenerateProductSearchParams()
// 		params.Query = url.QueryEscape(keyword)

// 		limit := int32(grab.Filter.GrabBasic.LimitGrab)
// 		limiter := helper.NewLimiter(limit)
// 		counter := helper.Counter{}

// 		products, errChan := grab.IterateProductPages(params)
// 		for product := range products {
// 			if limiter.LimitReached() {
// 				continue Keywords
// 			}

// 			err := grab.ProcessProduct(product)
// 			if err != nil {
// 				continue
// 			}

// 			fmt.Printf("grab [ keyword ] : mendapatkan produk [ %s ] [ %d ]\n", product.Name, counter.Total)
// 			limiter.Add()
// 			counter.Add()
// 		}

// 		err := errChan.GetError()
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 	}
// }

// type CategoryGrabber struct {
// 	ProductGrabber
// 	CatIds []string
// }

// func NewCategoryGrabber(grabber *Grabber) *CategoryGrabber {
// 	return &CategoryGrabber{
// 		ProductGrabber: ProductGrabber{
// 			Grabber: grabber,
// 		},
// 		CatIds: grabber.GrabTasker.TokpedCateg,
// 	}
// }

// func (grab *CategoryGrabber) Run() error {

// Category:
// 	for _, c := range grab.CatIds {
// 		catId, _ := strconv.Atoi(c)
// 		params := grab.GenerateProductSearchParams()
// 		params.CategoryId = catId

// 		limit := int32(grab.Filter.GrabBasic.LimitGrab)
// 		limiter := helper.NewLimiter(limit)

// 		products, errChan := grab.IterateProductPages(params)
// 		for product := range products {
// 			if limiter.LimitReached() {
// 				fmt.Println("filter [ produk ] : telah mencapai batas grab")
// 				continue Category
// 			}

// 			err := grab.ProcessProduct(product)
// 			if err != nil {
// 				continue
// 			}

// 			fmt.Printf("grab [ kategori ] : mendapatkan produk [ %s ]\n", product.Name)
// 			limiter.Add()
// 		}
// 		err := errChan.GetError()
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			return err
// 		}
// 	}

// 	return nil
// }

// // type CategoryCsvGrabber struct {
// // 	CategoryGrabber
// // 	*dumper.CategoryCsvDumper
// // }

// // func NewCategoryCsvGrabber(grabber *Grabber, pathfile string) *CategoryCsvGrabber {
// // 	return &CategoryCsvGrabber{
// // 		CategoryGrabber: CategoryGrabber{
// // 			ProductGrabber: ProductGrabber{
// // 				Grabber: grabber,
// // 			},
// // 		},
// // 		CategoryCsvDumper: dumper.NewCategoryCsvDumper(grabber.Api, grabber.Base, pathfile),
// // 	}
// // }

// func (g *CategoryCsvGrabber) Run() error {
// 	categories, err := g.Load()
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	fmt.Printf("grab [ file kategori ] : memulai grab file kategori\n")
// Categories:
// 	for _, c := range categories {
// 		if c.Status == "grabbed" {
// 			continue Categories
// 		}

// 		categ := g.GetCategoryByUrl(nil, c.Url)
// 		for i := range categ {
// 			g.CatIds = []string{fmt.Sprintf("%d", i.ID)}
// 			break
// 		}

// 		fmt.Printf("grab [ file kategori ] : memulai grab dari kategori [ %s ]\n", c.Name)
// 		g.CategoryGrabber.Run()
// 		c.Status = "grabbed"
// 		g.Save()
// 	}
// 	return nil
// }
