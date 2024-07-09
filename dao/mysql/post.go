package mysql

import (
	"GoBBS/models"
)

func CreatePost(post *models.Post) (err error) {
	err = db.Create(post).Error
	return
}
