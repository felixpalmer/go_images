package main

import (
  "bytes"
  "encoding/base64"
	"fmt"
  "image"
  "image/png"
	"net/http"
  "io"

	"code.google.com/p/go.net/websocket"
	"github.com/gorilla/mux"
)

func UIHandler(ws *websocket.Conn) {
	width, height := 800, 600
	canvas := NewCanvas(image.Rect(0, 0, width, height))
	zoom := 16000.0
	center := complex(-0.71, -0.25)
	colorizer := createColorizer("fractalGradients/gradient1.png")


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
    imgBuf := new(bytes.Buffer)
    imgEncoder := base64.NewEncoder(base64.StdEncoding, imgBuf)
    png.Encode(imgEncoder, canvas)
    imgEncoder.Close()
    drawMsg := fmt.Sprintf("DRAW %s\n", imgBuf)
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
