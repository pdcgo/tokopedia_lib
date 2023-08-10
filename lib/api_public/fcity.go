package api_public

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/chromedp/chromedp"
	"github.com/pdcgo/tokopedia_lib"
)

type generalType map[string]map[string]any

type Fcity struct {
	Description string `json:"Description"`
	Typename    string `json:"__typename"`
	Child       []any  `json:"child"`
	HexColor    string `json:"hexColor"`
	Icon        string `json:"icon"`
	InputType   string `json:"inputType"`
	IsNew       bool   `json:"isNew"`
	IsPopular   bool   `json:"isPopular"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	TotalData   string `json:"totalData"`
	ValMax      string `json:"valMax"`
	ValMin      string `json:"valMin"`
	Value       string `json:"value"`
}

func GetFcity(fname string) ([]*Fcity, error) {
	cities := []*Fcity{}
	file, err := GetDataAllFilterConfig(fname)
	if err != nil {
		return cities, err
	}
	defer file.Close()

	hasil := generalType{}

	err = json.NewDecoder(file).Decode(&hasil)

	for key, value := range hasil {
		if !strings.Contains(key, "filter_sort_product") {
			continue
		}

		if value["key"] != "fcity" {
			continue
		}

		data, err := json.Marshal(&value)
		if err != nil {
			return cities, err
		}

		var item Fcity
		err = json.Unmarshal(data, &item)
		if err != nil {
			return cities, err
		}

		cities = append(cities, &item)
	}

	return cities, err
}

func GetDataAllFilterConfig(fname string) (*os.File, error) {

	if _, err := os.Stat(fname); errors.Is(err, os.ErrNotExist) {
		driver, err := tokopedia_lib.NewDriverAccount("", "", "")
		if err != nil {
			return nil, err
		}

		data := ""

		err = driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {
			return chromedp.Run(
				dctx.Ctx,
				chromedp.Navigate("https://www.tokopedia.com/search?q=gamis"),
				chromedp.Evaluate("JSON.stringify(window.__cache)", &data),
			)
		})

		if err != nil {
			return nil, err
		}

		file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return nil, err
		}
		file.Write([]byte(data))
		file.Close()
	}

	return os.Open(fname)

}
