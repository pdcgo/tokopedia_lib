package model_public

type HeaderErr struct {
	ProcessTime string   `json:"processTime"`
	Messages    []string `json:"messages"`
	Reason      string   `json:"reason"`
	ErrorCode   string   `json:"errorCode"`
	Typename    string   `json:"__typename"`
}

type DataValue struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Selected bool   `json:"selected"`
	Typename string `json:"__typename"`
}

type AnnotationData struct {
	Variant   string      `json:"variant"`
	SortOrder int         `json:"sortOrder"`
	Values    []DataValue `json:"values"`
	Typename  string      `json:"__typename"`
}

type AnnotationDataList []*AnnotationData

func (list AnnotationDataList) FindValue(variant string, valueName string) *DataValue {
	for _, anno := range list {
		if anno.Variant == variant {
			for _, val := range anno.Values {
				if val.Name == valueName {
					return &val
				}
			}
		}
	}
	return nil
}

type DrogonAnnotationCategoryV2 struct {
	Header    HeaderErr          `json:"header"`
	ProductID int64              `json:"productID"`
	Data      AnnotationDataList `json:"data"`
	Typename  string             `json:"__typename"`
}

type GetDrogonAnnotationVar struct {
	CategoryID       string `json:"categoryID"`
	ExcludeSensitive string `json:"excludeSensitive"`
	ProductID        string `json:"productID"`
	VendorName       string `json:"vendorName"`
}

type GetDrogonAnnotationResp struct {
	Data struct {
		DrogonAnnotationCategoryV2 DrogonAnnotationCategoryV2 `json:"drogonAnnotationCategoryV2"`
	} `json:"data"`
}
