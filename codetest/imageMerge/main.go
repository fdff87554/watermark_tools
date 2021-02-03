package main

import (
	"image"
	"image/color"
)

//ColorToYCbCr convert image color to YCbCr colormodel
func ColorToYCbCr(img image.Image, ycbcrRatio int) image.Image {

	ycbcrImg := image.NewYCbCr(img.Bounds(), image.YCbCrSubsampleRatio(ycbcrRatio))
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			y_, cb, cr := color.RGBToYCbCr(uint8(r), uint8(g), uint8(b))

			ycbcrImg.Y[ycbcrImg.YOffset(x, y)] = y_
			ycbcrImg.Cb[ycbcrImg.COffset(x, y)] = cb
			ycbcrImg.Cr[ycbcrImg.COffset(x, y)] = cr
		}
	}

	return ycbcrImg
}

//ColorToGray convert image color to Grayscale
func ColorToGray(img image.Image) image.Image {

	grayImg := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			grayImg.Set(x, y, img.At(x, y))
		}
	}

	return grayImg
}

//imageMerge merge the two input image
func imageMerge(originalImg, insertImg image.Image) (image.Image) {

	originalImg = ColorToYCbCr(originalImg, 0)
	insertImg = ColorToGray(insertImg)
}

func main() {
	
}