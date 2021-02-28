package server

import (
	"sync"
	"testing"
)

func TestNotifyU(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	threshold := make(chan bool, 1)
	go func() {
		NotifyU(threshold, &wg)
		wg.Done()
	}()
	wg.Wait()
	//assert.Implements(t, NotifyU, Notify)
	//<- threshold

}
