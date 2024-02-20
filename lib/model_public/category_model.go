package model_public

type BaseCategories struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	IconImageUrl string `json:"iconImageUrl"`
	IsCrawlable  int    `json:"isCrawlable"`
	URL          string `json:"url,omitempty"`
}

type Categories struct {
	*BaseCategories
	Children []*Categories `json:"children,omitempty"`
	Typename string        `json:"__typename,omitempty"`
}

type JarvisRecommendationVar struct {
	ProductName string `json:"productName"`
}

type GetJarvisRecommendation struct {
	Categories []Categories `json:"categories"`
	Typename   string       `json:"__typename"`
}

type JarvisRecommendationData struct {
	GetJarvisRecommendation `json:"getJarvisRecommendation"`
}

type JarvisRecommendationResp struct {
	Data JarvisRecommendationData `json:"data"`
}

type CategoryAllListLiteData struct {
	Categories CategoriesData `json:"categories"`
	Typename   string         `json:"__typename"`
}

type BulkCatItem struct {
	*BaseCategories
	ParentID int `json:"parent_id"`
}

func (data *CategoryAllListLiteData) GetBulkCats(payload []int) ([][]*BulkCatItem, error) {

	hasil := make([][]*BulkCatItem, len(payload))

	var getCategories func(categs []*Categories) []int

	getCategories = func(categs []*Categories) []int {
		indexs := []int{}

		for _, categ := range categs {
			haveChild := len(categ.Children) != 0

			if haveChild {
				cindexs := getCategories(categ.Children)
				for _, ind := range cindexs {
					chains := hasil[ind]
					chains[0].ParentID = categ.ID

					hasil[ind] = append([]*BulkCatItem{
						{
							BaseCategories: categ.BaseCategories,
						},
					}, chains...)
				}
				indexs = append(indexs, cindexs...)
			}

			for ind, pay := range payload {
				if categ.ID == pay {
					hasil[ind] = append(hasil[ind], &BulkCatItem{
						BaseCategories: categ.BaseCategories,
					})
					indexs = append(indexs, ind)
				}
			}

		}

		return indexs
	}

	getCategories(data.Categories)

	return hasil, nil
}

type CategoryAllListLiteRespData struct {
	CategoryAllListLite *CategoryAllListLiteData `json:"categoryAllListLite"`
}

type CategoryAllListLiteResp struct {
	Data *CategoryAllListLiteRespData `json:"data"`
}
