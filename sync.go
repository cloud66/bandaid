package bandaid

import "sync/atomic"

// AtomBool is a thread safe atomic boolean
type AtomBool struct{ flag int32 }

// Set the value for the bool
func (b *AtomBool) Set(value bool) {
	var i int32
	if value {
		i = 1
	}
	atomic.StoreInt32(&(b.flag), int32(i))
}

// Get the value of the bool
func (b *AtomBool) Get() bool {
	if atomic.LoadInt32(&(b.flag)) != 0 {
		return true
	}
	return false
}
