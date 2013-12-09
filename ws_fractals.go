package main

import (
	"fmt"
	"image"
	"io"
	"net/http"
	"strconv"
	"strings"

	"code.google.com/p/go.net/websocket"
	"github.com/gorilla/mux"
)

func UIHandler(ws *websocket.Conn) {
	width, height := 512, 512
	canvas := NewCanvas(image.Rect(0, 0, width, height))
	zoom := 1600.0
	center := complex(-0.71, -0.25)
	colorizer := createColorizer("fractalGradients/gradient2.png")

	for {
		var msg string
		websocket.Message.Receive(ws, &msg)

		// Parse message
		fields := strings.Fields(msg)
		cmd := fields[0]
		if cmd == "ZOOMIN" {
			zoom *= 2.0
		} else if cmd == "ZOOMOUT" {
			zoom /= 2.0
		} else if cmd == "MOUSEDOWN" {
			x, _ := strconv.ParseInt(fields[1], 10, 0)
			y, _ := strconv.ParseInt(fields[2], 10, 0)
			center = toCmplx(int(x)-width/2, int(y)-height/2, zoom, center)
		} else if cmd == "SETGRADIENT" {
			colorizer = createColorizer(fields[1])
		} else {
			fmt.Println("Unknown command:", cmd)
			fmt.Println("Message:", msg)
		}

		// Create fractal and convert to base64
		drawInvMandelbrot(canvas, zoom, center, colorizer)
		drawMsg := fmt.Sprintf("DRAW %s", canvas.ToBase64())
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
