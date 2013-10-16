package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	width, height := 2048, 1024
	canvas := NewCanvas(image.Rect(0, 0, width, height))
	canvas.DrawGradient()

	// Draw a set of spirals randomly over the image
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 100; i++ {
		x := float64(width) * rand.Float64()
		y := float64(height) * rand.Float64()
		color := color.RGBA{uint8(rand.Intn(255)),
			uint8(rand.Intn(255)),
			uint8(rand.Intn(255)),
			255}

		canvas.DrawSpiral(color, Vector{x, y})
	}

	outFilename := "spirals.png"
	outFile, err := os.Create(outFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	log.Print("Saving image to: ", outFilename)
	png.Encode(outFile, canvas)
}
