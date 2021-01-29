// more information https://blog.golang.org/image
// and https://riptutorial.com/go/example/31686/loading-and-saving-image
package main

import (
	"fmt"
	"image/color"
	"image/png"
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

func main() {
	pngReader, err := os.Open("../testImage/cat_in_png.png")
	if err != nil {
		log.Fatal(err)
	}
	defer pngReader.Close()

	jpegReader, err := os.Open("../testImage/cat_in_jpg.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer jpegReader.Close()

	//should not use auto decode (image.Decode), usually get mistake
	pngImg, err := png.Decode(pngReader)
	if err != nil {
		log.Fatal(err)
	}
	switch pngImg.ColorModel() {
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

	//jpegImg, err := jpeg.Decode(jpegReader)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(jpegImg.ColorModel(), "jpeg")
}
