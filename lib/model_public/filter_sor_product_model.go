package model_public

type FilterSearch struct {
	Searchable  int    `json:"searchable"`
	Placeholder string `json:"placeholder"`
	Typename    string `json:"__typename"`
}

type FilterOption struct {
	Name        string        `json:"name"`
	Description string        `json:"Description"`
	Key         string        `json:"key"`
	Icon        string        `json:"icon"`
	Value       string        `json:"value"`
	InputType   string        `json:"inputType"`
	TotalData   string        `json:"totalData"`
	ValMax      string        `json:"valMax"`
	ValMin      string        `json:"valMin"`
	HexColor    string        `json:"hexColor"`
	Child       []interface{} `json:"child"`
	IsPopular   bool          `json:"isPopular"`
	IsNew       bool          `json:"isNew"`
	Typename    string        `json:"__typename"`
}

type SortProduct struct {
	Name        string `json:"name"`
	Key         string `json:"key"`
	Value       string `json:"value"`
	InputType   string `json:"inputType"`
	ApplyFilter string `json:"applyFilter"`
	Typename    string `json:"__typename"`
}

type FilterProduct struct {
	Title        string         `json:"title"`
	TemplateName string         `json:"template_name"`
	Search       FilterSearch   `json:"search"`
	Options      []FilterOption `json:"options"`
	Typename     string         `json:"__typename"`
}

type FilterSortProductData struct {
	Filter   []FilterProduct `json:"filter"`
	Sort     []SortProduct   `json:"sort"`
	Typename string          `json:"__typename"`
}

type FilterSortProduct struct {
	Data     FilterSortProductData `json:"data"`
	Typename string                `json:"__typename"`
}

type ParamsVar struct {
	Params string `json:"params"`
}

type FilterSortProductResp struct {
	Data struct {
		FilterSortProduct FilterSortProduct `json:"filter_sort_product"`
	} `json:"data"`
}
