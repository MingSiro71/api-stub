package redis_repositories

import (
	"context"
	"github.com/go-redis/redis/v9"
	"api_stub/repositories"
	"api_stub/vo"
)

type redisMessageRepository struct {
	ctx context.Context
	db *redis.Client
}

func NewRedisMessageRepository(ctx context.Context, db *redis.Client) repositories.MessageRepository{
	return &redisMessageRepository{ctx: ctx, db: db}
}

func (repo *redisMessageRepository) Push(id vo.Id, s string) error {
	err := repo.db.RPush(repo.ctx, id.Tos(), s).Err()
	return err
}

func (repo *redisMessageRepository) Pop(id vo.Id) (string, error) {
	val, err := repo.db.LPop(repo.ctx, id.Tos()).Result()
	return string(val), err
}