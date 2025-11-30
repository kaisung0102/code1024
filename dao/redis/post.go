package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var scorePerVote = 432.0

func VotePost(userID, postID string, value float64) (err error) {
	ctx := context.Background()
	_, err = Client.ZScore(ctx, RedisKeyPostTime, postID).Result()
	if err != nil {
		zap.L().Error("redis zscore failed", zap.Error(err))
		return ErrVoteInvalid
	}

	// 检查用户是否已投票
	preValue, _ := Client.ZScore(ctx, RedisKeyPostVote+postID, userID).Result()

	// 检查投票方向
	if value == preValue {
		zap.L().Error("user has voted, can't vote again")
		return ErrVoteAgain
	}

	tx := Client.TxPipeline()
	// 更新userID对于postID的投票方向
	tx.ZAdd(ctx, RedisKeyPostVote+postID, redis.Z{
		Score:  value,
		Member: userID,
	})

	// 更新post分数
	tx.ZIncrBy(ctx, RedisKeyPostScore,
		scorePerVote*(value-preValue),
		postID,
	)

	_, err = tx.Exec(ctx)
	if err != nil {
		zap.L().Error("redis tx exec failed", zap.Error(err))
		return ErrVoteInvalid
	}

	return err
}

func CreatePost(postID string, createTime int64) (err error) {
	ctx := context.Background()
	// 初始化帖子分数
	tx := Client.TxPipeline()
	tx.ZAdd(ctx, RedisKeyPostScore, redis.Z{
		Score:  0,
		Member: postID,
	})
	// 初始化帖子创建时间
	tx.ZAdd(ctx, RedisKeyPostTime, redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 执行事务
	_, err = tx.Exec(ctx)
	if err != nil {
		zap.L().Error("redis tx exec failed", zap.Error(err))
		return ErrVoteInvalid
	}
	return err
}
