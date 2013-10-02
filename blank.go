package main

import (
	"image"
	"image/png"
	"log"
	"os"
)

func main() {
	width, height := 128, 128
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	out_filename := "blank.png"
	out_file, err := os.Create(out_filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out_file.Close()
	log.Print("Saving image to: ", out_filename)
	png.Encode(out_file, m)
}
