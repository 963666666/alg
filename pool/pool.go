package pool

import (
	"sync"
)

type pool struct {
	wg *sync.WaitGroup
	ch chan int
}

func New(count int64) *pool {
	return &pool{
		wg: new(sync.WaitGroup),
		ch: make(chan int, count),
	}
}

func (p *pool) Add() {
	p.ch <- 0
	p.wg.Add(1)
}

func (p *pool) Down() {
	<-p.ch
	p.wg.Done()
}

func (p *pool) Wait() {
	p.wg.Wait()
}
