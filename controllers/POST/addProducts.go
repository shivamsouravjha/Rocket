package post

import (
	"net/http"
	requestStruct "rocket/structure/request"
	responseStruct "rocket/structure/response"

	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {
	addProductStruct := requestStruct.PostProduct{}
	if err := c.ShouldBind(&addProductStruct); err != nil {
		c.JSON(422, "Error")
		return
	}
	resp := responseStruct.SuccessResponse{}
	resp.Message = "Product added successfully"
	resp.Status = "true"
	c.JSON(http.StatusOK, resp)
}
