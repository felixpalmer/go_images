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
	outFilename := "blank.png"
	outFile, err := os.Create(outFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	log.Print("Saving image to: ", outFilename)
	png.Encode(outFile, m)
}
