package upload_app

import (
	"context"
	"log"
	"sync"

	"gorm.io/gorm"
)

type UploadConfig struct {
	Concurent int
}

type UploadApp struct {
	sync.Mutex
	limitGuard chan int
	iterator   *AkunUploadIterator
	Ctx        context.Context
	Cancel     func()
}

func NewUploadApp(db *gorm.DB, config *UploadConfig) *UploadApp {
	iterator := NewAkunUploadIterator(db)
	return &UploadApp{
		limitGuard: make(chan int, config.Concurent),
		iterator:   iterator,
		Cancel:     func() {},
	}
}

func (app *UploadApp) CreateNewContext() {
	app.Lock()
	defer app.Unlock()
	ctx, cancel := context.WithCancel(context.TODO())
	app.Ctx = ctx
	app.Cancel = cancel
}

func (app *UploadApp) Status() (*UploadStatus, error) {
	return app.iterator.GetStatus()

}

func (app *UploadApp) RunUpload() {

	for {

	}

}

func (app *UploadApp) Start() {
	app.Cancel()
	app.CreateNewContext()
	app.iterator.Reset()
	log.Println("upload running")

	go app.RunUpload()

}
