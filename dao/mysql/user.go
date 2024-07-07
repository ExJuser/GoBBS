package mysql

import (
	"GoBBS/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

//把每一步数据库操作封装成函数 等待logic业务层根据业务需求调用

const secret = "zhuchenchen.top"

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	var userCount int
	if tx := db.Raw("select count(user_id) from user where username=?", username).Find(&userCount); tx.Error != nil {
		return tx.Error
	}
	if userCount > 0 {
		return errors.New("用户已存在")
	}
	return
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	//对密码进行加密
	user.Password = encryptPassword(user.Password)
	//执行SQL语句入库
	return db.Exec("Insert into user(user_id, username, password) values (?,?,?)", user.UserID, user.Username, user.Password).Error
}

func encryptPassword(oPassword string) string {
	hash := md5.New()
	hash.Write([]byte(secret))
	return hex.EncodeToString(hash.Sum([]byte(oPassword)))
}
