package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	width, height := 128, 128
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	drawGradient(*m)
	outFilename := "gradient.png"
	outFile, err := os.Create(outFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	log.Print("Saving image to: ", outFilename)
	png.Encode(outFile, m)
}

func drawGradient(m image.RGBA) {
	size := m.Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			color := color.RGBA{
				uint8(255 * x / size.X),
				uint8(255 * y / size.Y),
				55,
				255}
			m.Set(x, y, color)
		}
	}
}
