package internal

import (
	"fmt"
	imagecompress "rocket/helper/imageCompress"
)

func StoreProduct(product []byte) {
	err := imagecompress.CompressImages(string(product))
	if err != nil {
		fmt.Println("error", err.Error())
	}
}
