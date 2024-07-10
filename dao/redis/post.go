package redis

import (
	"GoBBS/models"
	"context"
	"github.com/redis/go-redis/v9"
)

func GetPostIDInOrder(p *models.ParamPostList) ([]string, error) {
	var key string
	ctx := context.Background()
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	} else {
		key = getRedisKey(KeyPostTimeZSet)
	}
	start := (p.Page - 1) * p.Size
	stop := start + p.Size - 1
	return rdb.ZRevRange(ctx, key, start, stop).Result()
}

func GetPostVoteData(postIDs []string) (data []int64, err error) {
	ctx := context.Background()
	data = make([]int64, 0, len(postIDs))
	pipeline := rdb.Pipeline()
	for _, postID := range postIDs {
		pipeline.ZCount(ctx, getRedisKey(KeyPostVotedZSetPF+postID), "1", "1")
	}
	cmders, err := pipeline.Exec(ctx)
	if err != nil {
		return nil, err
	}
	for _, cmder := range cmders {
		voteCount := cmder.(*redis.IntCmd).Val()
		data = append(data, voteCount)
	}
	return
}
