package redis

import (
	"GoBBS/models"
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

func getIDsFromKey(key string, page, size int64) ([]string, error) {
	ctx := context.Background()
	start := (page - 1) * size
	stop := start + size - 1
	return rdb.ZRevRange(ctx, key, start, stop).Result()
}

func GetPostIDInOrder(p *models.ParamPostList) ([]string, error) {
	var key string

	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	} else {
		key = getRedisKey(KeyPostTimeZSet)
	}
	return getIDsFromKey(key, p.Page, p.Size)
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

func GetCommunityPostIDInOrder(p *models.ParamPostList) ([]string, error) {
	ctx := context.Background()
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}
	key := orderKey + strconv.FormatInt(p.CommunityID, 10)
	cKey := getRedisKey(KeyCommunitySetPF + strconv.FormatInt(p.CommunityID, 10))
	if rdb.Exists(ctx, key).Val() < 1 {
		pipeline := rdb.Pipeline()
		pipeline.ZInterStore(ctx, key, &redis.ZStore{
			Aggregate: "MAX",
			Keys:      []string{cKey, orderKey},
		})
		pipeline.Expire(ctx, key, time.Minute)
		_, err := pipeline.Exec(ctx)
		if err != nil {
			return nil, err
		}
	}
	return getIDsFromKey(key, p.Page, p.Size)
}
