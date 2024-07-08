package mysql

import (
	"GoBBS/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"gorm.io/gorm"
)

//把每一步数据库操作封装成函数 等待logic业务层根据业务需求调用

var (
	ErrUserExist       = errors.New("用户已存在")
	ErrUserNotExist    = errors.New("用户名不存在")
	ErrInvalidPassword = errors.New("密码错误")
)

const secret = "zhuchenchen.top"

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	var userCount int
	if tx := db.Raw("select count(user_id) from user where username=?", username).Find(&userCount); tx.Error != nil {
		return tx.Error
	}
	if userCount > 0 {
		return ErrUserExist
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

func Login(user *models.User) (err error) {
	oPassword := user.Password //用户传入的明文密码
	if tx := db.Where("username", user.Username).First(user); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return ErrUserNotExist //用户不存在错误
		}
		//其他数据库错误
		return tx.Error
	}
	password := encryptPassword(oPassword)
	//user.Password 从数据库查出来的加密后的密码
	if password != user.Password {
		return ErrInvalidPassword //密码不正确的错误
	}
	return
}

func encryptPassword(oPassword string) string {
	hash := md5.New()
	hash.Write([]byte(secret))
	return hex.EncodeToString(hash.Sum([]byte(oPassword)))
}
