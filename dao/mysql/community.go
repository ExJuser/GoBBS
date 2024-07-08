package mysql

import (
	"GoBBS/models"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
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

func GetCommunityDetailByID(id int64) (detail *models.CommunityDetail, err error) {
	tx := db.Raw("select community_id,community_name,introduction,create_time from community where community_id=?", id).First(&detail)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) { //没有查询到记录 ID出错
			err = ErrInvalidID
			return nil, err
		}
		//其他数据库错误
		err = tx.Error
	}
	return
}
