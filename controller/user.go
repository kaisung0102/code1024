package controller

import (
	"jachow/code1024/dao/mysql"
	"jachow/code1024/logic"
	"jachow/code1024/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	params := new(model.ParamsSignUp)
	if err := c.ShouldBindJSON(params); err != nil {
		zap.L().Error("SignUpHandler bind json failed", zap.Error(err))
		Response(c, CodeInvalidParam, http.StatusBadRequest, err.Error())
		return
	}

	// TODO: 调用logic层注册用户
	if err := logic.SignUp(params); err != nil {
		zap.L().Error("sign up failed", zap.Error(err))
		// 用户名已存在
		if err == mysql.ErrUserExist {
			Response(c, CodeUserExist, http.StatusBadRequest, err.Error())
			return
		}
		Response(c, CodeInvalidParam, http.StatusBadRequest, err.Error())
		return
	}

	Response(c, CodeSuccess, http.StatusOK, nil)
}

func LoginHandler(c *gin.Context) {
	params := new(model.ParamsLogin)
	if err := c.ShouldBindJSON(params); err != nil {
		zap.L().Error("LoginHandler bind json failed", zap.Error(err))
		Response(c, CodeInvalidParam, http.StatusBadRequest, err.Error())
		return
	}

	// TODO: 调用logic层登录用户
	token, err := logic.Login(params)
	if err != nil {
		zap.L().Error("login failed", zap.String("username", params.Username), zap.Error(err))
		// 用户名不存在
		if err == mysql.ErrUserNotExist {
			Response(c, CodeUserNotExist, http.StatusBadRequest, err.Error())
			return
		}
		// 密码错误
		Response(c, CodeInvalidPassword, http.StatusBadRequest, err.Error())
		return
	}

	Response(c, CodeSuccess, http.StatusOK, token)
}
