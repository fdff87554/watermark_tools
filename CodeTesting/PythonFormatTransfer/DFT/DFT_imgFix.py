import numpy as np
import cv2
import matplotlib.pyplot as plt

# Press the green button in the gutter to run the script.
if __name__ == '__main__':

    img = cv2.imread('./turing.jpeg', 0)
    freq = np.fft.fft2(img)
    freq_shift = np.fft.fftshift(freq)
    magnitude_spectrum = 20 * np.log(np.abs(freq_shift))

    rotate_img = cv2.imread('./turing_rotate.jpeg', 0)
    rotate_freq = np.fft.fft2(rotate_img)
    rotate_freq_shift = np.fft.fftshift(rotate_freq)
    rotate_magnitude_spectrum = 20 * np.log(np.abs(rotate_freq_shift))

    # 將 array 畫成圖像
    plt.figure(figsize=(20, 10))
    plt.subplot(221), plt.imshow(img, cmap='gray')
    plt.title('Input Image'), plt.xticks([]), plt.yticks([])
    plt.subplot(222), plt.imshow(magnitude_spectrum, cmap='gray')  # 將頻譜畫成圖，方便解讀頻率域資訊
    plt.title('Magnitude Spectrum'), plt.xticks([]), plt.yticks([])
    plt.subplot(223), plt.imshow(rotate_img, cmap='gray')
    plt.title('Input rotate Image'), plt.xticks([]), plt.yticks([])
    plt.subplot(224), plt.imshow(rotate_magnitude_spectrum, cmap='gray')  # 將頻譜畫成圖，方便解讀頻率域資訊
    plt.title('Rotate Magnitude Spectrum'), plt.xticks([]), plt.yticks([])

    # 儲存圖像
    plt.savefig('./Result.png', bbox_inches='tight', dpi=300)