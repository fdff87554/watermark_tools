package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
)

func merge(img, code color.Color) color.Color {

	r, g, b, _ := img.RGBA()
	chk, _, _, _ := code.RGBA()
	if chk == 0 {
		return color.RGBA{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
			A: uint8(rand.Intn(250-248) + 248),
		}
	} else {
		return color.RGBA{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
			A: uint8(rand.Intn(255-251) + 251),
		}
	}
}

func imageMerge(img, code image.Image) image.Image {

	weight, height := img.Bounds().Dx(), img.Bounds().Dy()
	mrgImg := image.NewRGBA(img.Bounds())
	for y := 0; y < height; y++ {
		for x := 0; x < weight; x++ {
			if y < code.Bounds().Dy() && x < code.Bounds().Dx() {
				mrgImg.Set(x, y, merge(img.At(x, y), code.At(x, y)))
			} else if rand.Int()%2 == 0 {
				r, g, b, _ := img.At(x, y).RGBA()
				mrgImg.Set(x, y, color.RGBA{
					R: uint8(r),
					G: uint8(g),
					B: uint8(b),
					A: uint8(rand.Intn(255-251) + 251),
				})
			} else {
				r, g, b, _ := img.At(x, y).RGBA()
				mrgImg.Set(x, y, color.RGBA{
					R: uint8(r),
					G: uint8(g),
					B: uint8(b),
					A: uint8(rand.Intn(248-240) + 240),
				})
			}
		}
	}

	return mrgImg
}

func imageSep(img image.Image) image.Image {

	weight, height := img.Bounds().Dx(), img.Bounds().Dy()
	sepImg := image.NewRGBA(img.Bounds())
	for y := 0; y < height; y++ {
		for x := 0; x < weight; x++ {
			_, _, _, chk := img.At(x, y).RGBA()
			if uint8(chk) >= 248 && uint8(chk) <= 250 {
				sepImg.Set(x, y, color.RGBA{
					R: 0,
					G: 0,
					B: 0,
					A: 255,
				})
			} else {
				sepImg.Set(x, y, color.RGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 255,
				})
			}
		}
	}

	return sepImg
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
	mrgWriter, err := os.Create("./output/merge.png")
	if err != nil {
		log.Fatal(err)
	}
	defer mrgWriter.Close()

	img, err := png.Decode(imgReader)
	if err != nil {
		log.Fatal(err)
	}
	code, err := png.Decode(codeReader)
	if err != nil {
		log.Fatal(err)
	}

	mergeImg := imageMerge(img, code)
	if err := png.Encode(mrgWriter, mergeImg); err != nil {
		log.Fatal(err)
	}

	deReader, err := os.Open("./output/merge.png")
	if err != nil {
		log.Fatal(err)
	}
	defer deReader.Close()
	deWriter, err := os.Create("./output/decode.png")
	if err != nil {
		log.Fatal(err)
	}
	defer deWriter.Close()
	mergeImg, err = png.Decode(deReader)
	if err != nil {
		log.Fatal(err)
	}

	deImg := imageSep(mergeImg)
	if err := png.Encode(deWriter, deImg); err != nil {
		log.Fatal(err)
	}
}
