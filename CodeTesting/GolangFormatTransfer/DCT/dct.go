package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

//This part is for math tools
//matrixTransfer will turn NxM matrix array to MxN
func matrixTransfer(matrix [][]float64) [][]float64 {

	row, col := len(matrix), len(matrix[0])

	m := make([][]float64, col)
	for i := range m {
		m[i] = make([]float64, row)
	}

	for y := 0; y < row; y++ {
		for x := 0; x < col; x++ {
			m[x][y] = matrix[y][x]
		}
	}

	return m
}

//matrixDot will do matrix dot for MxN and NxW matrix array
func matrixDot(m1 [][]float64, m2 [][]float64) ([][]float64, error) {

	row1, col1, row2, col2 := len(m1), len(m1[0]), len(m2), len(m2[0])
	if col1 != row2 {
		return nil, errors.New(fmt.Sprintf("ValueError: shapes (%d,%d) and (%d,%d) not aligned: %d (dim 1) != %d (dim0)", row1, col1, row2, col2, col1, row2))
	}

	matrix := make([][]float64, row1)
	for i := range matrix {
		matrix[i] = make([]float64, col2)
	}

	for i := 0; i < row1; i++ {
		for j := 0; j < col2; j++ {
			for k := 0; k < row2; k++ {
				matrix[i][j] += matrix[i][j] + m1[i][k]*m2[k][j]
			}
		}
	}

	return matrix, nil
}

//This part is for color handle of image




//This part is for DCT image transfer

func fdctBlock()

func fdct(img image.Gray, weight, height int) [][]float64 {

	freq := make([][]float64, height)
	for i := range freq {
		freq[i] = make([]float64, weight)
	}

	for y := 0; y < height; y+=8 {
		for x := 0; x < weight; x+=8 {
			block
		}
	}
}

func fdctGray(img image.Image) ([][][]float64, error) {

	grayImg := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			grayImg.Set(x, y, img.At(x, y))
		}
	}

	weight, height := img.Bounds().Dx(), img.Bounds().Dy()
	if weight%8 != 0 {
		weight += 8 - weight%8
	}
	if height%8 != 0 {
		height += 8 - height%8
	}
	//Gray only have one parameter Y in a pixel
	freq := make([][][]float64, 1)
	for i := range freq {
		freq[i] = make([][]float64, height)
		for j := range freq[i] {
			freq[i][j] = make([]float64, weight)
		}
	}

	for i := range freq {
		freq[i] = fdct(*grayImg, weight, height)
	}
}

//func fdctYCbCr(img image.Image) ([][][]float64, error) {
//
//	weight, height := img.Bounds().Dx(), img.Bounds().Dy()
//	//YCbCr have three parameter Y, Cb, Cr in a pixel
//	freq := make([][][]float64, 3)
//	for i := range freq {
//		freq[i] = make([][]float64, height)
//		for j := range freq[i] {
//			freq[i][j] = make([]float64, weight)
//		}
//	}
//}
//
//func fdctRGBA(img image.Image) ([][][]float64, error) {
//
//}

func FDCT(img image.Image) ([][][]float64, error) {

	switch img.ColorModel() {
	case color.GrayModel:
		return fdctGray(img)
	case color.Gray16Model:
		return fdctGray(img)
	//case color.YCbCrModel:
	//	return fdctYCbCr(img)
	//case color.NYCbCrAModel:
	//	return fdctYCbCr(img)
	default:
		return nil, errors.New("type is not allow")
	}
}

func IDCT(img [][][]float64) (image.Image, error) {


}

func main() {

	//Set image reader & writer
	imgReader, err := os.Open("../testimage/gray.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer imgReader.Close()

	img, _, err := image.Decode(imgReader)
	if err != nil {
		log.Fatal(err)
	}
	freq, err := FDCT(img)
	if err != nil {
		log.Fatal(err)
	}
	deImg, err := IDCT(freq)
	if err != nil {
		log.Fatal(err)
	}
}