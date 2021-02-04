package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func fixPixel(img color.Color) color.Color {

	r, g, b, _ := img.RGBA()
	//find the minimum of r, g, b
	if r <= g && r <= b {
		if r == 0 {
			r++
		} else {
			r--
		}
	} else if g <= r && g <= b {
		if g == 0 {
			g++
		} else {
			g--
		}
	} else {
		if b == 0 {
			b++
		} else {
			b--
		}
	}

	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: 255,
	}
}

//imageMerge merge the two input image
func imageMerge(originalImg, insertImg image.Image) image.Image {

	//this code is adding sample black white QRcode into an image by the sample of
	//https://elisklar.medium.com/imagehash-easy-steganography-240b92b586e2
	outImg := image.NewRGBA(originalImg.Bounds())
	for y := 0; y < originalImg.Bounds().Dy(); y++ {
		for x := 0; x < originalImg.Bounds().Dx(); x++ {
			if y < insertImg.Bounds().Dy() && x < insertImg.Bounds().Dx() {
				r, g, b, _ := originalImg.At(x, y).RGBA()
				chkR, chkG, chkB, _ := insertImg.At(x, y).RGBA()
				//if chk == 255, means insertImg pixels is white, else is black
				if ((chkR+chkG+chkB)%2 == 0 && (r+g+b)%2 != 0) || ((chkR+chkG+chkB)%2 != 0 && (r+g+b)%2 == 0) {
					outImg.Set(x, y, fixPixel(originalImg.At(x, y)))
				} else {
					outImg.Set(x, y, originalImg.At(x, y))
				}
			} else {
				outImg.Set(x, y, originalImg.At(x, y))
			}
		}
	}

	return outImg
}

func imageSep(img image.Image) image.Image {

	outImg := image.NewRGBA(img.Bounds())
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dy(); x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			if (r+g+b)%2 != 0 {
				outImg.Set(x, y, color.RGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 255,
				})
			} else {
				outImg.Set(x, y, color.RGBA{
					R: 0,
					G: 0,
					B: 0,
					A: 255,
				})
			}
		}
	}

	return outImg
}

func main() {

	imgReader, err := os.Open("../../testImage/cat_1200x600.png")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := imgReader.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	codeReader, err := os.Open("../../testImage/NISRA.png")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := codeReader.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	imgWriter, err := os.Create("../../testImage/output/merge.png")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := imgWriter.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	codeWriter, err := os.Create("../../testImage/output/unmerge.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := codeWriter.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	img, err := png.Decode(imgReader)
	if err != nil {
		log.Fatal(err)
	}
	codeImg, err := png.Decode(codeReader)
	if err != nil {
		log.Fatal(err)
	}

	outImg := imageMerge(img, codeImg)
	if err := png.Encode(imgWriter, outImg); err != nil {
		log.Fatal(err)
	}

	margeReader, err := os.Open("../../testImage/output/merge.png")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := margeReader.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	margeImg, err := png.Decode(margeReader)
	if err != nil {
		log.Fatal(err)
	}

	outCode := imageSep(margeImg)
	if err := jpeg.Encode(codeWriter, outCode, &jpeg.Options{Quality: 100}); err != nil {
		log.Fatal(err)
	}
}