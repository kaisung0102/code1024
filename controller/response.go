package controller

import (
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(c *gin.Context, code ResCode, httpcode int, data interface{}) {
	c.JSON(httpcode, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: data,
	})
}
