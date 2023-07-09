package db

import (
	"context"
	"errors"
	"fmt"
	rabbitMQ "rocket/helper/rabbitmq"
	"time"

	db "rocket/helper/sql"
	requestStruct "rocket/structure/request"
	"strings"
)

func AddProduct(ctx context.Context, userProductData requestStruct.PostProduct) error {
	var matchingSearch int
	sqlString := fmt.Sprintf("SELECT count(*) from user where id=  \"%v\" ", userProductData.UserID)
	err := db.Dbmap.SelectOne(&matchingSearch, sqlString)
	if err != nil || matchingSearch <= 0 {
		return errors.New("no user found")
	}
	images := strings.Join(userProductData.ProductImages, ",")
	createdProduct, err := db.Dbmap.Exec("INSERT INTO products (product_name, product_description,  product_price, product_images,created_at,updated_at,compressed_product_images) VALUES(?,?,?,?,?,?,?) ", userProductData.ProductName, userProductData.ProductDescription, userProductData.ProductPrice, images, time.Now(), time.Now(), "")
	if err != nil {
		return err
	}
	productID, _ := createdProduct.LastInsertId()
	rabbitMQ.RunPublish("image", fmt.Sprint(productID))
	return nil
}
