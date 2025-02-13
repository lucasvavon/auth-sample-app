package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisSessionRepository struct {
	client *redis.Client
}

func NewRedisSessionRepository() *RedisSessionRepository {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return &RedisSessionRepository{client: client}
}

func (r *RedisSessionRepository) SaveSession(ctx context.Context, sessionID string, userID uint, ttl time.Duration) error {
	return r.client.Set(ctx, sessionID, userID, ttl).Err()
}

func (r *RedisSessionRepository) GetSession(ctx context.Context, sessionID string) (string, error) {
	return r.client.Get(ctx, sessionID).Result()
}

func (r *RedisSessionRepository) DeleteSession(ctx context.Context, sessionID string) error {
	return r.client.Del(ctx, sessionID).Err()
}
