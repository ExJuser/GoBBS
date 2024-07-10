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

// GetPostListHandler2 根据前端传来的参数（按分数、按创建时间等）动态获取帖子列表
func GetPostListHandler2(context *gin.Context) {
	//去redis拿到id列表 再到数据库查询帖子详细信息
	p := models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	if err := context.ShouldBindQuery(&p); err != nil {
		zap.L().Error("context.ShouldBindQuery(&p) failed", zap.Error(err))
		ResponseError(context, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostList2(&p)
	if err != nil {
		zap.L().Error("logic.GetPostList2(&p) failed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	ResponseSuccess(context, data)
}

func GetCommunityPostListHandler(context *gin.Context) {
	p := models.ParamCommunityPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	if err := context.ShouldBindQuery(&p); err != nil {
		zap.L().Error("context.ShouldBindQuery(&p) failed", zap.Error(err))
		ResponseError(context, CodeInvalidParam)
		return
	}
	data, err := logic.GetCommunityPostList(&p)
	if err != nil {
		zap.L().Error("logic.GetCommunityPostList failed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	ResponseSuccess(context, data)
}
