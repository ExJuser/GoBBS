package middlewares

import (
	"GoBBS/controller"
	"GoBBS/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

func JWTAuthMiddleware() func(context *gin.Context) {
	return func(context *gin.Context) {
		//客户端携带token有三种方式：放在请求体、放在请求体、放在url
		authHeader := context.Request.Header.Get("Authorization")
		//没有携带token
		if authHeader == "" {
			controller.ResponseError(context, controller.CodeNeedLogin)
			context.Abort()
			return
		}
		//token格式不对
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(context, controller.CodeInvalidToken)
			context.Abort()
			return
		}
		//token无效
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(context, controller.CodeInvalidToken)
			context.Abort()
			return
		}
		//在请求上下文context中储存userID
		context.Set(controller.CtxUserIDKey, mc.UserID)
		context.Next() //后续的处理函数可以通过context.Get("userID")获取当前用户信息
	}
}
