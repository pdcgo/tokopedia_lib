package repo

import (
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"gorm.io/gorm"
)

type GroupRepo struct {
	db *gorm.DB
}

func NewGroupRepo(db *gorm.DB) *GroupRepo {
	return &GroupRepo{
		db: db,
	}
}

func (repo *GroupRepo) GetList() ([]string, error) {
	groups := []model.Group{}
	tx := repo.db.Model(model.Group{}).Find(&groups)

	if tx.Error != nil {
		return []string{}, tx.Error
	}

	list := []string{}
	for _, group := range groups {
		list = append(list, group.Name)
	}

	return list, nil
}

func (repo *GroupRepo) Delete(groupName string) error {

	group := model.Group{
		Name: groupName,
	}
	tx := repo.db.Where(&group).First(&group)
	if tx.Error != nil {
		return tx.Error
	}

	tx = repo.db.Delete(&group)
	return tx.Error
}
