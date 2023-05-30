package model_public

type Meta struct {
	AbTest     string `json:"ab_test"`
	Templating string `json:"templating"`
	Typename   string `json:"__typename"`
}

type Image struct {
	ImageURL        string `json:"imageUrl"`
	TrackerImageURL string `json:"trackerImageUrl"`
	Typename        string `json:"__typename"`
}

type ShopBadge struct {
	Title    string `json:"title"`
	ImageURL string `json:"imageURL"`
	Show     bool   `json:"show"`
	Typename string `json:"__typename"`
}

type AdParamsVar struct {
	AdParams string `json:"adParams"`
}

type TopadsProductQueryResp struct {
	Data struct {
		DisplayAdsV3 struct {
			Data []struct {
				ClickTrackURL      string `json:"clickTrackUrl"`
				ProductWishlistURL string `json:"product_wishlist_url"`
				Product            struct {
					ID                 string        `json:"id"`
					Name               string        `json:"name"`
					Wishlist           bool          `json:"wishlist"`
					Image              Image         `json:"image"`
					URL                string        `json:"url"`
					RelativeURI        string        `json:"relative_uri"`
					Price              string        `json:"price"`
					WholeSalePrice     []interface{} `json:"wholeSalePrice"`
					CountTalkFormat    string        `json:"count_talk_format"`
					CountReviewFormat  string        `json:"countReviewFormat"`
					Category           Category      `json:"category"`
					CategoryBreadcrumb string        `json:"categoryBreadcrumb"`
					Preorder           bool          `json:"preorder"`
					ProductWholesale   bool          `json:"product_wholesale"`
					FreeReturn         string        `json:"free_return"`
					IsNewProduct       bool          `json:"isNewProduct"`
					Cashback           string        `json:"cashback"`
					Rating             int           `json:"rating"`
					RatingAverage      string        `json:"ratingAverage"`
					TopLabel           []interface{} `json:"top_label"`
					BottomLabel        []interface{} `json:"bottomLabel"`
					LabelGroups        []LabelGroups `json:"labelGroups"`
					Campaign           struct {
						DiscountPercentage int    `json:"discountPercentage"`
						OriginalPrice      string `json:"originalPrice"`
						Typename           string `json:"__typename"`
					} `json:"campaign"`
					CustomvideoURL string `json:"customvideo_url"`
					Typename       string `json:"__typename"`
				} `json:"product"`
				Shop     ProductShop `json:"shop"`
				Tag      int         `json:"tag"`
				Typename string      `json:"__typename"`
			} `json:"data"`
			Header struct {
				Meta     Meta   `json:"meta"`
				Typename string `json:"__typename"`
			} `json:"header"`
			Typename string `json:"__typename"`
		} `json:"displayAdsV3"`
	} `json:"data"`
}
