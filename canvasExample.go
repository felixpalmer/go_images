package main

import (
	"image"
	"image/png"
	"log"
	"os"
)

func main() {
	width, height := 128, 128
	canvas := NewCanvas(image.Rect(0, 0, width, height))
	canvas.DrawGradient()
	outFilename := "canvas.png"
	outFile, err := os.Create(outFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	log.Print("Saving image to: ", outFilename)
	png.Encode(outFile, canvas)
}
