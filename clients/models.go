package clients

import (
	"github.com/gorilla/websocket"
	"time"
)

type Client struct {
	ID   uint64
	Name string
	Secure bool
	Duration time.Time
	Conn *websocket.Conn
}
