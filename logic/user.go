package logic

import (
	"GoBBS/dao/mysql"
	"GoBBS/models"
	"GoBBS/pkg/snowflake"
)

//存放业务逻辑代码 可能多次调用dao层服务

func SignUp(p *models.ParamSignUp) (err error) {
	//判断用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return
	}
	//生成UID
	userID := snowflake.GenID()
	//构造一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//保存进数据库（用户密码加密存储）
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.Login(user)
}
