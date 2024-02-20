package toped_flow

import (
	"context"
	"log"

	"github.com/pdcgo/tokopedia_lib/lib/repo"
)

// TODO: bakalan didelete

// deprecated
type TopedToTopedFlow struct {
	Ctx          context.Context
	limitGuard   chan int
	AkunIterator *repo.AkunUploadIterator
}

func NewTopedToTopedFlow() *TopedToTopedFlow {
	flow := TopedToTopedFlow{}

	return &flow
}

func (flow *TopedToTopedFlow) Run() error {
	log.Println("running Tokopedia To Tokopedia Upload...")
	err := flow.AkunIterator.Reset()
	if err != nil {
		return err
	}

MainLoop:
	for {
		select {
		case flow.limitGuard <- 1:
			// go flow.RunTask()
		case <-flow.Ctx.Done():
			break MainLoop
		}
	}

	return nil

}
