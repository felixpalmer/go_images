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

type Canvas struct {
	image.RGBA
}

func NewCanvas(r image.Rectangle) *Canvas {
	canvas := new(Canvas)
	canvas.RGBA = *image.NewRGBA(r)
	return canvas
}

func (c Canvas) DrawGradient() {
	size := c.Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			color := color.RGBA{
				uint8(255 * x / size.X),
				uint8(255 * y / size.Y),
				55,
				255}
			c.Set(x, y, color)
		}
	}
}
