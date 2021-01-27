package colortransfer

// the output will have deviation is count by functions
// so the code should use godoc, will have lower deviation

// RGBToYCbCr() converts an RGB to a YCbCr
func RGBToYCbCr(r, g, b uint8) (uint8, uint8, uint8) {
	// The JFIF specification says:
	// 		Y = 0 + 0.2990*R + 0.5870*G + 0.1140*B
	// 		Cb = 128 - 0.1687*R - 0.3313*G + 0.5000*B
	// 		Cr = 128 + 0.5000*R - 0.4187*G - 0.0813*B
	// https://www.w3.org/Graphics/JPEG/jfif3.pdf

	r_ := float32(r)
	g_ := float32(g)
	b_ := float32(b)

	y := 0 + 0.2990*r_ + 0.5870*g_ + 0.1140*b_
	cb := 128 - 0.168736*r_ - 0.331264*g_ + 0.5000*b_
	cr := 128 + 0.5000*r_ - 0.418688*g_ - 0.081312*b_

	return uint8(y), uint8(cb), uint8(cr)
}

// YCbCrToRGB() converts an YCbCr to a RGB
func YCbCrToRGB(y, cb, cr uint8) (uint8, uint8, uint8) {
	// R = Y + 1.402*(Cr - 128)
	// G = Y - 0.34414*(Cb - 128 ) - 0.71414*(Cr - 128)
	// B = Y + 1.772*( Cb - 128 )

	y_ := float32(y)
	cb_ := float32(cb)
	cr_ := float32(cr)

	r := y_ + 1.402*(cr_-128)
	g := y_ - 0.344136*(cb_-128) - 0.714136*(cr_-128)
	b := y_ + 1.772*(cb_-128)

	return uint8(r), uint8(g), uint8(b)
}
