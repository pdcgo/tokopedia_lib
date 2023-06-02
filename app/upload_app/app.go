package upload_app

import (
	"context"
	"sync"

	"github.com/pdcgo/tokopedia_lib/app/upload_app/shopee_flow"
	"github.com/pdcgo/tokopedia_lib/lib/repo"
	"gorm.io/gorm"
)

type RunStatus string

const (
	RUNNING RunStatus = "running"
	STOP    RunStatus = "stop"
)

type UploadAppStatus struct {
	Status RunStatus `json:"status"`
	repo.UploadStatus
}

type UploadApp struct {
	sync.Mutex
	Ctx       context.Context
	Cancel    func()
	RunStatus RunStatus
	flow      *shopee_flow.ShopeeToTopedFlow
}

func NewUploadApp(db *gorm.DB) *UploadApp {
	return &UploadApp{
		Cancel: func() {},
	}
}

func (app *UploadApp) CreateNewContext() {
	app.Lock()
	defer app.Unlock()
	ctx, cancel := context.WithCancel(context.TODO())
	app.Ctx = ctx
	app.Cancel = cancel
}

func (app *UploadApp) Status() (*UploadAppStatus, error) {
	stat, err := app.flow.AkunIterator.GetStatus()
	status := UploadAppStatus{
		Status:       app.RunStatus,
		UploadStatus: *stat,
	}
	return &status, err

}

func (app *UploadApp) Start() {
	app.Cancel()
	app.CreateNewContext()

	app.RunStatus = RUNNING
	defer func() {
		app.RunStatus = STOP
	}()

}
