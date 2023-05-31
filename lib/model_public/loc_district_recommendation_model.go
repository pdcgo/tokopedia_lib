package model_public

type District struct {
	DistrictID   int      `json:"districtId"`
	DistrictName string   `json:"district_name"`
	CityID       int      `json:"cityId"`
	CityName     string   `json:"city_name"`
	ProvinceID   int      `json:"provinceId"`
	ProvinceName string   `json:"province_name"`
	ZipCode      []string `json:"zip_code"`
	Typename     string   `json:"__typename"`
}

type KeroDistrictRecommendation struct {
	NextAvailable bool       `json:"next_available"`
	District      []District `json:"district"`
	Typename      string     `json:"__typename"`
}

type LocDisctricRecommendationVar struct {
	Page  string `json:"page"`
	Query string `json:"query"`
}

type LocDisctricRecommendationResp struct {
	Data struct {
		KeroDistrictRecommendation KeroDistrictRecommendation `json:"kero_district_recommendation"`
	} `json:"data"`
}
