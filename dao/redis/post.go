package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func VotePost(userID, postID string, value float64) (err error) {
	ctx := context.Background()
	_, err = Client.ZScore(ctx, RedisKeyPostTime, postID).Result()
	if err != nil {
		zap.L().Error("redis zscore failed", zap.Error(err))
		return ErrVoteInvalid
	}

	// 检查用户是否已投票
	preValue, err := Client.ZScore(ctx, RedisKeyPostVote+postID, userID).Result()
	if err != nil {
		zap.L().Error("redis zscore failed", zap.Error(err))
		return ErrVoteInvalid
	}

	// 检查投票方向
	if value == preValue {
		zap.L().Error("user has voted, can't vote again")
		return ErrVoteAgain
	}

	// 更新userID对于postID的投票方向
	Client.ZAdd(ctx, RedisKeyPostVote+postID, redis.Z{
		Score:  value,
		Member: userID,
	})

	DeltaScore := 432 * (value - preValue)

	// 获取当前post分数
	preScore, err := Client.ZScore(ctx, RedisKeyPostScore, postID).Result()
	if err != nil {
		zap.L().Error("redis zscore failed", zap.Error(err))
		return ErrVoteInvalid
	}

	// 更新post分数
	Client.ZAdd(ctx, RedisKeyPostScore, redis.Z{
		Score:  preScore + DeltaScore,
		Member: postID,
	})

	return err
}
