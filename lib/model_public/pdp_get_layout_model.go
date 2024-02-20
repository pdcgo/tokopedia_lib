package model_public

import (
	"errors"
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

	if len(splitPath) < 3 {
		err = errors.New("invalid url " + uri)
		return
	}

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

type PdpErrorCode int

const (
	PdpNotFoundCode PdpErrorCode = 2001
)

type PdpGetlayoutExtensions struct {
	Code             PdpErrorCode `json:"code"`
	DeveloperMessage string       `json:"developerMessage"`
	MoreInfo         string       `json:"moreInfo"`
	Timestamp        string       `json:"timestamp"`
}

type PdpGetlayoutErr struct {
	Message    string                 `json:"message"`
	Path       []string               `json:"path"`
	Extensions PdpGetlayoutExtensions `json:"extensions"`
}

type PdpGetlayoutQueryResp struct {
	Data   PdpGetlayoutData  `json:"data"`
	Errors []PdpGetlayoutErr `json:"errors"`
}

func (product *PdpGetlayoutQueryResp) IsNotFound() bool {
	if len(product.Errors) > 0 {
		for _, err := range product.Errors {
			if err.Extensions.Code == PdpNotFoundCode {
				return true
			}
		}
	}

	return false
}
