package repo

import (
	"sync"
	"time"

	"gorm.io/gorm"
)

type UploadStatus struct {
	AkunCount   int `json:"akun_count"`
	CountUpload int `json:"count_upload"`
	LimitUpload int `json:"limit_upload"`
}

type AkunUploadStatus struct {
	LimitUpload int    `json:"limit_upload"`
	CountUpload int    `json:"count_upload"`
	Active      bool   `json:"active_upload"`
	Lastup      int64  `json:"lastup"`
	InUpload    bool   `json:"in_upload"`
	LastError   string `json:"last_error"`
}

type AkunItem struct {
	AkunUploadStatus
	Username     string `json:"username" gorm:"primarykey"`
	Password     string `json:"password"`
	Secret       string `json:"secret"`
	Markup       string `json:"markup"`
	Spin         string `json:"spin"`
	Collection   string `json:"collection"`
	Hastag       string `json:"hastag"`
	TitlePattern string `json:"title_pattern"`
}

type AkunRepo struct {
	DB *gorm.DB
}

func NewAkunRepo(db *gorm.DB) *AkunRepo {
	err := db.AutoMigrate(&AkunItem{})
	if err != nil {
		panic(err)
	}
	return &AkunRepo{
		DB: db,
	}
}

type AkunUploadIterator struct {
	sync.Mutex
	db *gorm.DB
}

func NewAkunUploadIterator(db *gorm.DB) *AkunUploadIterator {
	return &AkunUploadIterator{
		db: db,
	}
}

func (iter *AkunUploadIterator) Reset() error {
	return iter.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&AkunItem{}).Updates(map[string]interface{}{"InUpload": false}).Error
}

func (iter *AkunUploadIterator) GetStatus() (*UploadStatus, error) {
	var hasil UploadStatus
	query := `SELECT count(username) as akun_count, sum(limit_upload) as limit_upload, sum(count_upload) as count_upload FROM akun_items WHERE active = 1`
	err := iter.db.Raw(query).Scan(&hasil).Error
	return &hasil, err
}
func (iter *AkunUploadIterator) InProcessCount() (int64, error) {
	iter.Lock()
	defer iter.Unlock()
	var count int64
	err := iter.db.Model(&AkunItem{}).Where(&AkunItem{
		AkunUploadStatus: AkunUploadStatus{
			InUpload: true,
		},
	}).Count(&count).Error
	return count, err
}

func (iter *AkunUploadIterator) DeactiveAll() error {
	return iter.db.Model(&AkunItem{}).Updates(&AkunItem{
		AkunUploadStatus: AkunUploadStatus{
			Active: false,
		},
	}).Error
}

func (iter *AkunUploadIterator) Get() (akun *AkunItem, updateinc func(count int, err error) error, finish func() error, err error) {
	iter.Lock()
	defer iter.Unlock()

	akun = &AkunItem{}
	updateinc = func(count int, err error) error { return nil }
	finish = func() error { return nil }
	err = iter.db.Transaction(func(tx *gorm.DB) error {

		query := tx.Model(&AkunItem{}).Where(AkunItem{
			AkunUploadStatus: AkunUploadStatus{
				Active:   true,
				InUpload: false,
			},
		}, "Active", "InUpload").Order("lastup asc")

		err := query.First(akun).Error

		if err != nil {
			return err
		}

		akun.AkunUploadStatus.Lastup = time.Now().UnixNano()
		akun.AkunUploadStatus.InUpload = true

		return tx.Save(akun).Error
	})

	if err != nil {
		return akun, updateinc, finish, err
	}
	finish = func() error {
		iter.Lock()
		defer iter.Unlock()
		akun.AkunUploadStatus.Active = false
		akun.AkunUploadStatus.CountUpload = 0

		return iter.db.Save(akun).Error
	}

	updateinc = func(count int, err error) error {

		akun.AkunUploadStatus.Lastup = time.Now().UnixNano()
		akun.AkunUploadStatus.InUpload = false
		akun.AkunUploadStatus.CountUpload += 1

		if akun.AkunUploadStatus.CountUpload >= akun.LimitUpload {
			return finish()
		}

		if err != nil {
			akun.AkunUploadStatus.LastError = err.Error()
		} else {
			akun.AkunUploadStatus.LastError = ""
		}

		iter.Lock()
		defer iter.Unlock()
		return iter.db.Save(akun).Error
	}

	return akun, updateinc, finish, nil
}
