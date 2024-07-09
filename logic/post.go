package logic

import (
	"GoBBS/dao/mysql"
	"GoBBS/models"
	"GoBBS/pkg/snowflake"
)

func CreatePost(post *models.Post) (err error) {
	//生成post的id
	post.ID = snowflake.GenID()
	return mysql.CreatePost(post)
}

func GetPostByID(postID int64) (data *models.Post, err error) {
	return mysql.GetPostByID(postID)
}
