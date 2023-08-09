package post

import (
	"net/http"
	requestStruct "rocket/structure/request"
	responseStruct "rocket/structure/response"
	"rocket/utils"

	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {
	addProductStruct := requestStruct.PostProduct{}
	if err := c.ShouldBind(&addProductStruct); err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	// err := db.AddProduct(c, addProductStruct)
	// if err != nil {
	// 	resp := responseStruct.SuccessResponse{}
	// 	resp.Message = err.Error()
	// 	resp.Status = "false"
	// 	c.JSON(http.StatusInternalServerError, resp)
	// 	return
	// }
	resp := responseStruct.SuccessResponse{}
	resp.Message = "Product added successfully"
	resp.Status = "true"
	c.JSON(http.StatusOK, resp)
}
