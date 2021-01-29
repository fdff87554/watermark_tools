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
	pndDct, err := formattransfer.FDCT(pngImg)
}
