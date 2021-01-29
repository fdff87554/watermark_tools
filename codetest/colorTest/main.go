package main

import (
	"errors"
	"fmt"
	"github.com/fdff87554/watermark_tools/tree/main/codetest/colorTest/colortransfer"
	"image"
	"image/png"
	"log"
	"os"
)

type ColorSamples int

const (
	Gray ColorSamples = iota
	Alpha
	CMYK
	NRGBA
	YCbCr
	//Paletted
)

func imageTranfer(img image.Image, colorType int) error {
	writer, err := os.Create(fmt.Sprintf("../testImage/output/transfer%d.png", colorType))
	if err != nil {
		return err
	}
	defer writer.Close()

	switch ColorSamples(colorType) {
	case Gray:
		img := colortransfer.ColorToGray(img)
		if err := png.Encode(writer, img); err != nil {
			log.Fatal(err)
		}
		return nil
	case Alpha:
		img := colortransfer.ColorToAlpha(img)
		if err := png.Encode(writer, img); err != nil {
			log.Fatal(err)
		}
		return nil
	case CMYK:
		img := colortransfer.ColorToCMYK(img)
		if err := png.Encode(writer, img); err != nil {
			log.Fatal(err)
		}
		return nil
	case NRGBA:
		img := colortransfer.ColorToNRGBA(img)
		if err := png.Encode(writer, img); err != nil {
			log.Fatal(err)
		}
		return nil
	case YCbCr:
		img := colortransfer.ColorToYCbCr(img, 0)
		if err := png.Encode(writer, img); err != nil {
			log.Fatal(err)
		}
		return nil
	//case Paletted:
	//	img := colortransfer.ColorToPaletted(img)
	//	if err := png.Encode(writer, img); err != nil {
	//		log.Fatal(err)
	//	}
	//	return nil
	default:
		return errors.New("none define type")
	}
}

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

	if err := imageTranfer(pngImg, 0); err != nil {
		log.Fatal(err)
	}
	if err := imageTranfer(pngImg, 1); err != nil {
		log.Fatal(err)
	}
	if err := imageTranfer(pngImg, 2); err != nil {
		log.Fatal(err)
	}
	if err := imageTranfer(pngImg, 3); err != nil {
		log.Fatal(err)
	}
	if err := imageTranfer(pngImg, 4); err != nil {
		log.Fatal(err)
	}
}
