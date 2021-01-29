package main

import (
	"image/png"
	"log"
	"os"
)

func main() {
	reader, err := os.Open("../testImage/cat_in_png.png")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	// The code here is for testing the different of paper function and godoc code
	// ---------------------------------------------------------------------------
	// writer_RGB, err := os.Create("./RGB.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer writer_RGB.Close()

	// writer_YCbCr_own, err := os.Create("./YCbCr_own.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer writer_YCbCr_own.Close()

	// writer_YCbCr_doc, err := os.Create("./YCbCr_doc.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer writer_YCbCr_doc.Close()

	// img, err := png.Decode(reader)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%T", img)

	// for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
	// 	for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
	// 		r, g, b, _ := img.At(x, y).RGBA()
	// 		if r != 0 && g != 0 && b != 0 {
	// 			writer_RGB.Write([]byte(fmt.Sprintf("%d, %d, %d\n", uint8(r), uint8(g), uint8(b))))
	// 			y1, cb1, cr1 := colortransfer.RGBToYCbCr(uint8(r), uint8(g), uint8(b))
	// 			r1_, g1_, b1_ := colortransfer.YCbCrToRGB(uint8(y1), uint8(cb1), uint8(cr1))
	// 			writer_YCbCr_own.Write([]byte(fmt.Sprintf("%d, %d, %d : %d, %d, %d\n", y1, cb1, cr1, r1_, g1_, b1_)))
	// 			y2, cb2, cr2 := colortransfer.GoDocRGBToYCbCr(uint8(r), uint8(g), uint8(b))
	// 			r2_, g2_, b2_ := colortransfer.YCbCrToRGB(uint8(y2), uint8(cb2), uint8(cr2))
	// 			writer_YCbCr_doc.Write([]byte(fmt.Sprintf("%d, %d, %d : %d, %d, %d\n", y2, cb2, cr2, r2_, g2_, b2_)))
	// 		}
	// 	}
	// }

	// The code here is for changing color type from RBG to YCbCr
	// ----------------------------------------------------------
	img, err := png.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			img.At(x, y)
		}
	}
}
