package core

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
)

type RedisCore struct {
	client *redis.Client
	logger *zap.SugaredLogger
}

func NewRedisCore(client *redis.Client, logger *zap.SugaredLogger) *RedisCore {
	return &RedisCore{client: client, logger: logger}
}

func (c *RedisCore) ReadSessionID(userID string) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	return c.client.Get(ctx, userID).Result()
}
