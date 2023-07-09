package imagedownloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadImage(imageLink string, productID string, img int) (string, error) {
	filePath := fmt.Sprintf("images/%s_%d_img.jpg", productID, img)
	response, err := http.Get(imageLink)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	defer response.Body.Close()
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	defer file.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	fmt.Println("Image downloaded successfully:", filePath)
	return filePath, nil
}
