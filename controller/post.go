package controller

import (
	"GoBBS/logic"
	"GoBBS/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
)

func CreatePostHandler(context *gin.Context) {
	post := models.Post{}
	//传入的参数就存在错误
	if err := context.ShouldBindJSON(&post); err != nil {
		zap.L().Debug("context.ShouldBindJSON(&post)", zap.Error(err))
		zap.L().Error("create post with invalid param")
		ResponseError(context, CodeInvalidParam)
		return
	}
	//从context中取到用户ID
	userID, err := getCurrentUserID(context)
	if err != nil {
		ResponseError(context, CodeNeedLogin)
		return
	}
	post.AuthorID = userID
	post.CreateTime = time.Now()
	//数据库层面的错误
	if err := logic.CreatePost(&post); err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	ResponseSuccess(context, nil)
}

func GetPostDetailHandler(context *gin.Context) {
	postIDStr := context.Param("id")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		zap.L().Error("get post with invalid param", zap.Error(err))
		ResponseError(context, CodeInvalidPassword)
		return
	}
	data, err := logic.GetPostByID(postID)
	if err != nil {
		zap.L().Error("logic.GetPostByID(postID) failed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	ResponseSuccess(context, data)
}

func GetPostListHandler(context *gin.Context) {
	page, size := getPageInfo(context)
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	ResponseSuccess(context, data)
}
