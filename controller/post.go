package controller

import (
	"jachow/code1024/logic"
	"jachow/code1024/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	// TODO: 调用logic层创建帖子
	params := new(model.Post)
	if err := c.ShouldBindJSON(params); err != nil {
		zap.L().Error("CreatePostHandler bind json failed", zap.Error(err))
		Response(c, CodeInvalidParam, http.StatusBadRequest, err.Error())
		return
	}
	// TODO: 调用logic层创建帖子

	// 通过上下文获取当前登录用户的ID
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("get current user id failed", zap.Error(err))
		Response(c, CodeInvalidParam, http.StatusBadRequest, err.Error())
		return
	}
	params.AuthorID = userID

	if err := logic.CreatePost(params); err != nil {
		zap.L().Error("create post failed", zap.Error(err))
		Response(c, CodeServerBusy, http.StatusInternalServerError, err.Error())
		return
	}
	Response(c, CodeSuccess, http.StatusOK, nil)
}

func GETPostHandler(c *gin.Context) {
	// TODO: 调用logic层获取帖子
	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		zap.L().Error("GETPostHandler parse post id failed", zap.Error(err))
		Response(c, CodeInvalidParam, http.StatusBadRequest, err.Error())
		return
	}
	// TODO: 调用logic层获取帖子
	post, err := logic.GetPost(postID)
	if err != nil {
		zap.L().Error("GETPostHandler get post failed", zap.Error(err))
		Response(c, CodeServerBusy, http.StatusInternalServerError, err.Error())
		return
	}
	Response(c, CodeSuccess, http.StatusOK, post)
}

func GetPostListHandler(c *gin.Context) {
	// TODO: 调用logic层获取帖子列表
	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "10")
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		zap.L().Error("GetPostListHandler parse offset failed", zap.Error(err))
		Response(c, CodeInvalidParam, http.StatusBadRequest, err.Error())
		return
	}
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		zap.L().Error("GetPostListHandler parse limit failed", zap.Error(err))
		Response(c, CodeInvalidParam, http.StatusBadRequest, err.Error())
		return
	}
	ApiPostList, err := logic.GetPostList(offset, limit)
	if err != nil {
		zap.L().Error("GetPostListHandler get post list failed", zap.Error(err))
		Response(c, CodeServerBusy, http.StatusInternalServerError, err.Error())
		return
	}
	Response(c, CodeSuccess, http.StatusOK, ApiPostList)
}

func VotePostHandler(c *gin.Context) {
	// TODO: 调用logic层投票
	params := new(model.ParamsVote)
	if err := c.ShouldBindJSON(params); err != nil {
		zap.L().Error("VotePostHandler bind json failed", zap.Error(err))
		Response(c, CodeInvalidParam, http.StatusBadRequest, err.Error())
		return
	}
	// TODO: 调用logic层投票
	// 通过上下文获取当前登录用户的ID
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("get current user id failed", zap.Error(err))
		Response(c, CodeInvalidParam, http.StatusBadRequest, err.Error())
		return
	}
	if err := logic.VotePost(userID, params); err != nil {
		zap.L().Error("VotePostHandler vote post failed", zap.Error(err))
		Response(c, CodeServerBusy, http.StatusInternalServerError, err.Error())
		return
	}
	Response(c, CodeSuccess, http.StatusOK, nil)
}

func GetPostListOrderByHandler(c *gin.Context) {
	// TODO: 调用logic层获取帖子列表
	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "10")
	orderBy := c.DefaultQuery("order", "score")
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		zap.L().Error("GetPostListOrderByHandler parse offset failed", zap.Error(err))
		Response(c, CodeInvalidParam, http.StatusBadRequest, err.Error())
		return
	}
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		zap.L().Error("GetPostListOrderByHandler parse limit failed", zap.Error(err))
		Response(c, CodeInvalidParam, http.StatusBadRequest, err.Error())
		return
	}
	ApiPostList, err := logic.GetPostListOrderBy(offset, limit)
	if err != nil {
		zap.L().Error("GetPostListOrderByHandler get post list failed", zap.Error(err))
		Response(c, CodeServerBusy, http.StatusInternalServerError, err.Error())
		return
	}
	Response(c, CodeSuccess, http.StatusOK, ApiPostList)
}
