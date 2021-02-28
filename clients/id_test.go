package clients

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)


func TestID_Next(t *testing.T) {
	idObject := InterfaceID(NewID())
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			id := idObject.Next()
			assert.IsType(t, uint64(5), id)
			//fmt.Println(id)
			wg.Done()
		}()
	}
	wg.Wait()
	assert.NotNil(t, idObject)
}

func TestID_CurrentID(t *testing.T) {
	idObject := InterfaceID(NewID())
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			_= idObject.Next()
			id := idObject.CurrentID()
			assert.IsType(t, uint64(5), id)
			fmt.Printf("[%v] value of ID : %v \n is ",i, id)
			wg.Done()
		}()
	}
	wg.Wait()
	assert.NotNil(t, idObject)

}