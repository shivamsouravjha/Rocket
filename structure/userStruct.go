package structure

import "time"

type User struct {
	ID        int       `json:"id" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Mobile    string    `json:"mobile" binding:"required"`
	Latitude  float64   `json:"latitude" binding:"required"`
	Longitude float64   `json:"longitude" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
	UpdatedAt time.Time `json:"updated_at" binding:"required"`
}
