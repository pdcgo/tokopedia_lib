package model

type IncomeDetailSection string
type IncomeDetailKey string

var TOTAL_SALES_SECTION IncomeDetailSection = "total_sales"

var (
	TOTAL_PRICE_KEY           IncomeDetailKey = "TOTAL_PRICE"
	TOTAL_SHIPPING_FEE_KEY    IncomeDetailKey = "TOTAL_SHIPPING_FEE"
	CASHLESS_SHIPPING_FEE_KEY IncomeDetailKey = "CASHLESS_SHIPPING_FEE"
)

type SOMComponent struct {
	Key        IncomeDetailKey `json:"key"`
	Label      string          `json:"label"`
	SubLabel   any             `json:"sub_label"`
	Value      string          `json:"value"`
	ValueRaw   int             `json:"value_raw"`
	Attributes any             `json:"attributes"`
	Type       string          `json:"type"`
	Typename   string          `json:"__typename"`
}

type SomSection struct {
	Key        IncomeDetailSection `json:"key"`
	Label      string              `json:"label"`
	SubLabel   any                 `json:"sub_label"`
	Value      string              `json:"value"`
	ValueRaw   int                 `json:"value_raw"`
	Attributes any                 `json:"attributes"`
	Components []*SOMComponent     `json:"components"`
	Typename   string              `json:"__typename"`
}

type SOMSummary struct {
	Key        string `json:"key"`
	Label      string `json:"label"`
	SubLabel   any    `json:"sub_label"`
	Value      string `json:"value"`
	ValueRaw   int    `json:"value_raw"`
	Attributes any    `json:"attributes"`
	State      string `json:"state"`
	Note       string `json:"note"`
	Typename   string `json:"__typename"`
}

type GetSomIncomeDetail struct {
	Title    string        `json:"title"`
	Sections []*SomSection `json:"sections"`
	Summary  *SOMSummary   `json:"summary"`
	Typename string        `json:"__typename"`
}

type SOMIncomeDetailRes struct {
	Data struct {
		GetSomIncomeDetail GetSomIncomeDetail `json:"get_som_income_detail"`
	} `json:"data"`
}

func (d *SOMIncomeDetailRes) GetSection(key IncomeDetailSection) *SomSection {
	for _, section := range d.Data.GetSomIncomeDetail.Sections {
		if section.Key == key {
			return section
		}
	}
	return nil
}

func (d *SOMIncomeDetailRes) GetSectionComponent(sec IncomeDetailSection, key IncomeDetailKey) *SOMComponent {
	section := d.GetSection(sec)
	if section != nil {
		for _, component := range section.Components {
			if component.Key == key {
				return component
			}
		}
	}

	return nil
}

func (d *SOMIncomeDetailRes) GetOngkir() int {
	component := d.GetSectionComponent(TOTAL_SALES_SECTION, TOTAL_SHIPPING_FEE_KEY)
	if component != nil {
		return component.ValueRaw
	}

	return 0
}
