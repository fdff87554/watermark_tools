import numpy as np
import cv2
import matplotlib.pyplot as plt

def LPF_processing(img, cover):
    """
    將圖片過濾高頻部分
    :param img: the img of original img
    :param cover: the cover base, which will define the block that remove the lower freq
    :return lpf_img: the freq with mask
    :return spectrum: the freq image
    """

    h, w = img.shape[:2]
    mask = np.zeros((h, w), np.uint8)
    mask[0:int(h / cover), 0:int(w / cover)] = 1  # 建立遮罩，長寬是原本影像的 1 / cover，位置是以左上角為起點
    lpf_img = img * mask  # 遮罩的地方乘1，也就讓低頻的部分通過(高頻部分都會因為乘上0而被捨棄掉)

    spectrum = 20 * np.log(np.abs(lpf_img))  # 將經過LPF之後的影像轉為頻譜
    return lpf_img, spectrum

if __name__ == '__main__':
    # image read
    img = cv2.imread('./Turing.jpeg', 0)

    # DCT
    dct_img = cv2.dct(img.astype(float))  # 將原圖轉換成 float32 格式才能使用 opencv 的 dct function
    dct_magnitude_spectrum = 20 * np.log(np.abs(dct_img))

    # LPF 實作
    LPF_img, LPF_dct_magnitude_spectrum = LPF_processing(dct_img, 5)

    # 將經過 LPF 後的頻譜轉回圖像
    idct_img = cv2.idct(LPF_img)

    # 將 array 畫成圖像
    plt.figure(figsize=(5, 6))  # 設定輸出影像長寬

    plt.subplot(221), plt.imshow(img, cmap='gray')  # 原圖
    plt.title('Input Image'), plt.xticks([]), plt.yticks([])

    plt.subplot(222), plt.imshow(dct_magnitude_spectrum, cmap='gray')  # DCT 頻譜圖
    plt.title('Spectrum after DCT'), plt.xticks([]), plt.yticks([])

    plt.subplot(223), plt.imshow(LPF_dct_magnitude_spectrum, cmap='gray')  # 透過遮罩做 LPF 的頻譜圖
    plt.title('DCT after LPF'), plt.xticks([]), plt.yticks([])

    plt.subplot(224), plt.imshow(idct_img.astype(np.uint8), cmap='gray')  # 轉回圖片
    plt.title('Image after LPF (IDCT)'), plt.xticks([]), plt.yticks([])

    # 儲存圖像
    plt.savefig('./Result.png', bbox_inches='tight', dpi=300)