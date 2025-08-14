package database

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type RedisDB struct {
	*redis.Client
	config *RedisConfig
}

// NewRedisConnStr builds the Redis address string
func NewRedisConnStr(cfg *RedisConfig) string {
	return fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
}

// NewRedisDB connects to Redis and returns a RedisDB
func NewRedisDB(addr string, config *RedisConfig) (*RedisDB, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Password,
		DB:       config.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return &RedisDB{Client: client, config: config}, nil
}

func (r *RedisDB) Close() error {
	if r.Client != nil {
		return r.Client.Close()
	}
	return nil
}

func (r *RedisDB) Health() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return r.Ping(ctx).Err()
}

// ---- Job-related helpers ----

func jobStatusKey(jobID string) string { return fmt.Sprintf("job:%s:status", jobID) }
func jobDataKey(jobID string) string   { return fmt.Sprintf("job:%s:data", jobID) }

func (r *RedisDB) SetJobStatus(ctx context.Context, jobID, status string, ttl time.Duration) error {
	return r.Set(ctx, jobStatusKey(jobID), status, ttl).Err()
}

func (r *RedisDB) GetJobStatus(ctx context.Context, jobID string) (string, error) {
	return r.Get(ctx, jobStatusKey(jobID)).Result()
}

func (r *RedisDB) SetJobResult(ctx context.Context, jobID string, data interface{}, ttl time.Duration) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return r.Set(ctx, jobDataKey(jobID), b, ttl).Err()
}

func (r *RedisDB) GetJobResult(ctx context.Context, jobID string, out interface{}) error {
	s, err := r.Get(ctx, jobDataKey(jobID)).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(s), out)
}
