# import numpy as np
# import scipy as misc
import imageio
import matplotlib.pyplot as plt

if __name__ == '__main__':
    img = imageio.imread('./test.jpeg')
    print(img.shape, img.dtype)

    plt.imshow(img)
    plt.show()
    plt.imshow(img[:, :, 0])
    plt.show()
    plt.imshow(img[:, :, 1])
    plt.show()
    plt.imshow(img[:, :, 2])
    plt.show()