package imagecompress

import (
	"errors"
	"fmt"
	imagedownloader "rocket/helper/imageDownloader"
	db "rocket/helper/sql"
	"rocket/structure"
	"strings"
	"sync"
	"time"
)

func CompressImages(productID string) error {
	var product structure.Product
	time.Sleep(2 * time.Second)
	errUpload := make(chan error)
	sqlString := fmt.Sprintf("SELECT * from products where product_id = \"%v\" ", productID)
	err := db.Dbmap.SelectOne(&product, sqlString)
	if err != nil {
		return errors.New("no user found")
	}
	images := strings.Split(product.ProductImages, ",")
	var wg sync.WaitGroup
	for i := 0; i < len(images); i++ {
		wg.Add(1)
		go imagedownloader.DownloadImage(images[i], &wg, errUpload)
		errorUploading := <-errUpload
		if errorUploading != nil {
			fmt.Println("sds", errorUploading)
			close(errUpload)
			return errorUploading
		}
	}
	close(errUpload)
	wg.Wait()
	return nil
}
