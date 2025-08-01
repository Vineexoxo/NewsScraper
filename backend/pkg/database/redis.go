package database

import (
	"context"
	"fmt"
	"time"
	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	*redis.Client
}

// NewRedis creates a new Redis connection
func NewRedis(addr, password string, db int) (*RedisDB, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return &RedisDB{Client: client}, nil
}

// Health checks Redis health
func (r *RedisDB) Health() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	return r.Ping(ctx).Err()
}