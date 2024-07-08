package routers

import (
	"GoBBS/controller"
	"GoBBS/logger"
	"GoBBS/middlewares"
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
	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})
	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"msg": "404",
		})
	})
	return r
}
