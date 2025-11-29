package router

import (
	"jachow/code1024/controller"
	"jachow/code1024/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	docs "jachow/code1024/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routers() *gin.Engine {
	docs.SwaggerInfo.BasePath = "/api/v1"
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v := r.Group("/api/v1")
	v.POST("/signup", controller.SignUpHandler)

	v.POST("/login", controller.LoginHandler)

	v.Use(middlewares.JWTAuthMiddleware())
	{
		v.GET("/community", controller.CommunityHandler)
		v.GET("/community/:id", controller.CommunityDetailHandler)
		v.POST("/post", controller.CreatePostHandler)
		v.GET("/post/:id", controller.GETPostHandler)
		v.GET("/posts", controller.GetPostListHandler)
		v.POST("/vote", controller.VotePostHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		zap.L().Error("404 Not Found", zap.String("path", c.Request.URL.Path))
		controller.Response(c, controller.CodeNotFound, http.StatusNotFound, nil)
	})

	return r
}
