package helper

import "sync/atomic"

type Counter struct {
	counter int32
}

func NewCounter() *Counter {
	return &Counter{
		counter: 0,
	}
}

func (c *Counter) Add() {
	atomic.AddInt32(&c.counter, 1)
}

func (c *Counter) Count() int32 {
	return c.counter
}
