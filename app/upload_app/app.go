package upload_app

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/pdcgo/tokopedia_lib/lib/uploader"
	"gorm.io/gorm"
)

type UploadConfig struct {
	Concurent int
}

type RunStatus string

const (
	RUNNING RunStatus = "running"
	STOP    RunStatus = "stop"
)

type UploadAppStatus struct {
	Status RunStatus
	UploadStatus
}

type UploadApp struct {
	sync.Mutex
	limitGuard chan int
	iterator   *AkunUploadIterator
	Ctx        context.Context
	Cancel     func()
	RunStatus  RunStatus
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

func (app *UploadApp) Status() (*UploadAppStatus, error) {
	stat, err := app.iterator.GetStatus()
	status := UploadAppStatus{
		Status:       app.RunStatus,
		UploadStatus: *stat,
	}
	return &status, err

}

func (app *UploadApp) RunTask() {
	defer func() {
		<-app.limitGuard
	}()

	_, updateinc, _, err := app.iterator.Get()
	if err != nil {
		pdc_common.ReportError(err)
		return
	}

	ctx, cancel := context.WithTimeout(app.Ctx, time.Minute*3)
	defer cancel()

	uploaderItem := uploader.NewTokopediaUploader(ctx, &api.TokopediaApi{})

	_, err = uploaderItem.RunUploader(func(eventcore uploader.EmitFunc, tokpedup *uploader.TokopediaUploader, payload *uploader.PayloadUpload, sub *common_concept.Subscriber) error {
		return nil
	})

	if err != nil {
		updateinc(0, err)
		pdc_common.ReportError(err)
		return
	} else {
		updateinc(1, err)
	}

}

func (app *UploadApp) RunUpload() {
	app.RunStatus = RUNNING
	defer func() {
		app.RunStatus = STOP
	}()

MainLoop:
	for {
		select {
		case app.limitGuard <- 1:
			go app.RunTask()
		case <-app.Ctx.Done():
			break MainLoop
		}
	}

}

func (app *UploadApp) Start() {
	app.Cancel()
	app.CreateNewContext()
	app.iterator.Reset()
	log.Println("upload running")

	go app.RunUpload()

}
