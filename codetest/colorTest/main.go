package main

import (
	"image/png"
	"log"
	"os"
)

func main() {
	pngReader, err := os.Open("../testImage/cat_in_png.png")
	if err != nil {
		log.Fatal(err)
	}
	defer pngReader.Close()

	pngImg, err := png.Decode(pngReader)
	if err != nil {
		log.Fatal(err)
	}
	grayImg := colortransfer.ColorToGray(pngImg)
	grayWriter, err := os.Create("../testImage/transferGray.png")
	if err != nil {
		log.Fatal(err)
	}
	defer grayWriter.Close()

	if err := png.Encode(grayWriter, grayImg); err != nil {
		log.Fatal(err)
	}
}
