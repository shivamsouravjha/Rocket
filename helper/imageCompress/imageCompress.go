package imagecompress

import (
	"errors"
	"fmt"
	imagedownloader "rocket/helper/imageDownloader"
	db "rocket/helper/sql"
	"rocket/structure"
	"strings"
	"sync"
)

func CompressImages(productID string) error {
	var product structure.Product
	errUpload := make(chan error, 1)
	paths := []string{}
	sqlString := fmt.Sprintf("SELECT * from products where product_id = \"%v\" ", productID)
	err := db.Dbmap.SelectOne(&product, sqlString)
	if err != nil {
		return errors.New("no user found")
	}
	images := strings.Split(product.ProductImages, ",")
	var wg sync.WaitGroup
	for i := 0; i < len(images); i++ {
		wg.Add(1)
		fmt.Println(images)
		go func(index int, wg *sync.WaitGroup) {
			defer wg.Done()
			path, err := imagedownloader.DownloadImage(images[index], productID, index)
			if err != nil {
				fmt.Println("Error downloading image:", err)
				errUpload <- err
				return
			}
			paths = append(paths, path)
		}(i, &wg)
	}
	wg.Wait()
	close(errUpload)

	if <-errUpload != nil {
		return errors.New("failed to upload image(s)")

	}
	compressImages := strings.Join(paths, ",")
	sqlString = fmt.Sprintf("UPDATE products set compressed_product_images = \"%v\" where product_id = \"%v\"", compressImages, productID)
	_, err = db.Dbmap.Exec(sqlString)
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}
