package ws

import (
	"net/http"
	"log"
	"github.com/gorilla/websocket"
	"fmt"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:1024,
	WriteBufferSize:1024,
}


func Echo (w http.ResponseWriter, r * http.Request){
	conn, err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	for {
		messageType , p ,err := conn.ReadMessage()
		if err != nil {
			return
		}
		print_binary(p)
		err = conn.WriteMessage(messageType, p)
		if err != nil {
			return
		}
	}
}

func print_binary(s []byte) {
	fmt.Printf("Received b:");
	for n := 0;n < len(s);n++ {
		fmt.Printf("%d,",s[n]);
	}
	fmt.Printf("\n");
}
