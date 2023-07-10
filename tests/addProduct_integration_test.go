package test

import (
	"context"
	product "rocket/helper/db"
	requestStruct "rocket/structure/request"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddProductIntegration(t *testing.T) {
	productData := requestStruct.PostProduct{
		UserID:             1,
		ProductName:        "SampleProductsData",
		ProductDescription: "This is a test product",
		ProductPrice:       10.00,
		ProductImages:      []string{"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTHZAq08u4YaR0Jsu2CgeptdxC74y-9QEeFYEAb6YHP&s"},
	}
	err := product.AddProduct(context.Background(), productData)
	assert.NoError(t, err)
}
