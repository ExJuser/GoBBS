package logic

import (
	"GoBBS/dao/mysql"
	"GoBBS/pkg/snowflake"
)

//存放业务逻辑代码 可能多次调用dao层服务

func SignUp() {
	//判断用户是否存在
	mysql.QueryUserByUsername()
	//生成UID
	snowflake.GenID()
	//保存进数据库（用户密码加密存储）
	mysql.InsertUser()
}
