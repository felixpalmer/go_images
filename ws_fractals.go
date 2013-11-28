package main

import (
	"fmt"
	"net/http"

	"code.google.com/p/go.net/websocket"
	"github.com/gorilla/mux"
)

func UIHandler(ws *websocket.Conn) {
	for {
    var msg string
    err := websocket.Message.Receive(ws, &msg)
    if err != nil {
      fmt.Println(err)
      break
    }
    fmt.Println(msg)
  }
}

func main() {
	router := mux.NewRouter()
	router.Handle("/ws", websocket.Handler(UIHandler))
	staticHandler := http.FileServer(http.Dir("."))
	router.PathPrefix("/").Handler(staticHandler)
	http.ListenAndServe("localhost:1234", router)
}
