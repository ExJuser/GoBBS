package mysql

import (
	"GoBBS/models"
)

func CreatePost(post *models.Post) (err error) {
	err = db.Create(post).Error
	return
}

func GetPostByID(postID int64) (*models.Post, error) {
	post := &models.Post{}
	tx := db.Where("post_id=?", postID).First(&post)
	return post, tx.Error
}

func GetPostList(page, size int64) ([]*models.Post, error) {
	posts := make([]*models.Post, 0, size)
	tx := db.Raw("select * from post limit ?,?", (page-1)*size, size).Scan(&posts)
	return posts, tx.Error
}
