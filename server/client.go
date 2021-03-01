package server

import (
	"ProtectMyBike/clients"
	"github.com/gorilla/websocket"
	"time"
)

type InterfaceClient interface {
	Create(name string, idObject clients.InterfaceID, server clients.InterfaceClient) *clients.Client
	Deactivate(id uint64, idObject clients.InterfaceID, server clients.InterfaceClient) *clients.Client
	GetAll(id uint64, idObject clients.InterfaceID, server clients.InterfaceClient)  *map[uint64]*clients.Client
}


func CreateClient(conn *websocket.Conn, action bool, id uint64) *clients.Client {
	idObject, server := clients.InterfaceConnection()
	var cli *clients.Client
	if action{
		cli = Create(conn,"raynard", idObject, server)
	}
	//fmt.Println(server.GetAClient(cli.ID))
	return cli
}


func Create(conn *websocket.Conn, name string, idObject clients.InterfaceID, server clients.InterfaceClient) *clients.Client{
	//idObject, server := clients.InterfaceConnection()
	cli := server.PairClient(idObject)
	//fmt.Println(cli)
	cli.Name = name
	cli.Secure = true
	cli.Conn = conn
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


func GetAll(id uint64, idObject clients.InterfaceID, server clients.InterfaceClient)  *map[uint64]*clients.Client {
	clie := server.GetAllClients()

	return clie
}
