package main

import (
	"formatTest/formattransfer"
	"image/png"
	"log"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func main() {
	imgReader, err := os.Open("../testImage/cat_in_png.png")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := imgReader.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	imgWriter, err := os.Create("../testImage/output/DFT_test.png")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := imgWriter.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	txtWriter, err := os.Create("../testImage/output/FFT_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := txtWriter.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	img, err := png.Decode(imgReader)
	if err != nil {
		log.Fatal(err)
	}
	imgFFFT := formattransfer.FFFT(img)

	imgIFFT := formattransfer.IFFT(imgFFFT, img.Bounds())

	if err := png.Encode(imgWriter, imgIFFT); err != nil {
		log.Fatal(err)
	}
}