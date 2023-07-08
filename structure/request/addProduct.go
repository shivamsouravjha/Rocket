package requestStruct

type PostProduct struct {
	UserID             int      `json:"user_id" binding:"required"`
	ProductName        string   `json:"product_name" binding:"required"`
	ProductDescription string   `json:"product_description" binding:"required"`
	ProductImages      []string `json:"product_images" binding:"required"`
	ProductPrice       float64  `json:"product_price" binding:"required"`
}
