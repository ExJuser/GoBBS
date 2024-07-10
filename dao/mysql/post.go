package mysql

import (
	"GoBBS/models"
	"strconv"
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
	tx := db.Raw("select post_id,author_id,community_id,status, title,content,create_time from post order by create_time desc limit ?,?", (page-1)*size, size).Scan(&posts)
	return posts, tx.Error
}

func GetPostListByIDs(postIDs []string) (posts []*models.Post, err error) {
	posts = make([]*models.Post, 0, len(postIDs))
	for _, postID := range postIDs {
		var postIDInt64 int64
		postIDInt64, err = strconv.ParseInt(postID, 10, 64)
		if err != nil {
			continue
		}
		var post *models.Post
		post, err = GetPostByID(postIDInt64)
		if err != nil {
			continue
		}
		posts = append(posts, post)
	}
	return
}
