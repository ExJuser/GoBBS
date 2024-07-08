package routers

import (
	"GoBBS/controller"
	"GoBBS/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//注册
	r.POST("/signup", controller.SignUpHandler)
	//登录
	r.POST("/login", controller.LoginHandler)

	//测试
	r.GET("/ping", func(context *gin.Context) {
		if isLogin() {
			//如果是登陆的用户
			context.String(http.StatusOK, "pong")
		} else {
			//否则直接返回请登录
			context.String(http.StatusOK, "请登录")
		}
	})
	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"msg": "404",
		})
	})
	return r
}
