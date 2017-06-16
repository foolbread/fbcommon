/*
author: foolbread
file: pool/workpool.go
date: 2017/6/16
*/
/*
	a easy goroutine pool
*/
package pool

import (
	"sync"
)

type WorkPool struct {
	ch chan interface{}
	worker func(interface{})
	count uint
	wait sync.WaitGroup
}

func NewWorkPool(work func(interface{}),c uint)*WorkPool{
	r := new(WorkPool)
	r.ch = make(chan interface{},100)
	r.worker = work
	r.count = c

	return r
}

func (w *WorkPool)CommitWork(data interface{}){
	w.ch <- data
}

func (w *WorkPool)Close(){
	close(w.ch)
}

func (w *WorkPool)WaitClose(){
	close(w.ch)
	w.wait.Wait()
}

func (w *WorkPool)Run(){
	for i := 0; i < int(w.count); i++ {
		w.wait.Add(1)
		go w.work()
	}
}

func (w *WorkPool)work(){
	for{
		data,ok := <-w.ch
		if !ok{
			w.wait.Done()
			return
		}

		w.worker(data)
	}
}