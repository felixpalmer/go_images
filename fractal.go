package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
)

// Utility function to convert a point on a Canvas to a
// complex number given a zoom level and the camplex value
// to be shown in the center of the Canvas
// A zoom of 1 means one pixel correspond to one unit in
// the complex plane
func toCmplx(x, y int, zoom float64, center complex128) complex128 {
	return center + complex(float64(x) / zoom, float64(y) / zoom)
}

// Perform iter iterations using the mandelbrot algorithm, and return
// the magnitude of the result
func mandelbrot(c complex128, iter int) float64 {
	z := complex(0, 0)
	for i := 0; i < iter; i++ {
		z = z*z + c
	}
	return cmplx.Abs(z)
}

func main() {
	width, height := 800, 600
	canvas := NewCanvas(image.Rect(0, 0, width, height))

  zoom := 250.0
  center := complex(-0.6, 0)
	for x := 0; x < width; x++ {
		for y := 0; y < width; y++ {
      c := toCmplx(x - width / 2, y - height / 2, zoom, center)
      mag := mandelbrot(c, 10)
			color := color.RGBA{uint8(mag), 0, 0, 255}
			canvas.Set(x, y, color)
		}
	}

	outFilename := "fractal.png"
	outFile, err := os.Create(outFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	log.Print("Saving image to: ", outFilename)
	png.Encode(outFile, canvas)
}
