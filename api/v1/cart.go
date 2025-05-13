package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_mall/pkg/utils"
	"web_mall/service"
)

// CreateCart 创建地址
func CreateCart(c *gin.Context) {
	var createCart service.CartService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createCart); err == nil {
		res := createCart.Create(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln(err)
	}
}

// ShowCart 查询购物车
func ShowCart(c *gin.Context) {
	var listCart service.CartService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listCart); err == nil {
		res := listCart.Show(c.Request.Context(), claims.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
		//fmt.Println(11111)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln(err)
	}
}

// UpdateCart 更新购物车
func UpdateCart(c *gin.Context) {
	var updateCart service.CartService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&updateCart); err == nil {
		res := updateCart.Update(c.Request.Context(), claims.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln(err)
	}
}

// DeleteCart 移出购物车
func DeleteCart(c *gin.Context) {
	var deleteCart service.CartService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteCart); err == nil {
		res := deleteCart.Delete(c.Request.Context(), claims.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln(err)
	}
}

// DeleteCarts 清空购物车
func DeleteCarts(c *gin.Context) {
	var deleteCart service.CartService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteCart); err == nil {
		res := deleteCart.DeleteAll(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln(err)
	}
}
