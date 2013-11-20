package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
  "math"
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
    if cmplx.Abs(z) > 1000 {
      return 1000
    }
	}
	return cmplx.Abs(z)
}

// Creates a function for converting a magnitude into a color
// based on a gradient image file
func createColorizer(filename string) func(float64) color.Color {
  gradient := CanvasFromFile(filename)
  limit := gradient.Bounds().Size().Y - 1
  return func(mag float64) color.Color {
    // Clamp magnitude to size of gradient
    m := int(math.Max(math.Min(300*mag, float64(limit)), 1))
    return gradient.At(0, m)
  }
}

func main() {
	width, height := 800, 600
	canvas := NewCanvas(image.Rect(0, 0, width, height))
  colorizer := createColorizer("fractalGradients/gradient3.png")

  zoom := 22400.0
  center := complex(-0.55, 0.6)
	for x := 0; x < width; x++ {
		for y := 0; y < width; y++ {
      c := toCmplx(x - width / 2, y - height / 2, zoom, center)
      mag := mandelbrot(c, 50)
			color := colorizer(mag)
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
