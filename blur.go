package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	// Load in the image file
	inFilename := "sweet_goat.jpeg"
	file, err := os.Open(inFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the image.
	m, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// Create a canvas from the image
	canvas := NewCanvas(m.Bounds())
	draw.Draw(canvas, m.Bounds(), m, image.ZP, draw.Src)

	canvas.Blur(15, new(WeightFunctionBox))
	outFilename := "blur.jpeg"
	outFile, err := os.Create(outFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	log.Print("Saving image to: ", outFilename)
	jpeg.Encode(outFile, canvas, nil)
}
