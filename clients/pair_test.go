package clients

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)


func TestService_PairClient(t *testing.T) {
	idObject, server := InterfaceConnection()
	var wg sync.WaitGroup
	typeMap := cliMap{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			cli := server.PairClient(idObject)
			fmt.Printf("[%v] value of id : %v \n is ", i,cli)
			assert.IsType(t, typeMap, cli)
			wg.Done()
		}()
	}
	wg.Wait()
	assert.NotNil(t, server, idObject)

}

func TestService_GetAllClients(t *testing.T) {
	idObject, server := InterfaceConnection()
	var wg sync.WaitGroup
	mapSize := 10
	typeMap := cliMap{}
	for i := 0; i < mapSize; i++ {
		wg.Add(1)
		go func() {
			cli := server.PairClient(idObject)

			assert.IsType(t, typeMap, cli)
			wg.Done()
		}()
	}
	wg.Wait()
	clients := server.GetAllClients()
	assert.NotNil(t, server, idObject)
	for k,v := range *clients{
		assert.NotNil(t, v, k)
		fmt.Printf("value of id {%v} is %v \n  ",k,v)
		if k > uint64(mapSize) {

			t.Error("error in map size")
		}
	}

}