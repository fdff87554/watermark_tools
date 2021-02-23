package colortransfer

import (
	"image"
	"image/color"
)

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

//ColorToAlpha convert image color to Alpha
func ColorToAlpha(img image.Image) image.Image {

	alphaImg := image.NewAlpha(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			alphaImg.Set(x, y, img.At(x, y))
		}
	}

	return alphaImg
}

//ColorToCMYK convert image color to CMYK
func ColorToCMYK(img image.Image) image.Image {

	CMYKImg := image.NewCMYK(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			CMYKImg.Set(x, y, img.At(x, y))
		}
	}

	return CMYKImg
}

//ColorToNRGBA convert image color to NRGBA
func ColorToNRGBA(img image.Image) image.Image {

	NRGBAImg := image.NewNRGBA(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			NRGBAImg.Set(x, y, img.At(x, y))
		}
	}

	return NRGBAImg
}

////ColorToPaletted convert image color to Paletted
//func ColorToPaletted(img image.Image) image.Image {
//	PalettedImg := image.NewPaletted(img.Bounds(), )
//	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
//		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
//			PalettedImg.Set(x, y, img.At(x, y))
//		}
//	}
//
//	return PalettedImg
//}

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

