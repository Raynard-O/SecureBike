package clients

import (
	"sync"
)
type cliMap map[uint64]*Client
type InterfaceClient interface {
	PairClient(id InterfaceID) *Client
	GetAllClients() *map[uint64]*Client
	GetAClient (id uint64) *Client
}

type Service struct {
	m sync.Mutex
	clients map[uint64]*Client
}

func Init() *Service {
	return &Service{
		clients: make(map[uint64]*Client),
	}
}

func InterfaceConnection() (InterfaceID,InterfaceClient) {
	idObject := InterfaceID(NewID())

	server := InterfaceClient(Init())
	return idObject,server
}


func (s *Service) PairClient(id InterfaceID) *Client {
	s.m.Lock()
	defer s.m.Unlock()
	cli := new(Client)
	cli.ID = id.Next()
	//cli.name= "name"
	s.clients[cli.ID] = cli
	//c := make(cliMap)
	//c[cli.ID] = cli
	return s.clients[cli.ID]
}

func (s *Service) GetAllClients() *map[uint64]*Client {
	return &s.clients
}

func (s *Service) GetAClient (id uint64) *Client{
	return s.clients[id]
}

//func (s *Service) UpdateClient(ID uint64, key string, value interface{})  {
//	//m := make(map[uint64]*client)
//	//cli := new(client)
//
//	j := s.clients[ID]
//	j.
//	client[key] = value
//}