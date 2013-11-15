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

func GradientHandler(ws *websocket.Conn) {
	width, height := 200, 200
	canvas := NewCanvas(image.Rect(0, 0, width, height))
	canvas.DrawGradient()
  // Warning, slow for large canvases!
  // This is definitely not an efficient way to send an image!
  for x := 0; x < width; x++ {
    for y := 0; y < height; y++ {
      r, g, b, _ := canvas.At(x, y).RGBA()
      color := fmt.Sprintf("#%02x%02x%02x", r/0xFF, g/0xFF, b/0xFF)
      msg := fmt.Sprintf("PIXEL %d %d %s\n", x, y, color)
      io.WriteString(ws, msg)
    }
    time.Sleep(time.Second / 30.0)
  }
}

func GraphHandler(ws *websocket.Conn) {
	width, height := 512, 512
  t := 0.0
  for {
    msg := "CLEAR\n"
    for x := 0; x < width; x++ {
      y := height/2 + int(50 * math.Sin(t/10) * math.Sin(float64(x)/10.0 + t))
      y2 := height/3 + int(60 * math.Sin(t/9) * math.Sin(float64(x)/20.0 + 2.0 * t))
      msg += fmt.Sprintf("PIXEL %d %d %s\n", x, y, "#ff0000")
      msg += fmt.Sprintf("PIXEL %d %d %s\n", x, y2, "#00ff00")
    }
    t += 0.3
    io.WriteString(ws, msg)
    time.Sleep(time.Second / 30.0)
  }
}

func HelloHandler(ws *websocket.Conn) {
  io.WriteString(ws, "HELLO")
}

func main() {
	router := mux.NewRouter()
	//router.Handle("/ws", websocket.Handler(GradientHandler))
	router.Handle("/ws", websocket.Handler(GraphHandler))
	staticHandler := http.FileServer(http.Dir("."))
	router.PathPrefix("/").Handler(staticHandler)
	http.ListenAndServe("localhost:1234", router)
}
