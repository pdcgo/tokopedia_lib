package api_public

import (
	"encoding/json"
	"strings"
)

type PubShippingItem struct {
	Name        string `json:"name"`
	Description string `json:"Description"`
	Key         string `json:"key"`
	Icon        string `json:"icon"`
	Value       string `json:"value"`
	InputType   string `json:"inputType"`
	TotalData   string `json:"totalData"`
	ValMax      string `json:"valMax"`
	ValMin      string `json:"valMin"`
	HexColor    string `json:"hexColor"`
	Child       []any  `json:"child"`
	IsPopular   bool   `json:"isPopular"`
	IsNew       bool   `json:"isNew"`
	Typename    string `json:"__typename"`
}

func GetPubShippings(fname string) ([]*PubShippingItem, error) {
	shippings := []*PubShippingItem{}
	file, err := GetDataAllFilterConfig(fname)
	if err != nil {
		return shippings, err
	}
	defer file.Close()

	hasil := generalType{}

	err = json.NewDecoder(file).Decode(&hasil)

	for key, value := range hasil {
		if !strings.Contains(key, "filter_sort_product") {
			continue
		}

		if value["key"] != "shipping" {
			continue
		}

		data, err := json.Marshal(&value)
		if err != nil {
			return shippings, err
		}

		var item PubShippingItem
		err = json.Unmarshal(data, &item)
		if err != nil {
			return shippings, err
		}

		shippings = append(shippings, &item)
	}

	return shippings, err
}
