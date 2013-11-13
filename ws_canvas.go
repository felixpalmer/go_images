package main

import (
	"fmt"
  "image"
  "io"
	"net/http"
  "time"

  "code.google.com/p/go.net/websocket"
	"github.com/gorilla/mux"
)

func WebsocketHandler(ws *websocket.Conn) {
	width, height := 200, 200
	canvas := NewCanvas(image.Rect(0, 0, width, height))
	canvas.DrawGradient()
  for {
    msg := ""
    for x := 0; x < width; x++ {
      for y := 0; y < width; y++ {
        r, g, b, _ := canvas.At(x, y).RGBA()
        color := fmt.Sprintf("#%02x%02x%02x", r/0xFF, g/0xFF, b/0xFF)
        msg += fmt.Sprintf("PIXEL %d %d %s\n", x, y, color)
      }
    }
    io.WriteString(ws, msg)
    time.Sleep(10 * time.Second)
  }
}

func main() {
	router := mux.NewRouter()
	router.Handle("/ws", websocket.Handler(WebsocketHandler))
	staticHandler := http.FileServer(http.Dir("."))
	router.PathPrefix("/").Handler(staticHandler)
	http.ListenAndServe("localhost:1234", router)
}
