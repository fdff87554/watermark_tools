package main

import (
	"formatTest/formattransfer"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	imgReader, err := os.Open("../testImage/cat_in_jpg.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := imgReader.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	imgWriter, err := os.Create("../testImage/output/cat_in_jpg.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := imgWriter.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	img, err := jpeg.Decode(imgReader)
	if err != nil {
		log.Fatal(err)
	}
	imgFreq, err := formattransfer.FDCT(img)
	if err != nil {
		log.Fatal(err)
	}
	imgBack, err := formattransfer.IDCT(imgFreq, img.Bounds().Dx(), img.Bounds().Dy())
	if err != nil {
		log.Fatal(err)
	}
	if err = jpeg.Encode(imgWriter, imgBack, &jpeg.Options{Quality: 100}); err != nil {
		log.Fatal(err)
	}
}