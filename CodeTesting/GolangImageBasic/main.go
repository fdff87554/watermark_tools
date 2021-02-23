// more information https://blog.golang.org/image
// and https://riptutorial.com/go/example/31686/loading-and-saving-image
package main

import (
	"fmt"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

//// convertJPEGToPNG converts from JPEG to PNG.
//func convertJPEGToPNG(w io.Writer, r io.Reader) error {
//	img, err := jpeg.Decode(r)
//	if err != nil {
//		return err
//	}
//	return png.Encode(w, img)
//}

////ColorToGray convert image color to Grayscale
//func ColorToGray(img image.Image) image.Gray {
//	grayImg := image.NewGray(img.Bounds())
//	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
//		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
//			grayImg.Set(x, y, img.At(x, y))
//		}
//	}
//
//	return *grayImg
//}

func main() {
	grayReader, err := os.Open("../testImage/gray.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer grayReader.Close()

	//jpegReader, err := os.Open("../testImage/cat_in_jpg.jpg")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer jpegReader.Close()

	writer, err := os.Create("./test.txt")
	if err != nil {
		log.Fatal()
	}
	defer func() {
		if err := writer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	////should not use auto decode (image.Decode), usually get mistake
	//jpegImg, err := jpeg.Decode(jpegReader)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//switch jpegImg.ColorModel() {
	//case color.AlphaModel:
	//	fmt.Println("AlphaModel")
	//case color.Alpha16Model:
	//	fmt.Println("Alpha16Model")
	//case color.YCbCrModel:
	//	fmt.Println("YCbCrModel")
	//case color.NYCbCrAModel:
	//	fmt.Println("NYCbCrAModel")
	//case color.GrayModel:
	//	fmt.Println("GrayModel")
	//case color.Gray16Model:
	//	fmt.Println("Gray16Model")
	//case color.RGBAModel:
	//	fmt.Println("RGBAModel")
	//case color.NRGBAModel:
	//	fmt.Println("NRGBAModel")
	//case color.NRGBA64Model:
	//	fmt.Println("NRGBA64Model")
	//case color.RGBA64Model:
	//	fmt.Println("RGBA64Model")
	//case color.CMYKModel:
	//	fmt.Println("CMYKModel")
	//default:
	//	fmt.Println("none is above")
	//}

	grayImg, err := jpeg.Decode(grayReader)
	if err != nil {
		log.Fatal(err)
	}
	switch grayImg.ColorModel() {
	case color.AlphaModel:
		fmt.Println("AlphaModel")
	case color.Alpha16Model:
		fmt.Println("Alpha16Model")
	case color.YCbCrModel:
		fmt.Println("YCbCrModel")
	case color.NYCbCrAModel:
		fmt.Println("NYCbCrAModel")
	case color.GrayModel:
		fmt.Println("GrayModel")
	case color.Gray16Model:
		fmt.Println("Gray16Model")
	case color.RGBAModel:
		fmt.Println("RGBAModel")
	case color.NRGBAModel:
		fmt.Println("NRGBAModel")
	case color.NRGBA64Model:
		fmt.Println("NRGBA64Model")
	case color.RGBA64Model:
		fmt.Println("RGBA64Model")
	case color.CMYKModel:
		fmt.Println("CMYKModel")
	default:
		fmt.Println("none is above")
	}

	//imgTranfer := ColorToGray(grayImg)
	//fmt.Println(grayImg.Bounds().Dy(), grayImg.Bounds().Dx(), grayImg.Bounds().Dy()*grayImg.Bounds().Dx())
	//if err := imgTranfer.Pix[imgTranfer.PixOffset(0, 5000)]; err != 0 {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(err)
	//}

	for y := 0; y < grayImg.Bounds().Dy(); y++ {
		for x := 0; x < grayImg.Bounds().Dx(); x++ {

			writer.WriteString(fmt.Sprintf("%d ", grayImg.At(x, y)))
			//fmt.Println(y*grayImg.Bounds().Dx() + x, x, y, grayImg.Bounds().Dx(), grayImg.Bounds().Dy())
		}

		writer.WriteString(fmt.Sprintf("\n"))
	}
}
