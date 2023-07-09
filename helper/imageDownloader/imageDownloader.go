package imagedownloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

func DownloadImage(imageLink string, wg *sync.WaitGroup, errUpload chan error) {
	defer wg.Done()
	filePath := "image.jpg" // Replace with the desired file path to save the image
	response, err := http.Get(imageLink)
	if err != nil {
		fmt.Println("Error:", err)
		errUpload <- err
		return
	}
	defer response.Body.Close()
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		errUpload <- err
		return
	}
	defer file.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		errUpload <- err
		return
	}
	fmt.Println("Image downloaded successfully.")
	return
}
