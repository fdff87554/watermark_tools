How it works
Each channel (red, green, blue) of each pixel in an image is represented by an 8-bit value. To hide the secret image inside the cover image, we replace the n least significant bits of the cover pixel value with the same number of most significant bits from the secret pixel value. Example, using 3 hidden bits:

Cover pixel: (167, 93, 27) == (10100111, 01011101, 00011011)    
Secret pixel: (67, 200, 105) == (01000011, 11001000, 01101001)    
Output pixel: (162, 94, 27) == (10100010, 01011110, 00011011)

The output colour is almost indistinguishable from the cover colour, but now contains information to extract an approximation of the secret pixel value, which gets padded with 0 to fill in the missing bits, so comes out to (64, 192, 96) == (01000000, 11000000, 01100000)    .

Using a larger number of hidden bits results in a better quality hidden image, but makes it easier to tell that the hidden image is there. Play with the 'hidden bits' slider to see.

There is an example on Wikipedia of a cat hidden in a picture of a tree. To do the example here, use the 'Unhide image' tool to select the 'Wikipedia tree' example, and set hidden bits to 2.

All of the computation is performed in your browser in Javascript. Be careful with larger images as it can be quite slow.

Once you've hidden your secret image inside a cover image, send the output image to your accomplice. Your accomplice then uses the 'Unhide image' tool to recover the secret image.