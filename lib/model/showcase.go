package model

type ShopShowcasesResult struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Count       int    `json:"count"`
	Type        int    `json:"type"`
	Highlighted bool   `json:"highlighted"`
	Alias       string `json:"alias"`
	URI         string `json:"uri"`
	UseAce      bool   `json:"useAce"`
	Badge       string `json:"badge"`
	ImageURL    string `json:"imageURL"`
	IsFeatured  any    `json:"isFeatured"`
	Typename    string `json:"__typename"`
}

type ShopShowcases struct {
	Error    ErrorMsg              `json:"error"`
	Result   []ShopShowcasesResult `json:"result"`
	Typename string                `json:"__typename"`
}

type GetFeaturedShowcase struct {
	Error struct {
		Message  string `json:"message"`
		Typename string `json:"__typename"`
	} `json:"error"`
	Result   []any  `json:"result"`
	Typename string `json:"__typename"`
}

type ShopShowcasesData struct {
	ShopShowcases       ShopShowcases       `json:"shopShowcases"`
	GetFeaturedShowcase GetFeaturedShowcase `json:"getFeaturedShowcase"`
}
