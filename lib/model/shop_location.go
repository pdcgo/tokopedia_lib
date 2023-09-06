package model

type GeneralTicker struct {
	Header       string `json:"header"`
	Body         string `json:"body"`
	BodyLinkText string `json:"body_link_text"`
	BodyLinkURL  string `json:"body_link_url"`
	Typename     string `json:"__typename"`
}

type ShopID struct {
	Int64    int    `json:"int64"`
	Valid    bool   `json:"valid"`
	Typename string `json:"__typename"`
}

type PartnerID struct {
	Int64    int    `json:"int64"`
	Valid    bool   `json:"valid"`
	Typename string `json:"__typename"`
}

type Ticker struct {
	TextInactive       string `json:"text_inactive"`
	TextCourierSetting string `json:"text_courier_setting"`
	LinkCourierSetting string `json:"link_courier_setting"`
	Typename           string `json:"__typename"`
}

type ShopLocationWarehouse struct {
	WarehouseID         int       `json:"warehouse_id"`
	WarehouseName       string    `json:"warehouse_name"`
	WarehouseType       int       `json:"warehouse_type"`
	ShopID              ShopID    `json:"shop_id"`
	PartnerID           PartnerID `json:"partner_id"`
	AddressDetail       string    `json:"address_detail"`
	PostalCode          int       `json:"postal_code,string"`
	Latlon              string    `json:"latlon"`
	DistrictID          int       `json:"district_id"`
	DistrictName        string    `json:"district_name"`
	CityID              int       `json:"city_id"`
	CityName            string    `json:"city_name"`
	ProvinceID          int       `json:"province_id"`
	ProvinceName        string    `json:"province_name"`
	Country             string    `json:"country"`
	Status              int       `json:"status"`
	IsCoveredByCouriers bool      `json:"is_covered_by_couriers"`
	Ticker              Ticker    `json:"ticker"`
	Typename            string    `json:"__typename"`
}

type ShopLocationLegacy struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	DistrictId   int    `json:"districtId"`
	DistrictName string `json:"districtName"`
	CityId       int    `json:"cityId"`
	CityName     string `json:"cityName"`
	StateId      int    `json:"stateId"`
	StateName    string `json:"stateName"`
	PostalCode   int    `json:"postalCode"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Fax          string `json:"fax"`
	Area         string `json:"area"`
	Position     int    `json:"position"`
}

func (w *ShopLocationWarehouse) CreateShopLocation() ShopLocationLegacy {
	return ShopLocationLegacy{
		Id:           w.WarehouseID,
		Name:         w.WarehouseName,
		Address:      w.AddressDetail,
		DistrictId:   w.DistrictID,
		DistrictName: w.DistrictName,
		CityId:       w.CityID,
		CityName:     w.CityName,
		StateId:      w.ProvinceID,
		StateName:    w.ProvinceName,
		PostalCode:   w.PostalCode,
		Email:        "",
		Phone:        "",
		Fax:          "",
		Area:         "",
		Position:     0,
	}
}

type ShopLocationWarehouses []ShopLocationWarehouse

func (warehouses ShopLocationWarehouses) GetLocations() []ShopLocationLegacy {

	locations := []ShopLocationLegacy{}
	for _, warehouse := range warehouses {
		shopLocation := warehouse.CreateShopLocation()
		locations = append(locations, shopLocation)
	}

	return locations
}

type ShopLocationAllData struct {
	GeneralTicker GeneralTicker          `json:"general_ticker"`
	Warehouses    ShopLocationWarehouses `json:"warehouses"`
	Typename      string                 `json:"__typename"`
}

type ShopLocationAll struct {
	Status   int                 `json:"status"`
	Message  string              `json:"message"`
	Error    Error               `json:"error"`
	Data     ShopLocationAllData `json:"data"`
	Typename string              `json:"__typename"`
}
