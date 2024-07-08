package controller

import (
	"GoBBS/logic"
	"GoBBS/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
)

// SignUpHandler 处理注册请求的方法
func SignUpHandler(context *gin.Context) {
	//参数获取与校验
	p := new(models.ParamSignUp)
	if err := context.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		var errs validator.ValidationErrors
		if ok := errors.As(err, &errs); !ok {
			context.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			//可选项：去掉提示信息中的前缀 ParamSignUp
			//"ParamSignUp.re_password": "re_password为必填字段"
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	//业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("signup failed,err:", zap.Error(err))
		context.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	//返回响应
	context.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func LoginHandler(context *gin.Context) {
	//获取请求参数和参数校验
	p := new(models.ParamLogin)
	if err := context.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		var errs validator.ValidationErrors
		if ok := errors.As(err, &errs); !ok {
			context.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
	}
	//业务逻辑处理
	if err := logic.Login(p); err != nil {
		zap.L().Error("login failed,err:", zap.String("username", p.Username), zap.Error(err))
		context.JSON(http.StatusOK, gin.H{
			"msg": "用户名或密码错误",
		})
		return
	}
	//返回响应
	context.JSON(http.StatusOK, gin.H{
		"msg": "登陆成功",
	})
}
