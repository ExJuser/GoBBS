package controller

import (
	"GoBBS/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SignUpHandler 处理注册请求的方法
func SignUpHandler(context *gin.Context) {
	//参数校验

	//业务处理
	logic.SignUp()
	//返回响应
	context.JSON(http.StatusOK, "ok")
}
