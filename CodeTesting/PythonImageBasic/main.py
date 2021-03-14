from PIL import Image

import numpy as np
import matplotlib.pyplot as plt

if __name__ == '__main__':

    with Image.open('./images/input/cover.png') as cover:
        cover = cover.convert('RGB')

    # Test pil image tool
    cover.save('./images/output/pil_out.png')
    with Image.open('./images/output/pil_out.png') as p_cover:
        p_cover = p_cover.convert('RGB')

    cover = np.asarray(cover)
    p_cover = np.asarray(p_cover)
    # print(p_cover - cover)
    c_h, c_w, _ = cover.shape

    # Test fft exchange
    freq = np.fft.fft2(cover)
    img = np.fft.ifft2(freq)
    print(np.abs(img))
    print(img / (c_h*c_w))
    # img_freq = np.fft.fft2(np.abs(img))

    # print("int", np.abs(img).astype(int))
    # print("uint8", np.abs(img).astype(np.uint8))

    # print(np.abs(img).astype(np.uint8) - cover)
    # plt.imshow(np.abs(img).astype(np.uint8) - cover)
    # plt.show()
    # im = Image.fromarray(np.abs(img).astype(np.uint8))
    # im.save('./images/output/ifft_img.png')
    #
    # with Image.open('./images/output/ifft_img.png') as ifft_cover:
    #     ifft_cover = ifft_cover.convert('RGB')