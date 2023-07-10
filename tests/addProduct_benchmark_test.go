package test

import (
	"context"
	"fmt"
	product "rocket/helper/db"
	requestStruct "rocket/structure/request"
	"testing"
)

func BenchmarkAddProduct(b *testing.B) {
	productData := requestStruct.PostProduct{
		UserID:             1,
		ProductName:        "SampleProductsData",
		ProductDescription: "This is a test product",
		ProductPrice:       10.00,
		ProductImages:      []string{"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTHZAq08u4YaR0Jsu2CgeptdxC74y-9QEeFYEAb6YHP&s"},
	}
	for i := 0; i < b.N; i++ {
		err := product.AddProduct(context.Background(), productData)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

//APP_ENV=test go test -run=10000 -bench=BenchmarkAddProduct -benchmem
