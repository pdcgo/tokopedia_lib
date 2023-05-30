package upload_app

import (
	"context"
	"log"
	"sync"
)

type UploadStatus struct {
	AkunCount   int `json:"account_count"`
	Uploaded    int `json:"uploaded"`
	NotUploaded int `json:"not_uploaded"`
}

type UploadApp struct {
	sync.Mutex
	Ctx    context.Context
	Cancel func()
	Status *UploadStatus
}

func NewUploadApp() *UploadApp {

	return &UploadApp{}
}

func (app *UploadApp) CreateNewContext() {
	app.Lock()
	defer app.Unlock()
	ctx, cancel := context.WithCancel(context.TODO())
	app.Ctx = ctx
	app.Cancel = cancel
}

func (app *UploadApp) Start() {
	app.CreateNewContext()

	log.Println("upload running")
}
