package handlers


import (
	"net/http"
	"fmt"

	"github.com/gorilla/websocket"
	log "github.com/Sirupsen/logrus"
	"time"
)

type WebsocketHandler struct {
	commands chan bool
}

func (handler WebsocketHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	up := websocket.Upgrader{WriteBufferSize:1024}
	conn, err := up.Upgrade(response, request, nil)	// This closes the response writer
	if err != nil {
		log.Error(fmt.Sprintf("Error establishing websocket connection : %v", err))
		return
	}


	// Write pump
	go func() {
		for {
			mtype, message, err := conn.ReadMessage()
			if mtype == websocket.CloseMessage {
				fmt.Printf("closed !")
				return
			}
			if err != nil {
				fmt.Printf("READ ERROR : %v\n", err)
				return
			} else {
				fmt.Printf("Message received. Type : %v ; payload : %v\n", mtype, string(message))
			}
		}
	} ()

	// TODO : do stuff !
	for range time.Tick(2*time.Second) {
		writeErr := conn.WriteMessage(websocket.TextMessage, []byte("hi ! " + time.Now().String()))
		if writeErr != nil {
			fmt.Printf("WRITE ERROR : %v\n", writeErr)
			return
		}
	}
}
