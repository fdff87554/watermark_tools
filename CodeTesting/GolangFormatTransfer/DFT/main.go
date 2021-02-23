package main

import (
	"image"
	"image/jpeg"
	"log"
	"os"
)

func fdft(img image.Image) [][]float64 {

}

func main() {

	imgReader, err := os.Open("../testimage/gray.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer imgReader.Close()
	img, err := jpeg.Decode(imgReader)
	if err != nil {
		log.Fatal(err)
	}
	imgWriter, err := os.Create("../testimage/output.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer imgWriter.Close()

	imgFreq := fdft(img)
	imgIDFT := idft(imgFreq)

	if err := jpeg.Encode(imgWriter, imgIDFT, &jpeg.Options{Quality: 100}); err != nil {
		log.Fatal(err)
	}
}
