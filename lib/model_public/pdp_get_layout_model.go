package model_public

type Tokonow struct {
	ShopID      string `json:"shopID"`
	WhID        string `json:"whID"`
	ServiceType string `json:"serviceType"`
}

type UserLocation struct {
	CityID     string `json:"cityID"`
	AddressID  string `json:"addressID"`
	DistrictID string `json:"districtID"`
	PostalCode string `json:"postalCode"`
	Latlon     string `json:"latlon"`
}

type PdpGetlayoutQueryVar struct {
	ShopDomain   string       `json:"shopDomain"`
	ProductKey   string       `json:"productKey"`
	LayoutID     string       `json:"layoutID"`
	APIVersion   int          `json:"apiVersion"`
	Tokonow      Tokonow      `json:"tokonow"`
	UserLocation UserLocation `json:"userLocation"`
	ExtParam     string       `json:"extParam"`
}

type PdpGetlayoutQueryResp struct {
	Data struct {
		PdpGetLayout PdpGetLayout `json:"pdpGetLayout"`
	} `json:"data"`
}
