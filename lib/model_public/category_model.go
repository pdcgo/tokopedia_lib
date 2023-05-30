package model_public

type Categories struct {
	ID       int          `json:"id"`
	Name     string       `json:"name"`
	URL      string       `json:"url,omitempty"`
	Children []Categories `json:"children,omitempty"`
	Typename string       `json:"__typename,omitempty"`
}

type JarvisRecommendationVar struct {
	ProductName string `json:"productName"`
}

type JarvisRecommendationResp struct {
	Data struct {
		GetJarvisRecommendation struct {
			Categories []Categories `json:"categories"`
			Typename   string       `json:"__typename"`
		} `json:"getJarvisRecommendation"`
	} `json:"data"`
}

type CategoryAllListLiteResp struct {
	Data struct {
		CategoryAllListLite struct {
			Categories []Categories `json:"categories"`
			Typename   string       `json:"__typename"`
		} `json:"categoryAllListLite"`
	} `json:"data"`
}
