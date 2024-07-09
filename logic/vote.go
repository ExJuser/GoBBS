package logic

import (
	"GoBBS/dao/redis"
	"GoBBS/models"
	"go.uber.org/zap"
)

func VoteForPost(userID int64, v *models.ParamVoteData) error {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.Int64("postID", v.PostID),
		zap.Int8("Direction", v.Direction))
	return redis.VoteForPost(userID, v.PostID, v.Direction)
}
