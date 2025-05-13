package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_mall/pkg/utils"
	"web_mall/service"
)

// CreateAddress 创建地址
func CreateAddress(c *gin.Context) {
	var createAddress service.AddressService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createAddress); err == nil {
		res := createAddress.Create(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln(err)
	}
}

// GetAddress 获取单个地址
func GetAddress(c *gin.Context) {
	var createAddress service.AddressService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createAddress); err == nil {
		res := createAddress.Get(c.Request.Context(), claims.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln(err)
	}
}

// ListAddress 获取用户所有地址
func ListAddress(c *gin.Context) {
	var createAddress service.AddressService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createAddress); err == nil {
		res := createAddress.List(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln(err)
	}
}

// UpdateAddress 更新单个地址
func UpdateAddress(c *gin.Context) {
	var createAddress service.AddressService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createAddress); err == nil {
		res := createAddress.Update(c.Request.Context(), claims.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln(err)
	}
}

// DeleteAddress 删除某个地址
func DeleteAddress(c *gin.Context) {
	var createAddress service.AddressService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createAddress); err == nil {
		res := createAddress.Delete(c.Request.Context(), claims.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln(err)
	}
}
