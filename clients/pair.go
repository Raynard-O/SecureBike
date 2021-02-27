package clients

import (
	"sync"
)
type cliMap map[uint64]*client
type InterfaceClient interface {
	PairClient(id InterfaceID) cliMap
	GetAllClients() *map[uint64]*client
}

type Service struct {
	m sync.Mutex
	clients map[uint64]*client
}

func Init() *Service {
	return &Service{
		clients: make(map[uint64]*client),
	}
}

func InterfaceConnection() (InterfaceID,InterfaceClient) {
	idObject := InterfaceID(NewID())

	server := InterfaceClient(Init())
	return idObject,server
}


func (s *Service) PairClient(id InterfaceID) cliMap {
	s.m.Lock()
	defer s.m.Unlock()
	cli := new(client)
	cli.id = id.Next()
	//cli.name= "name"
	s.clients[cli.id] = cli
	c := make(cliMap)
	c[cli.id] = cli
	return c
}

func (s *Service) GetAllClients() *map[uint64]*client {
	return &s.clients
}
