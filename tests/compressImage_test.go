package test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	imagecompress "rocket/helper/imageCompress"
)

func TestCompressImages(t *testing.T) {
	db, err := sql.Open("mysql", "root:newpassword@tcp(localhost:3306)/productio")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Arrange
	productID := "208"

	// Act
	err = imagecompress.CompressImages(productID)

	fmt.Println(err)
	// Assert
	if err != nil {
		t.Errorf("Error compressing images: %v", err)
	}

	// Check that the compressed images were stored in the database
	rows, err := db.QueryContext(context.Background(), "SELECT compressed_product_images FROM products WHERE product_id = ?", productID)
	if err != nil {
		t.Errorf("Error querying database: %v", err)
	}
	defer rows.Close()

	if rows.Next() {
		var compressedImages string
		err := rows.Scan(&compressedImages)
		if err != nil {
			t.Errorf("Error scanning row: %v", err)
		}
		// Assert that the compressed images are correct
		if compressedImages != "../images/208_0_img.jpg" {
			t.Errorf("Expected compressed images to be %s, got %s", "../images/150_0_img.jpg", compressedImages)
		}
	} else {
		t.Errorf("No compressed images found for product ID %s", productID)
	}
}
