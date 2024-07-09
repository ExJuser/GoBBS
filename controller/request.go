package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

const CtxUserIDKey = "userID"

var ErrUserNotLogin = errors.New("用户未登录")

func getCurrentUserID(ctx *gin.Context) (userID int64, err error) {
	uid, ok := ctx.Get(CtxUserIDKey)
	if !ok {
		err = ErrUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrUserNotLogin
		return
	}
	return
}

func getPageInfo(context *gin.Context) (int64, int64) {
	//获取分页参数
	pageNumStr := context.Query("page")
	pageSizeStr := context.Query("size")
	page, err := strconv.ParseInt(pageNumStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
