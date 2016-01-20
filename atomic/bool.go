//@auther foolbread
//@time 2015-08-13
//@file bool.go
package atomic

import (
	"sync/atomic"
)

type AtomicBool struct {
	v int32
}

func (a *AtomicBool) Get() bool {
	return atomic.LoadInt32(&a.v) != 0
}

func (a *AtomicBool) Set(b bool) {
	if b {
		atomic.StoreInt32(&a.v, 1)
	} else {
		atomic.StoreInt32(&a.v, 0)
	}
}
