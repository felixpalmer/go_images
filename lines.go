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

  // Draw a series of lines from the top left corner to the bottom of the image
	for x := 0; x < width; x += 8 {
	  draw_line(*canvas,
	            color.RGBA{0, 0, 0, 255},
	            Vector{0.0, 0.0}, Vector{float64(x), float64(height)})
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

func draw_line(m Canvas, color color.RGBA, from Vector, to Vector) {
	// Get the number of pixels we'll need to draw
	delta := to.Sub(from)
	length := delta.Length()
	x_step, y_step := delta.X/length, delta.Y/length
	limit := int(length + 0.5)
	for i := 0; i < limit; i++ {
		x := from.X + float64(i)*x_step
		y := from.Y + float64(i)*y_step

		// Alias for now
		m.Set(int(x), int(y), color)
	}
}

