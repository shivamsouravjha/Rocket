package routes

import (
	post "rocket/controllers/POST"

	"github.com/gin-gonic/gin"
)

func v0Routes(route *gin.RouterGroup) {
	v1Routes := route.Group("/products")
	{
		v1Routes.GET("/create", post.AddProduct)
	}
}
