package clients

import (
	"fmt"
	"time"
)

func CreateClient(name string) {
	idObject, server := InterfaceConnection()
	//var wg sync.WaitGroup
	////var cli *clients.Client
	//go func() {
	//	wg.Add(1)
		cli := server.PairClient(idObject)
		fmt.Println(cli)
		cli.Name = name
		cli.Secure = true
		cli.Duration = time.Now().UTC()

	clie := server.GetAClient(cli.ID)
	fmt.Println(clie)
	//wg.Wait()
}

