package redis

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated   = errors.New("不允许重复投票")
)

func CreatePost(postID int64, createTime time.Time, communityID int64) error {
	pipeline := rdb.TxPipeline()
	ctx := context.Background()
	pipeline.ZAdd(ctx, getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(createTime.Unix()),
		Member: postID,
	})

	pipeline.ZAdd(ctx, getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(createTime.Unix()),
		Member: postID,
	})
	cKey := getRedisKey(KeyCommunitySetPF + strconv.FormatInt(communityID, 10))
	pipeline.SAdd(ctx, cKey, postID)
	_, err := pipeline.Exec(ctx)
	return err
}

func VoteForPost(userID, postID int64, direction int8) (err error) {
	ctx := context.Background()
	postIDStr := strconv.FormatInt(postID, 10)
	userIDStr := strconv.FormatInt(userID, 10)
	postTime := rdb.ZScore(ctx, getRedisKey(KeyPostTimeZSet), postIDStr).Val()
	//帖子发布时间超过一周 不可点赞
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	//查询之前有没有投过票 没有为0 赞成为1
	ov := rdb.ZScore(ctx, getRedisKey(KeyPostVotedZSetPF+postIDStr), userIDStr).Val()
	pointsDiff := (float64(direction) - ov) * scorePerVote
	if pointsDiff == 0 {
		return ErrVoteRepeated
	}
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(ctx, getRedisKey(KeyPostScoreZSet), pointsDiff, postIDStr)
	if direction == 0 {
		pipeline.ZRem(ctx, getRedisKey(KeyPostVotedZSetPF+postIDStr), userID)
	} else {
		pipeline.ZAdd(ctx, getRedisKey(KeyPostVotedZSetPF+postIDStr), redis.Z{
			Score:  float64(direction), //赞成还是反对
			Member: userID,
		})
	}
	_, err = pipeline.Exec(ctx)
	return
}
