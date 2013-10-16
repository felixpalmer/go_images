package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	width, height := 512, 512
	canvas := NewCanvas(image.Rect(0, 0, width, height))
	canvas.DrawGradient()

	// Draw a series of lines from the top left corner to the bottom of the image
	for x := 0; x < width; x += 8 {
		canvas.DrawLine(color.RGBA{0, 0, 0, 255},
			Vector{0.0, 0.0},
			Vector{float64(x), float64(height)})
	}

	outFilename := "lines.png"
	outFile, err := os.Create(outFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	log.Print("Saving image to: ", outFilename)
	png.Encode(outFile, canvas)
}
