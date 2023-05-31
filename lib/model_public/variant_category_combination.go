package model_public

type VariantUnitValue struct {
	VariantUnitValueID int    `json:"VariantUnitValueID"`
	Status             int    `json:"Status"`
	Value              string `json:"Value"`
	EquivalentValueID  int    `json:"EquivalentValueID"`
	EnglishValue       string `json:"EnglishValue"`
	Hex                string `json:"Hex"`
	Typename           string `json:"__typename"`
}

type VariantDetailUnit struct {
	VariantUnitID int                `json:"VariantUnitID"`
	Status        int                `json:"Status"`
	UnitName      string             `json:"UnitName"`
	UnitShortName string             `json:"UnitShortName"`
	UnitValues    []VariantUnitValue `json:"UnitValues"`
	Typename      string             `json:"__typename"`
}

type VariantDetail struct {
	VariantID  int                 `json:"VariantID"`
	HasUnit    int                 `json:"HasUnit"`
	Identifier string              `json:"Identifier"`
	Name       string              `json:"Name"`
	Status     int                 `json:"Status"`
	Units      []VariantDetailUnit `json:"Units"`
	Typename   string              `json:"__typename"`
}

type AllVariant struct {
	VariantID int    `json:"VariantID"`
	Name      string `json:"Name"`
	Typename  string `json:"__typename"`
}

type GetVariantCategoryCombinationData struct {
	CategoryID            int             `json:"categoryID"`
	VariantIDCombinations [][]int         `json:"variantIDCombinations"`
	AllVariants           []AllVariant    `json:"allVariants"`
	VariantDetails        []VariantDetail `json:"variantDetails"`
	Typename              string          `json:"__typename"`
}

type GetVariantCategoryCombination struct {
	Header   HeaderErr                         `json:"header"`
	Data     GetVariantCategoryCombinationData `json:"data"`
	Typename string                            `json:"__typename"`
}

type VariantCategoryCombinationVar struct {
	AllVariants     string `json:"allVariants"`
	CategoryID      int    `json:"categoryID"`
	ProductVariants string `json:"productVariants"`
	Type            string `json:"type"`
}

type VariantCategoryCombinationResp struct {
	Data struct {
		GetVariantCategoryCombination GetVariantCategoryCombination `json:"getVariantCategoryCombination"`
	} `json:"data"`
}
