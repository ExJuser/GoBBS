package logic

import (
	"GoBBS/dao/mysql"
	"GoBBS/models"
)

func GetCommunityList() ([]models.Community, error) {
	//查数据库查找到所有的community并返回
	return mysql.GetCommunityList()
}
