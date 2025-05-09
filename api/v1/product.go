package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_mall/pkg/utils"
	"web_mall/service"
)

// CreateProduct 创建商品
func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	var createProduct service.ProductService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createProduct); err == nil {
		res := createProduct.Create(c.Request.Context(), claims.ID, files)
		c.JSON(http.StatusBadRequest, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
