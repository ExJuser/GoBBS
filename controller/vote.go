package controller

import (
	"GoBBS/logic"
	"GoBBS/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func VoteForPostHandler(context *gin.Context) {
	v := models.ParamVoteData{}
	if err := context.ShouldBindJSON(&v); err != nil {
		var errs validator.ValidationErrors
		if ok := errors.As(err, &errs); ok {
			ResponseErrorWithMsg(context, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		} else {
			ResponseError(context, CodeInvalidParam)
		}
		return
	}
	userID, err := getCurrentUserID(context)
	if err != nil {
		ResponseError(context, CodeNeedLogin)
		return
	}
	if err = logic.VoteForPost(userID, &v); err != nil {
		zap.L().Error("logic.VoteForPost(userID, &v) failed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	ResponseSuccess(context, nil)
}
