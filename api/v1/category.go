package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_mall/pkg/utils"
	"web_mall/service"
)

// GetCategory 获取商品列表
func GetCategory(c *gin.Context) {
	var categoryService service.CategoryService
	if err := c.ShouldBind(&categoryService); err == nil {
		res := categoryService.ListCategory(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln(err)
	}
}
