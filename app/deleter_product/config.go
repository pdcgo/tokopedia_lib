package deleter_product

import (
	"encoding/json"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/pdcgo/tokopedia_lib/lib/model"
)

type AkunDeleteItem struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Secret   string `json:"secret"`
}

type DeleteConfig struct {
	LimitConcurent int                 `json:"limit_concurent"`
	LimitProduct   int                 `json:"limit_product"`
	Title          []string            `json:"title"`
	StatusProduct  model.ProductStatus `json:"product_status"`
	CategoryID     string              `json:"category_id"`
	StartTime      int64               `json:"start_time"`
	EndTime        int64               `json:"end_time"`
	TStartTime     time.Time           `json:"-"`
	TEndTime       time.Time           `json:"-"`
	Akuns          []*AkunDeleteItem   `json:"akuns"`
}

func NewDeleteConfig(fname string) (*DeleteConfig, error) {
	var cfg DeleteConfig

	data, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &cfg)

	return &cfg, err
}

func SaveDeleteConfig(fname string, cfg *DeleteConfig) error {
	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(cfg)
	if err != nil {
		return err
	}
	return nil
}

func (cfg *DeleteConfig) UnmarshalJSON(data []byte) error {
	type Alias DeleteConfig
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(cfg),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	cfg.TStartTime = time.Unix(aux.StartTime, 0)
	cfg.TEndTime = time.Unix(aux.EndTime, 0)

	return nil
}

type FilterHandler func(product *model.SellerProductItem) bool

func (cfg *DeleteConfig) GenerateFilter() FilterHandler {

	handlers := []FilterHandler{
		cfg.GenerateFilterTime(),
		cfg.GenerateFilterTitle(),
	}

	return func(product *model.SellerProductItem) bool {
		for _, handler := range handlers {
			if !handler(product) {

				return false
			}
		}

		return true
	}
}

func (cfg *DeleteConfig) GenerateFilterTime() FilterHandler {
	return func(product *model.SellerProductItem) bool {
		data := product.CreateTime

		return data.After(cfg.TStartTime) && data.Before(cfg.TEndTime)
	}
}

func (cfg *DeleteConfig) GenerateFilterTitle() FilterHandler {
	fstring := []string{}
	fregex := []*regexp.Regexp{}

	regpola := `regex-->`

	for _, title := range cfg.Title {
		if strings.HasPrefix(title, regpola) {
			title = strings.Replace(title, regpola, "", 1)
			reitem := regexp.MustCompile(title)
			fregex = append(fregex, reitem)
		} else {
			fstring = append(fstring, title)
		}
	}

	return func(product *model.SellerProductItem) bool {
		datatitle := product.Name
		datatitle = strings.ToLower(datatitle)

		for _, fstr := range fstring {
			if strings.Contains(datatitle, fstr) {
				return true
			}
		}

		for _, freg := range fregex {
			if freg.MatchString(datatitle) {
				return true
			}
		}
		return false
	}
}
