package mysql

import (
	"GoBBS/models"
	"go.uber.org/zap"
)

func GetCommunityList() (communities []models.Community, err error) {
	tx := db.Raw("select community_id,community_name from community").Scan(&communities)
	if tx.Error != nil { //数据库查询出错
		err = tx.Error
		return
	}
	if tx.RowsAffected == 0 { //没有数据
		zap.L().Warn("there is no community")
		return
	}
	return
}
