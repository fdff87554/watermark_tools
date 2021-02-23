package main

import (
	"crypto/sha256"
	"fmt"
	"flag"
	"os"

	"github.com/skip2/go-qrcode"
)

func main () {

	originStr := flag.String("s", "A0000000001", "a string you want to add to qrcode")
	qrFileName := flag.String("qr", "qr_hash.png", "qrcode file name")
	hashFileName := flag.String("ha", "hash.dat", "hash file name")

	flag.Parse()

	hash := sha256.Sum256([]byte(*originStr))
	hex_hash := fmt.Sprintf("%x", hash[:])
	fmt.Printf("hash: %s\n", hex_hash)

	f, err := os.Create(*hashFileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	
	if _, err = f.WriteString(hex_hash); err != nil {
		panic(err)
	}
	fmt.Printf("hash saved in %s\n", *hashFileName)

	err = qrcode.WriteFile(hex_hash, qrcode.Highest, 256, *qrFileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("qrcode saved in %s\n", *qrFileName)
}