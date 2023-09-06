package model

type Account struct {
	ID         int    `gorm:"primaryKey;autoIncrement:false" json:"id"`
	ProfileUrl string `json:"profile_url"`
	ShopName   string `json:"shopname"`
	ShopDomain string `json:"-"`
	UnreadChat int    `json:"unread_chat"`
	Online     bool   `json:"online"`
	NewOrder   int    `json:"-"`
	Saldo      int    `json:"saldo"`
	Diskusi    int    `json:"diskusi"`
	NotifHash  string `json:"-"`

	AccountData *AccountData `gorm:"foreignKey:ShopID" json:"akun_data"`
	Orders      []Order      `gorm:"foreignKey:ShopID" json:"-"`
}

func (a *Account) GetUsername() string {
	if a.AccountData == nil {
		return a.ShopName
	}

	return a.AccountData.Username
}

type AccountData struct {
	ID          int    `gorm:"primaryKey;autoIncrement:true" json:"-"`
	Username    string `gorm:"unique" json:"username"`
	Password    string `json:"password"`
	OtpPassword string `grom:"constraint:OnDelete:SET NULL" json:"otp_password"`
	ShopID      int    `json:"-"`
	Pinned      bool   `json:"pinned"`
	Deleted     bool   `json:"-"`

	Account Account `gorm:"foreignKey:ShopID" json:"-"`
	Groups  []Group `gorm:"many2many:account_groups" json:"groups"`
}

func (ad *AccountData) AppendGroup(group Group) {
	groupExist := false

	for _, g := range ad.Groups {
		if g.Name == group.Name {
			groupExist = true
			break
		}
	}

	if !groupExist {
		ad.Groups = append(ad.Groups, group)
	}
}

type AccountGroups struct {
	GroupId       string `json:"group_id"`
	AccountDataId string `json:"account_data_id"`
}
