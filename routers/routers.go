package routers

import (
	"GoBBS/controller"
	"GoBBS/logger"
	"GoBBS/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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
	r.GET("/ping", JWTAuthMiddleware(), func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})
	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"msg": "404",
		})
	})
	return r
}

func JWTAuthMiddleware() func(context *gin.Context) {
	return func(context *gin.Context) {
		//客户端携带token有三种方式：放在请求体、放在请求体、放在url
		authHeader := context.Request.Header.Get("Authorization")
		//没有携带token
		if authHeader == "" {
			context.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求体重auth为空",
			})
			context.Abort()
			return
		}
		//token格式不对
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			context.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求体中auth格式有误",
			})
			context.Abort()
			return
		}
		//token无效
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			context.Abort()
			return
		}
		context.Set("userID", mc.UserID)
		context.Next()
	}
}
