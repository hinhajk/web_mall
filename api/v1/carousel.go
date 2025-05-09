package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"web_mall/pkg/utils"
	"web_mall/service"
)

func ListCarousel(c *gin.Context) {
	var listCarousel service.Carousel
	if err := c.ShouldBind(&listCarousel); err == nil {
		res := listCarousel.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
