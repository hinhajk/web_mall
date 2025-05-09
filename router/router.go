package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_mall/api/v1"
	"web_mall/middleware"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.Cors())                    //使用跨域中间件
	router.StaticFS("/static", http.Dir("./static")) //加载静态文件的路径
	v1 := router.Group("/api/v1")
	{
		//ping一下，保证服务能够连通
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})

		//用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		//轮播图
		v1.GET("carousels", api.ListCarousel)

		//商品操作
		v1.GET("/product/lists", api.GetProductList) // 获取商品列表

		//在验证token情况下进行的操作
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.PUT("user/update", api.UserUpdate)
			authed.PUT("user/avatar", api.UploadAvatar)      //上传头像
			authed.POST("user/sending-email", api.SendEmail) //邮件发送
			authed.POST("user/valid-email", api.ValidEmail)  //验证邮箱

			authed.POST("money", api.ShowMoney) //显示用户金额

			authed.POST("/product/create", api.CreateProduct) //创建商品

		}
	}
	return router
}
