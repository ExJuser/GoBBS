package logic

import (
	"GoBBS/dao/mysql"
	"GoBBS/models"
	"GoBBS/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(post *models.Post) (err error) {
	//生成post的id
	post.ID = snowflake.GenID()
	return mysql.CreatePost(post)
}

func GetPostByID(postID int64) (data *models.APIPostDetail, err error) {
	//查询并拼接数据
	post := &models.Post{}
	community := &models.CommunityDetail{}
	user := &models.User{}
	if post, err = mysql.GetPostByID(postID); err != nil {
		zap.L().Error("mysql.GetPostByID(postID) failed", zap.Int64("postID", postID), zap.Error(err))
		return nil, err
	}
	if community, err = mysql.GetCommunityDetailByID(post.CommunityID); err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed", zap.Int64("post.CommunityID", post.CommunityID), zap.Error(err))
		return nil, err
	}
	if user, err = mysql.GetUserByID(post.AuthorID); err != nil {
		zap.L().Error("mysql.GetUserByID(post.AuthorID)", zap.Int64("post.AuthorID", post.AuthorID), zap.Error(err))
		return nil, err
	}
	data = &models.APIPostDetail{
		AuthorName: user.Username,
		Post:       post,
		Community:  community,
	}
	return
}
