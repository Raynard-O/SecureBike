package main

import (
	"ProtectMyBike/clients"
	"ProtectMyBike/server"
	"fmt"
	"sync"
)

func main()  {
	idObject, servers := clients.InterfaceConnection()
	var cli *clients.Client

	cli = server.Create("raynard", idObject, servers)
	cliALL := servers.GetAllClients()
	cl := server.Deactivate( cli.ID , idObject, servers)
	fmt.Println(cli, cl, cliALL)
	var wg sync.WaitGroup
	wg.Add(1)
	var err error
	//threshold := make(chan bool, 1)
	go func() {
		input := [][]int{{1,15,2},{4,2,2}}
		err = server.CheckSecurity(input)
		wg.Done()
	}()
	wg.Wait()

}

