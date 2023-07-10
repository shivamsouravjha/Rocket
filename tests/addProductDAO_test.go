package db

import (
	"context"
	"database/sql"
	product "rocket/helper/db"
	requestStruct "rocket/structure/request"
	"strings"
	"testing"
)

func TestAddProduct(t *testing.T) {
	db, err := sql.Open("mysql", "root:newpassword@tcp(localhost:3306)/productio")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	productData := requestStruct.PostProduct{
		UserID:             1,
		ProductName:        "SampleProduct",
		ProductDescription: "This is a test product",
		ProductPrice:       10.00,
		ProductImages:      []string{"https://example.com/image.jpg"},
	}

	ctx := context.Background()
	err = product.AddProduct(ctx, productData)

	if err != nil {
		t.Errorf("Error adding product: %v", err)
	}

	rows, err := db.QueryContext(ctx, "SELECT * FROM products WHERE product_name = ?", productData.ProductName)
	if err != nil {
		t.Errorf("Error querying database: %v", err)
	}
	defer rows.Close()

	if rows.Next() {
		var productID int
		var productName string
		var productDescription string
		var productPrice float64
		var productImages string
		var createdAt string
		var updatedAt string
		var compressed_product_images string
		err := rows.Scan(&productID, &productName, &productDescription, &productPrice, &compressed_product_images, &productImages, &createdAt, &updatedAt)
		if err != nil {
			t.Errorf("Error scanning row: %v", err)
		}
		if productName != productData.ProductName {
			t.Errorf("Expected product name to be %s, got %s", productData.ProductName, productName)
		}
		if productDescription != productData.ProductDescription {
			t.Errorf("Expected product description to be %s, got %s", productData.ProductDescription, productDescription)
		}
		if productPrice != productData.ProductPrice {
			t.Errorf("Expected product price to be %f, got %f", productData.ProductPrice, productPrice)
		}
		if productImages != strings.Join(productData.ProductImages, ",") {
			t.Errorf("Expected product images to be %s, got %s", productData.ProductImages, productImages)
		}
	} else {
		t.Errorf("No product found with product name %s", productData.ProductName)
	}
}
