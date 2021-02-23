import cv2
import numpy as np

def fixPixel(color):

	b, g, r = color

	if r <= g and r <= b:
		if r == 0:
			r += 1
		else:
			r -= 1

	elif g <= r and g <= b:
		if g == 0:
			g += 1
		else:
			g -= 1

	else:
		if b == 0:
			b += 1
		else:
			b -= 1

	return (b, g, r)


def imgEmbed(img, waterprint):

	h = img.shape[0]
	w = img.shape[1]
	bound_h = waterprint.shape[0]
	bound_w = waterprint.shape[1]
	outImg = np.zeros((h, w, 3), np.uint8)

	for y in range(0, h):
		for x in range(0, w):

			b, g, r = img[y, x].astype(np.uint16)

			if y < bound_w and x < bound_h:

				chkb, chkg, chkr = waterprint[y, x].astype(np.uint16)

				if ((chkr+chkg+chkb)%2 == 0 and (r+g+b)%2 != 0) \
					or ((chkr+chkg+chkb)%2 != 0 and (r+g+b)%2 == 0):

					outImg[y, x] = fixPixel(img[y, x])

				else:

					outImg[y, x] = img[y, x]
			else:

				outImg[y, x] = img[y, x]

	return outImg


def imgExtract(embed):

	h = embed.shape[0]
	w = embed.shape[1]
	outImg = np.zeros((h, w, 3), np.uint8)

	for y in range(0, h):
		for x in range(0, w):

			b, g, r = embed[y, x].astype(np.uint16)

			if (r+g+b)%2 != 0:

				outImg[y, x] = (255, 255, 255)
			else:

				outImg[y, x] = (0, 0, 0)

	return outImg



if __name__ == '__main__':
	# read images
	img = cv2.imread('./testImage/cat_1200x600.png')
	waterprint = cv2.imread('./testImage/NISRA.png')

	# show images
	cv2.imshow('origin image', img)
	cv2.waitKey(0)
	cv2.imshow('waterprint', waterprint)
	cv2.waitKey(0)

	# embed images
	imgEmbedded = imgEmbed(img, waterprint)
	cv2.imshow('extract', imgEmbedded)
	cv2.waitKey(0)
	cv2.imwrite('./testImage/output/embedded.png', imgEmbedded)

	# extract image
	decode = imgExtract(imgEmbedded)
	cv2.imshow('merge', decode)
	cv2.waitKey(0)
	cv2.imwrite('./testImage/output/decode.png', decode)
	