package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const UserID = "userID"

func getCurrentUserID(c *gin.Context) (int64, error) {
	// 从上下文获取当前登录用户的ID
	userID, exists := c.Get(UserID)
	if !exists {
		return 0, errors.New("userID not found in context")
	}
	return userID.(int64), nil
}
