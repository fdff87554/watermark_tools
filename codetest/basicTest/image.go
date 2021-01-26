// more information https://blog.golang.org/image
// and https://riptutorial.com/go/example/31686/loading-and-saving-image
package main

import (
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
)

// convertJPEGToPNG converts from JPEG to PNG.
func convertJPEGToPNG(w io.Writer, r io.Reader) error {
	img, err := jpeg.Decode(r)
	if err != nil {
		return err
	}
	return png.Encode(w, img)
}

func main() {
	// testing basic image input and output in golang
	reader, err := os.Open("../testImage/cat_in_jpg.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	writer, err := os.Create("../testImage/cat_transfer_from_jpg.png")
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	convertJPEGToPNG(writer, reader)
}
