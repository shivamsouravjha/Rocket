package structure

type Product struct {
	ProductID          int     `db:"product_id" binding:"required"`
	ProductName        string  `db:"product_name" binding:"required"`
	ProductDescription string  `db:"product_description" binding:"required"`
	ProductImages      string  `db:"product_images" binding:"required"`
	ProductPrice       float64 `db:"product_price" binding:"required"`
	CompressedImages   string  `db:"compressed_product_images" binding:"required"`
	CreatedAt          string  `db:"created_at" binding:"required"`
	UpdatedAt          string  `db:"updated_at" binding:"required"`
}
