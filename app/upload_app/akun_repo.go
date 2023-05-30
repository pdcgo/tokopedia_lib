package upload_app

import "gorm.io/gorm"

type AkunUploadStatus struct {
	LimitUpload int     `json:"limit_upload"`
	CountUpload int     `json:"count_upload"`
	Active      bool    `json:"active_upload"`
	Lastup      float32 `json:"lastup"`
}

type AkunItem struct {
	AkunUploadStatus
	Username   string `json:"username" gorm:"primarykey"`
	Password   string `json:"password"`
	Secret     string `json:"secret"`
	Markup     string `json:"markup"`
	Spin       string `json:"spin"`
	Collection string `json:"collection"`
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
