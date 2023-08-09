package model_public

type HeaderMainDataResp struct {
	Data struct {
		DynamicHomeIcon struct {
			CategoryGroup []struct {
				ID           int    `json:"id"`
				Title        string `json:"title"`
				Desc         string `json:"desc"`
				CategoryRows []struct {
					ID         int    `json:"id"`
					Name       string `json:"name"`
					URL        string `json:"url"`
					ImageURL   string `json:"imageUrl"`
					Type       string `json:"type"`
					CategoryID int    `json:"categoryId"`
					Typename   string `json:"__typename"`
				} `json:"categoryRows"`
				Typename string `json:"__typename"`
			} `json:"categoryGroup"`
			Typename string `json:"__typename"`
		} `json:"dynamicHomeIcon"`
		CategoryAllListLite struct {
			Categories []*Categories `json:"categories"`
			Typename   string        `json:"__typename"`
		} `json:"categoryAllListLite"`
	} `json:"data"`
}
