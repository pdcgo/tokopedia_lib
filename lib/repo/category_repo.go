package repo

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/api"
)

type BaseInterface interface {
	Path(item ...string) string
}

type CategoryRepo struct {
	fname string
	base  BaseInterface
	Data  *api.CategoryAllListLiteRes
}

func NewCategoryRepo(base BaseInterface) *CategoryRepo {
	repo := CategoryRepo{
		fname: "data/tokopedia_categories.json",
		base:  base,
	}

	repo.load()

	return &repo
}

func (repo *CategoryRepo) load() {
	fname := repo.base.Path(repo.fname)
	if _, err := os.Stat(fname); errors.Is(err, os.ErrNotExist) {
		return
	}

	data, err := os.ReadFile(fname)
	if err != nil {
		pdc_common.ReportError(err)
		return
	}

	var hasil api.CategoryAllListLiteRes
	err = json.Unmarshal(data, &hasil)
	if err != nil {
		pdc_common.ReportError(err)
		return
	}

	repo.Data = &hasil
}

func (repo *CategoryRepo) Get() *api.CategoryAllListLiteRes {
	return repo.Data

}

type UpdateTopedCategoryPayload struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Secret   string `json:"secret" form:"secret"`
}

func (repo *CategoryRepo) Save() error {
	fname := repo.base.Path(repo.fname)
	f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(repo.Data)
	return err

}

func (repo *CategoryRepo) UpdateCateg(payload *UpdateTopedCategoryPayload) error {

	driver, _ := tokopedia_lib.NewDriverAccount(payload.Username, payload.Password, payload.Secret)
	sellerApi, saveSession, _ := driver.CreateApi()
	defer saveSession()

	data, err := sellerApi.CategoryAllListLite()
	if err != nil {
		return err
	}

	repo.Data = data
	return repo.Save()

}
