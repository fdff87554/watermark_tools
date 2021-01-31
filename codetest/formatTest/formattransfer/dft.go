package formattransfer

import (
	"formatTest/colortransfer"
	"github.com/mjibson/go-dsp/fft"
	"image"
	"image/color"
	"math/cmplx"
)

func FFFT(img image.Image) [][]complex128 {
	switch img.ColorModel() {
	case color.GrayModel:
	case color.Gray16Model:
	default:
		img = colortransfer.ColorToGray(img)
	}
	bounds := img.Bounds()
	imgPixels := make([][]float64, bounds.Dy())
	for y := 0; y < bounds.Dy(); y++ {
		imgPixels[y] = make([]float64, bounds.Dx())
		for x := 0; x < bounds.Dx(); x++ {
			r, _, _, _ := img.At(x, y).RGBA()
			imgPixels[y][x] = float64(r)
		}
	}
	// apply discrete fourier transform on realPixels.
	coeffs := fft.FFT2Real(imgPixels)

	return coeffs
}

func IFFT(freq [][]complex128, bounds image.Rectangle) image.Image {
	// use inverse fourier transform to transform fft
	// values back to the original image.
	coeffs := fft.IFFT2(freq)

	// write everything to a new image
	outImage := image.NewGray(bounds)

	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			px := uint8(cmplx.Abs(coeffs[y][x]))
			outImage.SetGray(x, y, color.Gray{Y: px})
		}
	}

	return outImage
}