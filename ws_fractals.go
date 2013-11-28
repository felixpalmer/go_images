package main

import (
	"fmt"
  "image"
	"net/http"
  "io"

	"code.google.com/p/go.net/websocket"
	"github.com/gorilla/mux"
)

func UIHandler(ws *websocket.Conn) {
	width, height := 800, 600
	canvas := NewCanvas(image.Rect(0, 0, width, height))
	zoom := 1600.0
	center := complex(-0.71, -0.25)
	colorizer := createColorizer("fractalGradients/gradient2.png")

	for {
    var msg string
    err := websocket.Message.Receive(ws, &msg)
    if err != nil {
      fmt.Println(err)
      break
    }
    fmt.Println(msg)

    // Create fractal and convert to base64
    drawInvMandelbrot(canvas, zoom, center, colorizer)
    drawMsg := fmt.Sprintf("DRAW %s\n", canvas.ToBase64())
    io.WriteString(ws, drawMsg)
  }
}

func main() {
	router := mux.NewRouter()
	router.Handle("/ws", websocket.Handler(UIHandler))
	staticHandler := http.FileServer(http.Dir("."))
	router.PathPrefix("/").Handler(staticHandler)
	http.ListenAndServe("localhost:1234", router)
}
