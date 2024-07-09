package redis

const (
	KeyPrefix = "gobbs:"
	// KeyPostTimeZSet 记录帖子发布时间
	KeyPostTimeZSet = "post:time"
	// KeyPostScoreZSet 记录帖子的分数
	KeyPostScoreZSet = "post:score"
	// KeyPostVotedZSetPF 记录一个帖子都有谁投过票 需要继续拼接post_id
	KeyPostVotedZSetPF = "post:voted:"
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
