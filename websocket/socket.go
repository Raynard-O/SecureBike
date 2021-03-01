package socket

import (
	"ProtectMyBike/clients"
	"ProtectMyBike/server"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}


func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "websockets.html")
}

func SocketWeb(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
	log.Println("Client Connected")
	err := conn.WriteMessage(1, []byte("Hi Client!, Enter Microcontroller ID"))
	if err != nil {
		log.Println(err)
	}


	// Read message from browser
	msgType, msg, err := conn.ReadMessage()
	if err != nil {
		return
	}

	// Print the message to the console
	m:= fmt.Sprintf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

	idObject, servers := clients.InterfaceConnection()
	var cli *clients.Client

	cli = server.Create(conn,m, idObject, servers)
	fmt.Println(cli)
	//Write message back to browser
	if err = conn.WriteMessage(msgType, msg); err != nil {
		return
	}

	err = conn.WriteMessage(1, []byte("Enter A Request"))
	if err != nil {
		log.Println(err)
	}

	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		// Print the message to the console
		//m:= fmt.Sprintf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		//idObject, servers := clients.InterfaceConnection()
		//var cli *clients.Client
		//
		//cli = server.Create(conn,m, idObject, servers)
		//fmt.Println(cli)


		mm := string(msg)
		switch mm {
		case "identify":
			//b := []byte(strconv.FormatInt(int64(cli.ID), 10))
			b := fmt.Sprintf("client ID: %v ", cli.ID)
			if err = conn.WriteMessage(msgType,[]byte(b)); err != nil {
				return
			}
		case "disconnect":
			cl := server.Deactivate( cli.ID , idObject, servers)
			b := fmt.Sprintf("client %v is Diconnected", cl.ID)
			if err = conn.WriteMessage(msgType,[]byte(b)); err != nil {
				return
			}
		default:
			var wg sync.WaitGroup
			wg.Add(1)
			var err error
			threshold := make(chan string)
			go func() {
				input := [][]int{{1,15,2},{4,2,2}}
				err = server.CheckSecurity(input,threshold)
				wg.Done()
			}()
			j := <-threshold
			if err = conn.WriteMessage(msgType,[]byte(j)); err != nil {
				return
			}
			//fmt.Println(j)
			wg.Wait()
		}

	}
}