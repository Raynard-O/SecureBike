package server

import (
	"ProtectMyBike/clients"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateClient(t *testing.T) {
		cli := CreateClient(true, 0)
		assert.NotNil(t, cli)
		fmt.Println(cli)
}


func TestDeactivate(t *testing.T) {
	idObject, server := clients.InterfaceConnection()
	var cli *clients.Client

	cli = Create("raynard", idObject, server)

	cl := Deactivate( cli.ID , idObject, server)

	fmt.Println(cl)
	assert.NotNil(t, cli)

}
