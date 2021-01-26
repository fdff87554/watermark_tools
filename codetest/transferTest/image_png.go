package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
)

func main() {
	reader, err := os.Open("./test_img/cat_in_png.png")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	img, err := png.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%T", img)

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			fmt.Println(r, g, b, a)
			y, cb, cr := RGBToYCbCr(r, g, b)

		}
	}
}
