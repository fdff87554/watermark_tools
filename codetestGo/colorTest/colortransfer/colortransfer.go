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

	_r := float32(r)
	_g := float32(g)
	_b := float32(b)

	y := 0 + 0.2990*_r + 0.5870*_g + 0.1140*_b
	cb := 128 - 0.168736*_r - 0.331264*_g + 0.5000*_b
	cr := 128 + 0.5000*_r - 0.418688*_g - 0.081312*_b

	return uint8(y), uint8(cb), uint8(cr)
}

// YCbCrToRGB() converts an YCbCr to a RGB
func YCbCrToRGB(y, cb, cr uint8) (uint8, uint8, uint8) {
	// R = Y + 1.402*(Cr - 128)
	// G = Y - 0.34414*(Cb - 128 ) - 0.71414*(Cr - 128)
	// B = Y + 1.772*( Cb - 128 )

	_y := float32(y)
	_cb := float32(cb)
	_cr := float32(cr)

	r := _y + 1.402*(_cr-128)
	g := _y - 0.344136*(_cb-128) - 0.714136*(_cr-128)
	b := _y + 1.772*(_cb-128)

	return uint8(r), uint8(g), uint8(b)
}
