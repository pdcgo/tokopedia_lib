package repo

import (
	"errors"
	"fmt"

	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"gorm.io/gorm"
)

type AccountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) *AccountRepo {
	return &AccountRepo{
		db: db,
	}
}

func (repo *AccountRepo) GetChatAccountData(groupName string, shopid int) (*model.AccountData, error) {

	var account model.AccountData
	tx := repo.db.Preload("Groups", "name = ?", groupName).
		Where("shop_id = ?", shopid).
		First(&account)

	if len(account.Groups) == 0 {
		msg := fmt.Sprintf("%d tidak ditemukan pada group %s", shopid, groupName)
		return &account, errors.New(msg)
	}

	return &account, tx.Error
}

func (repo *AccountRepo) IterateGroupAccount(groupName string, handler func(account model.AccountData) error) error {

	group := model.Group{
		Name: groupName,
	}
	tx := repo.db.Preload("AccountDatas").Where(group).First(&group)

	if tx.Error != nil {
		return tx.Error
	}

	for _, accountData := range group.AccountDatas {
		err := handler(accountData)
		if err != nil {
			return err
		}
	}

	return nil
}

func (repo *AccountRepo) AddAccountData(groupName string, accountData model.AccountData) error {

	group := model.Group{
		Name: groupName,
	}
	tx := repo.db.Where(&group).FirstOrCreate(&group)
	if tx.Error != nil {
		return tx.Error
	}

	tx = repo.db.Preload("Groups").Where(&model.AccountData{
		Username: accountData.Username,
	}).FirstOrCreate(&accountData)
	if tx.Error != nil {
		return tx.Error
	}

	accountData.AppendGroup(group)
	tx = repo.db.Save(&accountData)
	return tx.Error
}

type ListAccountFilter struct {
	GroupName string `form:"group_name" schema:"group_name" json:"group_name"`
	Name      string `form:"name" schema:"name" json:"name"`
	HaveChat  bool   `form:"have_chat" schema:"have_chat" json:"have_chat"`
	Online    bool   `form:"online" schema:"online" json:"online"`
	Pinned    bool   `form:"pinned" schema:"pinned" json:"pinned"`
	Saldo     bool   `form:"saldo" schema:"saldo" json:"saldo"`
}

func (repo *AccountRepo) List(filter *ListAccountFilter) ([]*model.Account, error) {

	accounts := []*model.Account{}
	tx := repo.db.Preload("AccountData.Groups").Joins("AccountData")

	if filter.GroupName != "" {
		tx = tx.Where(`EXISTS (
			SELECT 1 FROM "groups", account_groups
			WHERE AccountData__id = account_groups.account_data_id
			AND "groups".id = account_groups.group_id
			AND "groups".name = ?
		)`, filter.GroupName)
	}

	if filter.Name != "" {
		tx = tx.Where("shop_name LIKE %?%", filter.Name)
	}

	if filter.HaveChat {
		tx = tx.Where("unread_chat > 0")
	}

	if filter.Saldo {
		tx = tx.Where("saldo > 0")
	}

	if filter.Pinned {
		tx = tx.Where("AccountData__pinned = true")
	}

	tx = tx.Find(&accounts)

	if tx.Error != nil {
		return accounts, tx.Error
	}

	return accounts, nil
}
