package middlewares

import (
	"jachow/code1024/controller"
	"jachow/code1024/pkg"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取token
		AuthorHeader := c.Request.Header.Get("Authorization")
		if AuthorHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": controller.CodeNeedLogin,
				"msg":  controller.CodeNeedLogin.Msg(),
			})
			c.Abort()
			return
		}
		parts := strings.SplitN(AuthorHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusOK, gin.H{
				"code": controller.CodeInvalidToken,
				"msg":  controller.CodeInvalidToken.Msg(),
			})
			c.Abort()
			return
		}
		token := parts[1]
		claims, err := pkg.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": controller.CodeInvalidToken,
				"msg":  controller.CodeInvalidToken.Msg(),
			})
			c.Abort()
			return
		}
		c.Set(controller.UserID, claims.UserID)
		c.Next()
	}
}
