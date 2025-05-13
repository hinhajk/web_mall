package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_mall/pkg/utils"
	"web_mall/service"
)

// CreateFavorites 创建收藏夹
func CreateFavorites(c *gin.Context) {
	var createFavorites service.FavoritesService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createFavorites); err == nil {
		res := createFavorites.Create(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln(err)
	}
}

// DeleteFavorites 删除收藏夹
func DeleteFavorites(c *gin.Context) {
	var deleteFavorites service.FavoritesService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteFavorites); err == nil {
		res := deleteFavorites.Delete(c.Request.Context(), claims.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln(err)
	}
}

// ShowFavorites 查看收藏夹
func ShowFavorites(c *gin.Context) {
	var showFavorites service.FavoritesService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showFavorites); err == nil {
		res := showFavorites.Show(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln(err)
	}
}
