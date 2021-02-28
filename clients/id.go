package clients

import "sync"

type InterfaceID interface {
	Next() uint64
	CurrentID() uint64
}

func NewID() *ID {
	return &ID{}
}
//Id class
type ID struct {
	id uint64
	lock sync.Mutex
}
//Next add microcontroller to server
//the process is locked for memory lock
func (i *ID) Next() uint64 {
	i.lock.Lock()
	//var ID int8
	i.id++

	defer i.lock.Unlock()
	return i.id
}
func (i *ID) CurrentID() uint64 {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.id
}
