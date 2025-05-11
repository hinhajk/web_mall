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
		utils.LogObj.Infoln(err)
	}
}

// GetProductList 获取商品列表
func GetProductList(c *gin.Context) {
	var productList service.ProductService
	if err := c.ShouldBind(&productList); err == nil {
		res := productList.ProductList(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln(err)
	}
}

// SearchProduct 搜索商品
func SearchProduct(c *gin.Context) {
	var searchProduct service.ProductService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&searchProduct); err == nil {
		res := searchProduct.Search(c.Request.Context(), claims.ID)
		c.JSON(http.StatusBadRequest, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln(err)
	}
}

// ShowProduct 获取商品详细信息
func ShowProduct(c *gin.Context) {
	var showProduct service.ProductService
	if err := c.ShouldBind(&showProduct); err == nil {
		res := showProduct.Show(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln(err)
	}
}

// DeleteProduct 搜索商品
func DeleteProduct(c *gin.Context) {
	var deleteProduct service.ProductService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteProduct); err == nil {
		res := deleteProduct.Delete(c.Request.Context(), claims.ID, c.Param("id"))
		c.JSON(http.StatusBadRequest, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln(err)
	}
}

// UpdateProduct 搜索商品
func UpdateProduct(c *gin.Context) {
	var updateProduct service.ProductService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&updateProduct); err == nil {
		res := updateProduct.Update(c.Request.Context(), claims.ID, c.Param("id"))
		c.JSON(http.StatusBadRequest, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln(err)
	}
}

// UpdateAvatar 更新商品图片信息
func UpdateAvatar(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	var createProduct service.ProductService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createProduct); err == nil {
		res := createProduct.UpdateAvatar(c.Request.Context(), claims.ID, c.Param("id"), files)
		c.JSON(http.StatusBadRequest, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln(err)
	}
}
