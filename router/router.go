package router

import (
	"net/http"

	v1 "github.com/binz96/blog/api/v1"
	"github.com/binz96/blog/middleware"
	"github.com/binz96/blog/setting"
	"github.com/gin-gonic/gin"
)

func Run() {
	gin.SetMode(setting.GinMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	rv1 := r.Group("/api/v1")
	{
		// GET http://localhost:8000/api/v1/hello
		rv1.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code":    "0",
				"message": "ok",
				"data":    "hello",
			})
		})

		// 用户模块路由接口
		rv1.POST("user/add", v1.CreateUser)
		rv1.GET("users", middleware.JwtToken(), v1.GetUsers)
		rv1.PUT("user/:id", middleware.JwtToken(), v1.UpdateUser)
		rv1.DELETE("user/:id", middleware.JwtToken(), v1.DeleteUser)

		// 分类模块路由接口
		rv1.POST("category/add", middleware.JwtToken(), v1.CreateCategory)
		rv1.GET("categories", middleware.JwtToken(), v1.GetCategories)
		rv1.GET("category/:cid", v1.GetArticlesInCategory)
		rv1.PUT("category/:id", middleware.JwtToken(), v1.UpdateCategory)
		rv1.DELETE("category/:id", middleware.JwtToken(), v1.DeleteCategory)

		// 文章模块路由接口
		rv1.POST("article/add", middleware.JwtToken(), v1.CreateArticle)
		rv1.GET("article/:id", v1.GetArticle)
		rv1.GET("articles", v1.GetArticles)
		rv1.PUT("article/:id", middleware.JwtToken(), v1.UpdateArticle)
		rv1.DELETE("article/:id", middleware.JwtToken(), v1.DeleteArticle)

		// 登录
		rv1.POST("login", v1.Login)

		rv1.POST("upload", middleware.JwtToken(), v1.UploadFile)
	}

	r.Run(":" + setting.HttpPort)
}
