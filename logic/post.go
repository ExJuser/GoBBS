package logic

import (
	"GoBBS/dao/mysql"
	"GoBBS/dao/redis"
	"GoBBS/models"
	"GoBBS/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(post *models.Post) (err error) {
	//生成post的id
	post.ID = snowflake.GenID()
	if err = mysql.CreatePost(post); err != nil {
		return
	}
	return redis.CreatePost(post.ID, post.CreateTime)
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

func GetPostList(page, size int64) (postsDetails []*models.APIPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return
	}
	postsDetails = make([]*models.APIPostDetail, 0, len(posts))
	for _, post := range posts {
		community := &models.CommunityDetail{}
		user := &models.User{}
		if community, err = mysql.GetCommunityDetailByID(post.CommunityID); err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed", zap.Int64("post.CommunityID", post.CommunityID), zap.Error(err))
			continue
		}
		if user, err = mysql.GetUserByID(post.AuthorID); err != nil {
			zap.L().Error("mysql.GetUserByID(post.AuthorID)", zap.Int64("post.AuthorID", post.AuthorID), zap.Error(err))
			continue
		}
		postsDetails = append(postsDetails, &models.APIPostDetail{
			AuthorName: user.Username,
			Post:       post,
			Community:  community,
		})
	}
	return
}

func GetPostList2(p *models.ParamPostList) (postsDetails []*models.APIPostDetail, err error) {
	postIDs, err := redis.GetPostIDInOrder(p)
	if err != nil {
		return
	}
	if len(postIDs) == 0 {
		zap.L().Warn("redis.GetPostIDInOrder(p) returns 0 data")
		return
	}
	posts, err := mysql.GetPostListByIDs(postIDs)
	voteData, err := redis.GetPostVoteData(postIDs)
	if err != nil {
		return
	}
	for i, post := range posts {
		community := &models.CommunityDetail{}
		user := &models.User{}
		if community, err = mysql.GetCommunityDetailByID(post.CommunityID); err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed", zap.Int64("post.CommunityID", post.CommunityID), zap.Error(err))
			continue
		}
		if user, err = mysql.GetUserByID(post.AuthorID); err != nil {
			zap.L().Error("mysql.GetUserByID(post.AuthorID)", zap.Int64("post.AuthorID", post.AuthorID), zap.Error(err))
			continue
		}
		postsDetails = append(postsDetails, &models.APIPostDetail{
			AuthorName: user.Username,
			VoteCount:  voteData[i],
			Post:       post,
			Community:  community,
		})
	}
	return
}
