package main

import (
	"image/color"
	"math"
	"math/cmplx"
)

// Utility function to convert a point on a Canvas to a
// complex number given a zoom level and the camplex value
// to be shown in the center of the Canvas
// A zoom of 1 means one pixel correspond to one unit in
// the complex plane
func toCmplx(x, y int, zoom float64, center complex128) complex128 {
	return center + complex(float64(x)/zoom, float64(y)/zoom)
}

// Perform iter iterations using the mandelbrot algorithm, and return
// the magnitude of the result
// Strictly speaking, this isn't actually the Mandelbrot set, as we
// return the size of z after iter iterations, rather than the number
// of iteration it takes for z to reach a set value. Nevertheless, it
// still produces nice pictures!
func invMandelbrot(c complex128, iter int) float64 {
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

func drawInvMandelbrot(canvas *Canvas, zoom float64, center complex128, colorizer func(float64) color.Color) {
	size := canvas.Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			c := toCmplx(x-size.X/2, y-size.Y/2, zoom, center)
			mag := invMandelbrot(c, 50)
			color := colorizer(mag)
			canvas.Set(x, y, color)
		}
	}
}

