package redis

import (
	"GoBBS/models"
	"context"
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
