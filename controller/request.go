package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
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
