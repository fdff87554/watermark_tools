package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func merge(img color.Color, code color.Color) color.Color {

	r, g, b, _ := img.RGBA()
	cr, cg, cb, _ := code.RGBA()
	return color.RGBA{
		R: uint8(r)&248 + uint8(cr)&7,
		G: uint8(g)&248 + uint8(cg)&7,
		B: uint8(b)&248 + uint8(cb)&7,
		A: 255,
	}
}

func imageMerge(img image.Image, code image.Image) image.Image {

	outImg := image.NewRGBA(img.Bounds())
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			if y < code.Bounds().Dy() && x < code.Bounds().Dx() {
				outImg.Set(x, y, merge(img.At(x, y), code.At(x, y)))
			} else {
				outImg.Set(x, y, img.At(x, y))
			}
		}
	}

	return outImg
}

func imageSep(img image.Image) image.Image {

	outImg := image.NewRGBA(img.Bounds())
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			outImg.Set(x, y, color.RGBA{
				R: uint8(r) & 3 << 6,
				G: uint8(g) & 3 << 6,
				B: uint8(b) & 3 << 6,
				A: 255,
			})
		}
	}

	return outImg
}

func main() {

	imgReader, err := os.Open("./input/cat_1200x600.png")
	if err != nil {
		log.Fatal(err)
	}
	defer imgReader.Close()
	codeReader, err := os.Open("./input/qrcode.png")
	if err != nil {
		log.Fatal(err)
	}
	defer codeReader.Close()

	mergeWriter, err := os.Create("./output/merge.png")
	if err != nil {
		log.Fatal(err)
	}
	defer mergeWriter.Close()
	mergeReader, err := os.Open("./output/merge.png")
	if err != nil {
		log.Fatal(err)
	}
	defer mergeReader.Close()
	codeWriter, err := os.Create("./output/qrcode.png")
	if err != nil {
		log.Fatal(err)
	}
	defer codeWriter.Close()

	img, err := png.Decode(imgReader)
	if err != nil {
		log.Fatal(err)
	}
	code, err := png.Decode(codeReader)
	if err != nil {
		log.Fatal(err)
	}

	mergeImg := imageMerge(img, code)
	if err := png.Encode(mergeWriter, mergeImg); err != nil {
		log.Fatal(err)
	}

	mergeImg, err = png.Decode(mergeReader)
	if err != nil {
		log.Fatal(err)
	}
	codeImg := imageSep(mergeImg)
	if err := png.Encode(codeWriter, codeImg); err != nil {
		log.Fatal(err)
	}
}
