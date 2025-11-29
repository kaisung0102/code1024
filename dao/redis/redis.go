package redis

import (
	"context"
	"errors"
	"jachow/code1024/config"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

const (
	RedisKeyPrefix    = "code1024:"
	RedisKeyPostTime  = RedisKeyPrefix + "post:time"
	RedisKeyPostScore = RedisKeyPrefix + "post:score"
	RedisKeyPostVote  = RedisKeyPrefix + "post:vote:"
)

var (
	ErrVoteAgain   = errors.New("不能重复投票")
	ErrVoteInvalid = errors.New("投票无效")
)

func InitRedis(redisConfig *config.RedisConfig) {
	// TODO: 初始化Redis

	Client = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(redisConfig.Timeout)*time.Second)
	defer cancel()
	// 测试连接
	_, err := Client.Ping(ctx).Result()
	if err != nil {
		zap.L().Error("redis ping failed", zap.Error(err))
		return
	}
}

var (
	// ctx    = context.Background()
	Client *redis.Client
)