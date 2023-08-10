package model_public

import (
	"net/url"
	"strings"
)

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
	ShopDomain   string        `json:"shopDomain"`
	ProductKey   string        `json:"productKey"`
	LayoutID     string        `json:"layoutID"`
	APIVersion   int           `json:"apiVersion"`
	Tokonow      *Tokonow      `json:"tokonow,omitempty"`
	UserLocation *UserLocation `json:"userLocation,omitempty"`
	ExtParam     string        `json:"extParam"`
}

func NewPdpGetlayoutQueryVar(uri string) (queryVar *PdpGetlayoutQueryVar, err error) {
	u, err := url.Parse(uri)
	if err != nil {
		return queryVar, err
	}

	path := u.EscapedPath()
	query := u.Query()

	splitPath := strings.Split(path, "/")
	shopDomain := splitPath[len(splitPath)-2]
	productKey := splitPath[len(splitPath)-1]

	queryVar = &PdpGetlayoutQueryVar{
		ShopDomain: shopDomain,
		ProductKey: productKey,
		APIVersion: 1,
		ExtParam:   url.QueryEscape(query.Get("extParam")),
	}

	return queryVar, err
}

type PdpGetlayoutData struct {
	PdpGetLayout PdpGetLayout `json:"pdpGetLayout"`
}

type PdpGetlayoutQueryResp struct {
	Data PdpGetlayoutData `json:"data"`
}
