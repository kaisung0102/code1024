package redis

import (
	"context"
	"jachow/code1024/config"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func InitRedis(redisConfig *config.RedisConfig) {
	// TODO: 初始化Redis
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	// 测试连接
	_, err := client.Ping(ctx).Result()
	if err != nil {
		zap.L().Error("redis ping failed", zap.Error(err))
		return
	}
}
