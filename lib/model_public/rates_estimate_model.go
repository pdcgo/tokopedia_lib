package model_public

// weight: 0.25,
// domain: multiaccesoris, -> shop domain
// productId: 2248059401,
// origin: 5637|15810|YHeWJestVbuD66PoGVDnlQaHYJAZ8CG10NtXqKCEcuLjE3IncRxwVBiPSs6Tw6PTDkzUsF9kBXgJqw%3D%3D, -> geolocation PDPGetDataP2
// destination: 2935||-8.127172,112.417213, 
// POTime: 0,
// isFulfillment: false,
// deviceType: default_v3,
// shopTier: 3,
// bo_metadata: "",
// free_shipping_flag: 0,
// warehouse_id: 113930
type RatesEstimateQueryVar struct {
	Weight           float64 `json:"weight"`
	Domain           string  `json:"domain"`
	ProductID        string  `json:"productId"`
	Origin           string  `json:"origin"`
	Destination      string  `json:"destination"`
	POTime           int     `json:"POTime"`
	IsFulfillment    bool    `json:"isFulfillment"`
	DeviceType       string  `json:"deviceType"`
	ShopTier         int     `json:"shopTier"`
	BoMetadata       string  `json:"bo_metadata"`
	FreeShippingFlag int     `json:"free_shipping_flag"`
	WarehouseID      string  `json:"warehouse_id"`
}

//////////////////////////////////////////////

type Address struct {
	CityName     string `json:"city_name"`
	ProvinceName string `json:"province_name"`
	DistrictName string `json:"district_name"`
	Country      string `json:"country"`
	PostalCode   string `json:"postal_code"`
	Address      string `json:"address"`
	Lat          string `json:"lat"`
	Long         string `json:"long"`
	Phone        string `json:"phone"`
	AddrName     string `json:"addr_name"`
	Address1     string `json:"address_1"`
	ReceiverName string `json:"receiver_name"`
	Typename     string `json:"__typename"`
}

type RatesService struct {
	ServiceName  string `json:"service_name"`
	ServiceID    int    `json:"service_id"`
	ServiceOrder int    `json:"service_order"`
	Status       int    `json:"status"`
	RangePrice   struct {
		MinPrice int    `json:"min_price"`
		MaxPrice int    `json:"max_price"`
		Typename string `json:"__typename"`
	} `json:"range_price"`
	Texts struct {
		TextServiceDesc  string `json:"text_service_desc"`
		TextServiceNotes string `json:"text_service_notes"`
		TextRangePrice   string `json:"text_range_price"`
		TextEtd          string `json:"text_etd"`
		TextPrice        string `json:"text_price"`
		Typename         string `json:"__typename"`
	} `json:"texts"`
	Products []struct {
		ShipperName        string `json:"shipper_name"`
		ShipperID          int    `json:"shipper_id"`
		ShipperProductID   int    `json:"shipper_product_id"`
		ShipperProductName string `json:"shipper_product_name"`
		ShipperWeight      int    `json:"shipper_weight"`
		Price              struct {
			Price          int    `json:"price"`
			FormattedPrice string `json:"formatted_price"`
			Typename       string `json:"__typename"`
		} `json:"price"`
		Texts struct {
			TextEtd          string `json:"text_etd"`
			TextRangePrice   string `json:"text_range_price"`
			TextEtaSummarize string `json:"text_eta_summarize"`
			Typename         string `json:"__typename"`
		} `json:"texts"`
		Cod struct {
			IsCodAvailable int    `json:"is_cod_available"`
			Typename       string `json:"__typename"`
		} `json:"cod"`
		Eta struct {
			TextEta   string `json:"text_eta"`
			ErrorCode int    `json:"error_code"`
			Typename  string `json:"__typename"`
		} `json:"eta"`
		Features struct {
			DynamicPrice struct {
				TextLabel string `json:"text_label"`
				Typename  string `json:"__typename"`
			} `json:"dynamic_price"`
			Typename string `json:"__typename"`
		} `json:"features"`
		Typename string `json:"__typename"`
	} `json:"products"`
	ServiceBasedShipment struct {
		IsAvailable bool   `json:"is_available"`
		TextPrice   string `json:"text_price"`
		TextEta     string `json:"text_eta"`
		Typename    string `json:"__typename"`
	} `json:"service_based_shipment"`
	Cod struct {
		IsCod    int    `json:"is_cod"`
		CodText  string `json:"cod_text"`
		Typename string `json:"__typename"`
	} `json:"cod"`
	OrderPriority struct {
		IsNow    bool   `json:"is_now"`
		Typename string `json:"__typename"`
	} `json:"order_priority"`
	Etd struct {
		MinEtd   int    `json:"min_etd"`
		MaxEtd   int    `json:"max_etd"`
		Typename string `json:"__typename"`
	} `json:"etd"`
	Typename string `json:"__typename"`
}

type RatesEstimateV3 struct {
	Data struct {
		Address Address `json:"address"`
		Shop    struct {
			DistrictID   int    `json:"district_id"`
			DistrictName string `json:"district_name"`
			PostalCode   string `json:"postal_code"`
			Origin       int    `json:"origin"`
			AddrStreet   string `json:"addr_street"`
			Latitude     string `json:"latitude"`
			Longitude    string `json:"longitude"`
			ProvinceID   int    `json:"province_id"`
			CityID       int    `json:"city_id"`
			CityName     string `json:"city_name"`
			Typename     string `json:"__typename"`
		} `json:"shop"`
		Rates struct {
			ID       string         `json:"id"`
			RatesID  string         `json:"rates_id"`
			Type     string         `json:"type"`
			Services []RatesService `json:"services"`
			Typename string         `json:"__typename"`
		} `json:"rates"`
		Texts struct {
			TextMinPrice    string `json:"text_min_price"`
			TextDestination string `json:"text_destination"`
			TextEta         string `json:"text_eta"`
			Typename        string `json:"__typename"`
		} `json:"texts"`
		FreeShipping struct {
			Flag          int    `json:"flag"`
			ShippingPrice string `json:"shipping_price"`
			EtaText       string `json:"eta_text"`
			ErrorCode     int    `json:"error_code"`
			IconURL       string `json:"icon_url"`
			Title         string `json:"title"`
			Typename      string `json:"__typename"`
		} `json:"free_shipping"`
		TokocabangFrom struct {
			Title    string `json:"title"`
			Content  string `json:"content"`
			IconURL  string `json:"icon_url"`
			Typename string `json:"__typename"`
		} `json:"tokocabang_from"`
		IsBlackbox bool   `json:"is_blackbox"`
		Typename   string `json:"__typename"`
	} `json:"data"`
	Typename string `json:"__typename"`
}

type RatesEstimateQueryResp []struct {
	Data struct {
		RatesEstimateV3 RatesEstimateV3 `json:"ratesEstimateV3"`
	} `json:"data"`
}
