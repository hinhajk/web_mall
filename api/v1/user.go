package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_mall/pkg/utils"
	"web_mall/service"
)

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	var userRegister service.UserService //定义请求参数
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register(c.Request.Context()) //将上下文传递进去
		c.JSON(http.StatusOK, res)
	} else {
		//fmt.Println("1111")
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln("user register api:", err)
	}
}

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln("user login api:", err)
	}
}

// UserUpdate 用户更新
func UserUpdate(c *gin.Context) {
	var userUpdate service.UserService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdate); err == nil {
		res := userUpdate.Update(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln("user update api err:", err)
	}
}

// UploadAvatar 上传头像
func UploadAvatar(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	filesSize := fileHeader.Size
	var userAvatar service.UserService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userAvatar); err == nil {
		res := userAvatar.LoadAvatar(c.Request.Context(), claims.ID, file, filesSize)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln("upload avatar api err:", err)
	}
}

// SendEmail 发送邮箱
func SendEmail(c *gin.Context) {
	var sendEmail service.SendEmailService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&sendEmail); err == nil {
		res := sendEmail.Send(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln("send email api err:", err)
	}
}

// ValidEmail 验证邮箱
func ValidEmail(c *gin.Context) {
	var validEmail service.ValidEmailService
	if err := c.ShouldBind(&validEmail); err == nil {
		res := validEmail.Valid(c.Request.Context(), c.GetHeader("Authorization"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln("valid email api err:", err)
	}
}

// ShowMoney 显示用户金额
func ShowMoney(c *gin.Context) {
	var showMoney service.ShowMoneyService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showMoney); err == nil {
		res := showMoney.Show(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogObj.Infoln("show money api err:", err)
	}
}
