package model_public

type CategoryRow struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	ImageURL   string `json:"imageUrl"`
	Type       string `json:"type"`
	CategoryID int    `json:"categoryId"`
	Typename   string `json:"__typename"`
}

type CategoryGroup struct {
	ID           int           `json:"id"`
	Title        string        `json:"title"`
	Desc         string        `json:"desc"`
	CategoryRows []CategoryRow `json:"categoryRows"`
	Typename     string        `json:"__typename"`
}

type DynamicHomeIcon struct {
	CategoryGroup []CategoryGroup `json:"categoryGroup"`
	Typename      string          `json:"__typename"`
}

type CategoriesData []*Categories

type CategoriesHandler func(parents []*Categories, category *Categories) (stop bool, err error)

func (data CategoriesData) iterateCategory(parents []*Categories, handler CategoriesHandler) error {
	for _, category := range data {

		stop, err := handler(parents, category)
		if stop {
			return nil
		}
		if err != nil {
			return err
		}

		if len(category.Children) > 0 {
			childrenParents := append(parents, category)
			childrens := CategoriesData(category.Children)

			err := childrens.iterateCategory(childrenParents, handler)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (data CategoriesData) Iterate(handler CategoriesHandler) error {

	err := data.iterateCategory([]*Categories{}, handler)
	return err
}

func (data CategoriesData) GetCategoryByUrl(url string) (*Categories, error) {

	result := &Categories{}
	err := data.Iterate(func(parents []*Categories, category *Categories) (stop bool, err error) {

		if category.URL == url {
			result = category
			return true, nil
		}

		return false, nil
	})

	return result, err
}

type HeaderMainData struct {
	DynamicHomeIcon     DynamicHomeIcon         `json:"dynamicHomeIcon"`
	CategoryAllListLite CategoryAllListLiteData `json:"categoryAllListLite"`
}

type HeaderMainDataResp struct {
	Data HeaderMainData `json:"data"`
}
