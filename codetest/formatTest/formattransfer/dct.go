package formattransfer

import (
	"formatTest/colortransfer"
	"formatTest/util"
	"image"
	"image/color"
	"math"
)

//dctMatrix will create a dct matrix in N x N size
func dctMatrix(N int) [][]float64 {

	U := make([][]float64, N)
	for i := range U {
		U[i] = make([]float64, N)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == 0 {
				U[i][j] = 1 / math.Sqrt(float64(N))
			} else {
				U[i][j] = math.Sqrt(2.0/float64(N))*math.Cos(((2.0 * float64(j) + 1) * float64(i) * math.Pi)/2.0 * float64(N))
			}
		}
	}

	return U
}

func blockFdct(A [][]float64) ([][]float64, error) {

	U := dctMatrix(8)
	Ut := util.MatrixTranfer(U)
	A, err := util.MatrixDot(U, A)
	if err != nil {
		return nil, err
	}
	A, err = util.MatrixDot(A, Ut)
	if err != nil {
		return nil, err
	}

	return A, nil
}

func fdct(img [][]float64) ([][]float64, error) {

	weight, height := len(img[0]), len(img)
	imgFreq := make([][]float64, height)
	for i := range imgFreq{
		imgFreq[i] = make([]float64, weight)
	}

	for y := 0; y < height; y+=8 {
		for x := 0; x < weight; x+=8 {
			block := img[x:x+8][y:y+8]
			blockFreq, err := blockFdct(block)
			if err != nil {
				return nil, err
			}
			for i := 0; i < 8; i++ {
				for j := 0; j < 8; j++ {
					imgFreq[x+i][y+j] = blockFreq[i][j]
				}
			}
		}
	}

	return imgFreq, nil
}

func fdctGray(img image.Image) ([][]float64, error){

	weight, height := img.Bounds().Dx(), img.Bounds().Dy()
	if weight % 8 != 0 {
		weight += 8 - weight%8
	}
	if height % 8 != 0 {
		height += 8 - height%8
	}

	imgGray := colortransfer.ColorToGray(img)
	grayInform := make([][]float64, height)
	for i := range grayInform {
		grayInform[i] = make([]float64, weight)
	}
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			grayInform[x][y] = float64(imgGray.Pix[imgGray.PixOffset(x, y)])
		}
	}

	imgFreq, err := fdct(grayInform)
	if err != nil {
		return nil, err
	}

	return imgFreq, nil
}

//func fdctColor(img image.Image) ([][]float32, error){
//
//	weight, height := img.Bounds().Dx(), img.Bounds().Dy()
//	ycbcrImg := ColorToYCbCr(img, 0)
//}

// FDCT is for Forward Discrete Cosine Transformation
func FDCT(img image.Image) ([][]float64, error){

	switch img.ColorModel() {
	case color.GrayModel:
		return fdctGray(img)
	case color.Gray16Model:
		return fdctGray(img)
	default:
		return fdctGray(img)
	}
}

func blockIdct(A [][]float64) ([][]float64, error) {

	U := dctMatrix(8)
	Ut := util.MatrixTranfer(U)
	A, err := util.MatrixDot(Ut, A)
	if err != nil {
		return nil, err
	}
	A, err = util.MatrixDot(A, U)
	if err != nil {
		return nil, err
	}

	return A, nil
}

func idct(freq [][]float64) ([][]float64, error) {

	weight, height := len(freq[0]), len(freq)
	img := make([][]float64, height)
	for i := range img{
		img[i] = make([]float64, weight)
	}

	for y := 0; y < height; y+=8 {
		for x := 0; x < weight; x+=8 {
			blockFreq := freq[x:x+8][y:y+8]
			block, err := blockIdct(blockFreq)
			if err != nil {
				return nil, err
			}
			for i := 0; i < 8; i++ {
				for j := 0; j < 8; j++ {
					img[x+i][y+j] = block[i][j]
				}
			}
		}
	}

	return img, nil
}

func idctGray(freq [][]float64, w, h int) (image.Image, error){

	grayInform, err := idct(freq)
	if err != nil {
		return nil, err
	}

	grayImg := image.NewGray(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			grayImg.SetGray(x, y, color.Gray{Y: uint8(grayInform[x][y])})
		}
	}

	return grayImg, nil
}

// IDCT is for Inverse Discrete Consine Transformation
func IDCT(freq [][]float64, w, h int) (image.Image, error){

	return idctGray(freq, w, h)
}
