package router

import (
	"jachow/code1024/controller"
	"jachow/code1024/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Routers() *gin.Engine {
	r := gin.Default()
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
	}

	r.NoRoute(func(c *gin.Context) {
		zap.L().Error("404 Not Found", zap.String("path", c.Request.URL.Path))
		controller.Response(c, controller.CodeNotFound, http.StatusNotFound, nil)
	})

	return r
}
