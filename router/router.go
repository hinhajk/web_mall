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
		v1.GET("product/lists", api.GetProductList)   // 获取商品列表
		v1.GET("product/show/:id", api.ShowProduct)   //获取商品详细信息
		v1.GET("product/imgs/:id", api.ShowImgs)      //获取商品图片信息
		v1.GET("product/categories", api.GetCategory) //获取商品分类

		//在验证token情况下进行的操作
		authed := v1.Group("/") //需要登录保护
		authed.Use(middleware.JWT())
		{
			authed.PUT("user/update", api.UserUpdate)
			authed.PUT("user/avatar", api.UploadAvatar)      //上传头像
			authed.POST("user/sending-email", api.SendEmail) //邮件发送
			authed.POST("user/valid-email", api.ValidEmail)  //验证邮箱

			authed.POST("money", api.ShowMoney) //显示用户金额

			//商品操作
			authed.POST("product/create", api.CreateProduct)          //创建商品
			authed.POST("product/search", api.SearchProduct)          //搜索商品
			authed.DELETE("product/delete/:id", api.DeleteProduct)    //删除商品
			authed.PUT("product/update/:id", api.UpdateProduct)       //更新商品信息（除图片外)
			authed.PUT("product/update-avatar/:id", api.UpdateAvatar) //更新商品图片信息

			//收藏夹操作
			authed.GET("favorites", api.ShowFavorites)          //查看收藏夹
			authed.POST("favorites", api.CreateFavorites)       //创建收藏夹
			authed.DELETE("favorites/:id", api.DeleteFavorites) //删除收藏夹

		}
	}
	return router
}
