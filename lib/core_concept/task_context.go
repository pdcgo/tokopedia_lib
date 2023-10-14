package core_concept

import (
	"context"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/rs/zerolog"
)

type TaskContext struct {
	Ctx        context.Context
	Err        chan error
	CancelTask func()
}

func (ctx *TaskContext) GetCtx() context.Context {
	return ctx.Ctx
}

func (ctx *TaskContext) Cancel() {
	ctx.CancelTask()
}

func (ctx *TaskContext) SetErrorCustom(err error, handler func(event *zerolog.Event) *zerolog.Event) {
	pdc_common.ReportErrorCustom(err, handler)
	ctx.Err <- err
}

func (ctx *TaskContext) SetError(err error) {
	pdc_common.ReportError(err)
	ctx.Err <- err
}

func NewTaskContext(ctx context.Context) *TaskContext {
	ctx, cancel := context.WithCancel(ctx)
	errchan := make(chan error, 10)
	return &TaskContext{
		Ctx: ctx,
		Err: errchan,
		CancelTask: func() {
			cancel()
			close(errchan)
		},
	}
}
