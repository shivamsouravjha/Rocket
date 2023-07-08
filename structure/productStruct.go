package structure

import "time"

type Product struct {
	ProductID          int
	ProductName        string
	ProductDescription string
	ProductPrice       float64
	CompressedImages   string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
