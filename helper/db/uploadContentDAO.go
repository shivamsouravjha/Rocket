package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	rabbitMQ "rocket/helper/rabbitmq"
	db "rocket/helper/sql"
	"rocket/structure"
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
	product := structure.Product{
		ProductName:        userProductData.ProductName,
		ProductDescription: userProductData.ProductDescription,
		CompressedImages:   images,
		ProductPrice:       userProductData.ProductPrice,
	}
	insertedSeriesIdPush, _ := json.Marshal(product)
	rabbitMQ.RunPublish("SeriesContent", string(insertedSeriesIdPush))
	return nil
}
