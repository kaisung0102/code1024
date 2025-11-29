package controller

import (
	"jachow/code1024/logic"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	// TODO: 调用logic层获取所有社区
	CommunityList, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("get community failed", zap.Error(err))
		Response(c, CodeServerBusy, http.StatusInternalServerError, err.Error())
		return
	}
	Response(c, CodeSuccess, http.StatusOK, CommunityList)
}

func CommunityDetailHandler(c *gin.Context) {
	// TODO: 调用logic层获取指定社区详情
	communityID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		zap.L().Error("get community detail failed", zap.Error(err))
		Response(c, CodeInvalidParam, http.StatusBadRequest, err.Error())
		return
	}
	community, err := logic.GetCommunityByID(communityID)
	if err != nil {
		zap.L().Error("get community detail failed", zap.Error(err))
		Response(c, CodeServerBusy, http.StatusInternalServerError, err.Error())
		return
	}
	Response(c, CodeSuccess, http.StatusOK, community)
}
