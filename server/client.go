package server

import (
	"ProtectMyBike/clients"
	"time"
)

type InterfaceClient interface {
	Create(name string, idObject clients.InterfaceID, server clients.InterfaceClient) *clients.Client
	Deactivate(id uint64, idObject clients.InterfaceID, server clients.InterfaceClient) *clients.Client
}


func CreateClient(action bool, id uint64) *clients.Client {
	idObject, server := clients.InterfaceConnection()
	var cli *clients.Client
	if action{
		cli = Create("raynard", idObject, server)
	}
	Deactivate( cli.ID , idObject, server)

	//fmt.Println(server.GetAClient(cli.ID))
	return cli
}


func Create(name string, idObject clients.InterfaceID, server clients.InterfaceClient) *clients.Client{
	//idObject, server := clients.InterfaceConnection()
	cli := server.PairClient(idObject)
	//fmt.Println(cli)
	cli.Name = name
	cli.Secure = true
	cli.Duration = time.Now().UTC()
	clie := server.GetAClient(cli.ID)
	//fmt.Println(clie)
	return clie
}



//Deactivate
// takes in the interfaces
// deactivate an id
func Deactivate(id uint64, idObject clients.InterfaceID, server clients.InterfaceClient) *clients.Client {
	clie := server.GetAClient(id)
	clie.Secure = false
	clie2 := server.GetAClient(id)
	return clie2
}

