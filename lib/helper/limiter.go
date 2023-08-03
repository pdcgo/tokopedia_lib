package helper

import "sync/atomic"

type Limiter struct {
	limit int32
	count int32
}

func NewLimiter(limit int32) *Limiter {
	return &Limiter{
		limit: limit,
	}
}

func (l *Limiter) Add() {
	atomic.AddInt32(&l.count, 1)
}

func (l *Limiter) LimitReached() bool {
	return l.count >= l.limit
}

func (l *Limiter) ResetLimiter() {
	l.count = 0
}

func (l *Limiter) Count() int32 {
	return l.count
}
