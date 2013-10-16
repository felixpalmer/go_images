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
	width, height := 128, 128
	canvas := NewCanvas(image.Rect(0, 0, width, height))
	//canvas.DrawGradient()

  // Draw a set of spirals randomly over the image
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 1; i++ {
    x := 10.0
    y := 60.0
    color := color.RGBA{55,
				                60,
                        200,
                        255}
			                
		canvas.DrawSpiral(color, Vector{x, y})
  }

  canvas.Blur(2)
	outFilename := "blur.png"
	outFile, err := os.Create(outFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	log.Print("Saving image to: ", outFilename)
	png.Encode(outFile, canvas)
}
