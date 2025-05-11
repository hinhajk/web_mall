package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"web_mall/pkg/utils"
	"web_mall/service"
)

func ShowImgs(c *gin.Context) {
	var listProductImg service.ListProductImg
	if err := c.ShouldBind(&listProductImg); err == nil {
		res := listProductImg.ListImgs(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
