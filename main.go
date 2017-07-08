package main

import (
	"net/http"

	"github.com/kehrlann/webgockets/handlers"
	"fmt"
)

func main() {
	http.HandleFunc("/", handlers.HandleIndex)
	http.Handle("/ws", &handlers.WebsocketHandler{})
	server := &http.Server{Addr:":3000",Handler:nil}
	fmt.Println("Started !")
	server.ListenAndServe()
}
