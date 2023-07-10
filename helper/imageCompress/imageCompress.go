package imagecompress

import (
	"errors"
	"fmt"
	"sync"

	imagedownloader "rocket/helper/imageDownloader"
	db "rocket/helper/sql"
	"rocket/structure"
	"strings"
)

func CompressImages(productID string) error {
	var product structure.Product
	sqlString := fmt.Sprintf("SELECT * from products where product_id = \"%v\" ", productID)
	err := db.Dbmap.SelectOne(&product, sqlString)
	if err != nil {
		return errors.New("no product found")
	}

	images := strings.Split(product.ProductImages, ",")
	numImages := len(images)

	errUpload := make(chan error, numImages)
	paths := make([]string, numImages)

	var wg sync.WaitGroup
	wg.Add(numImages)

	for i, image := range images {
		go func(index int, image string) {
			defer wg.Done()
			fmt.Println(index)
			path, err := imagedownloader.DownloadImage(image, productID, index)
			if err != nil {
				errUpload <- err
				return
			}
			paths[index] = path
			fmt.Println(paths[index], "sds")
		}(i, image)
	}

	wg.Wait()
	close(errUpload)

	for err := range errUpload {
		fmt.Println(err)
	}
	if len(errUpload) > 0 {
		return errors.New("failed to upload image(s)")
	}

	compressedImages := strings.Join(paths, ",")
	sqlString = fmt.Sprintf("UPDATE products SET compressed_product_images = \"%v\" WHERE product_id = \"%v\"", compressedImages, productID)
	_, err = db.Dbmap.Exec(sqlString)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
