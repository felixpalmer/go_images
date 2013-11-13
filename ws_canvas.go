package main

import (
	"fmt"
  "image"
  "io"
  "math"
	"net/http"
  "time"

  "code.google.com/p/go.net/websocket"
	"github.com/gorilla/mux"
)

func WebsocketHandler(ws *websocket.Conn) {
	width, height := 200, 200
	canvas := NewCanvas(image.Rect(0, 0, width, height))
	canvas.DrawGradient()
  t := 0.0
  for {
    msg := ""
    for x := 0; x < width; x++ {
      y := height/2 + int(50 * math.Sin(t/10) * math.Sin(float64(x)/10.0 + t))
      //r, g, b, _ := canvas.At(x, y).RGBA()
      //color := fmt.Sprintf("#%02x%02x%02x", r/0xFF, g/0xFF, b/0xFF)
      color := "#ff0000"
      msg += fmt.Sprintf("PIXEL %d %d %s\n", x, y, color)
    }
    t += 0.3
    io.WriteString(ws, msg)
    time.Sleep(time.Second / 30.0)
  }
}

func main() {
	router := mux.NewRouter()
	router.Handle("/ws", websocket.Handler(WebsocketHandler))
	staticHandler := http.FileServer(http.Dir("."))
	router.PathPrefix("/").Handler(staticHandler)
	http.ListenAndServe("localhost:1234", router)
}
