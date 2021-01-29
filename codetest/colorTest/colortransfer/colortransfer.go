package colortransfer

import "image"

func ColorToGray(img image.Image) image.Image {
	grayImg := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			grayImg.Set(x, y, img.At(x, y))
		}
	}

	return grayImg
}

//func ColorToYCbCr(img image.Image, ycbcrRatio int) image.Image {
//	ycbcrImg := image.NewYCbCr(img.Bounds(), image.YCbCrSubsampleRatio(ycbcrRatio))
//	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
//		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
//			ycbcrImg.Set(x, y, img.At(x, y))
//		}
//	}
//
//	return ycbcrImg
//}
