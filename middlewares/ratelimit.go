package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	r2 "github.com/juju/ratelimit"
	r1 "go.uber.org/ratelimit"
)

func RateLimit1Middleware() gin.HandlerFunc {
	// 100次/秒
	limiter := r1.New(100, r1.Per(time.Second))
	return func(c *gin.Context) {
		// 从limiter中获取一个令牌
		if time.Until(limiter.Take()) > 0 { // 底层是 原子类的CAS操作
			c.JSON(429, gin.H{
				"message": "rate limit exceeded",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func RateLimit2Middleware() gin.HandlerFunc {
	// 100次/秒
	limiter := r2.NewBucket(time.Second, 100)
	return func(c *gin.Context) {
		// 从limiter中获取一个令牌
		if limiter.TakeAvailable(1) != 1 {
			c.JSON(429, gin.H{
				"message": "rate limit exceeded",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
