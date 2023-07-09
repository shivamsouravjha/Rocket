package imagecompress

import (
	"errors"
	"fmt"
	db "rocket/helper/sql"
	"rocket/structure"
	"time"
)

func CompressImages(productID string) error {
	var product structure.Product
	time.Sleep(2 * time.Second)
	sqlString := fmt.Sprintf("SELECT * from products where product_id = \"%v\" ", productID)
	err := db.Dbmap.SelectOne(&product, sqlString)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("no user found")
	}
	return nil
}
