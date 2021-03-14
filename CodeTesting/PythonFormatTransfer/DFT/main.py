from PIL import Image
import numpy as np
import matplotlib.pyplot as plt

# def HPF_processing(img, cover):
#     """
#     將圖片過濾低頻部分
#     :param img: the img of original img
#     :param cover: the cover pixels, which will define the block that remove the lower freq
#     :return: hpf_img
#     """
#
#     # 利用原圖的 長寬 取得中心座標
#     rows, cols = img.shape
#     crow, ccol = int(rows/2), int(cols/2)
#
#     # 利用 np 的 fft 進行傅立葉轉換，由於是 image，所以用 fft2 進行二維傅立葉轉換
#     freq = np.fft.fft2(img)
#     # 頻率最低的地方會左上角，將結果平移至中心，比較好理解
#     freq_shift = np.fft.fftshift(freq)
#
#     # 利用中心點各取 cover/2 來形成一個 cover*cover 的遮罩直接過濾掉中心低頻
#     freq_shift[int(crow-(cover/2)):int(crow+(cover/2)), int(ccol-(cover/2)):int(ccol+(cover/2))] = 0
#
#     # 恢復頻率原來的狀況
#     freq_rshift = np.fft.ifftshift(freq_shift)
#     # 回轉頻率域 to 圖域
#     hpf_img = np.fft.ifft2(freq_rshift)
#
#     return np.abs(hpf_img)
#
# def LPF_processing(img, cover):
#     """
#     將圖片過濾高頻部分
#     :param img: the img of original img
#     :param cover: the cover pixels, which will define the block that remove the lower freq
#     :return: lpf_img
#     """
#
#     # 利用原圖的 長寬 取得中心座標
#     rows, cols = img.shape
#     crow, ccol = int(rows / 2), int(cols / 2)
#
#     # 利用 np 的 fft 進行傅立葉轉換，由於是 image，所以用 fft2 進行二維傅立葉轉換
#     freq = np.fft.fft2(img)
#     # 頻率最低的地方會左上角，將結果平移至中心，比較好理解
#     freq_shift = np.fft.fftshift(freq)
#
#     # 利用中心點各取 cover/2 來形成一個 cover*cover 的遮罩讓低頻通過
#     mask = np.zeros((rows, cols), np.uint8)
#     mask[int(crow-(cover/2)):int(crow+(cover/2)), int(ccol-(cover/2)):int(ccol+(cover/2))] = 1
#     freq_shift = freq_shift*mask
#     freq_rshift = np.fft.ifftshift(freq_shift)
#     lpf_img = np.fft.ifft2(freq_rshift)
#
#     return np.abs(lpf_img)


# Press the green button in the gutter to run the script.
if __name__ == '__main__':

    with Image.open('./turing.jpeg') as img:
        img = img.convert('L')

    freq = np.fft.fft2(img)
    freq_shift = np.fft.fftshift(freq)
    magnitude_spectrum = 20 * np.log(np.abs(freq_shift))

    # # 高通濾波器實作
    # HPF_img = HPF_processing(img, 30)
    #
    # # 低通濾波器實作
    # LPF_img = LPF_processing(img, 30)

    # 將 array 畫成圖像
    plt.figure(figsize=(20, 10))
    plt.subplot(221), plt.imshow(img, cmap='gray')
    plt.title('Input Image'), plt.xticks([]), plt.yticks([])
    plt.subplot(222), plt.imshow(magnitude_spectrum, cmap='gray')  # 將頻譜畫成圖，方便解讀頻率域資訊
    plt.title('Magnitude Spectrum'), plt.xticks([]), plt.yticks([])
    plt.subplot(223), plt.imshow(HPF_img, cmap='gray')
    plt.title('Image after HPF'), plt.xticks([]), plt.yticks([])
    plt.subplot(224), plt.imshow(LPF_img, cmap='gray')
    plt.title('Image after LPF'), plt.xticks([]), plt.yticks([])

    # 儲存圖像
    plt.savefig('./Result.png', bbox_inches='tight', dpi=300)
