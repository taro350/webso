package main

import (
	"flag"
	"fmt"
	"net/http"
	"webso/con_data"

	// "webso/go_sys/con_config"

	// "github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "0.0.0.0:1200", "server address")
var upgrader = websocket.Upgrader{}

func handleAPI(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Print("Upgrade error occurs.", err)
		return // connection doesn't exist so exit from handling
	}

	// id, err := uuid.NewRandom()
	// if err != nil {
	// 	fmt.Print("UUID error occurs.", err)
	// }

	// c.Uuid = "ws-" + id.String()

	// fmt.Println(c.Uuid)

	// send data to clients
	// con_data.SendMsg("^Av", "client-websocket-id", c.Uuid, *c)

Loop:
	for {
		in := con_data.Msg{}
		err := c.ReadJSON(in)
		if err != nil {
			c.Close()
			break Loop
		}

		switch in.Type {
		case "test":
			fmt.Println(in.Data)
			break
		default:
			fmt.Println("unknown type")
			break
		}

	}

}

func main() {
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/ws", handleAPI)

	fmt.Println("Server running")
	http.ListenAndServe(*addr, r)

}
