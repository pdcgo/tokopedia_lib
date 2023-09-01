package model

type Group struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"-"`
	Name string `gorm:"unique" json:"name"`

	AccountDatas []AccountData `gorm:"many2many:account_groups" json:"-"`
}
