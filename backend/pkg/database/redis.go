package database

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"

	pb "github.com/shishir54234/NewsScraper/backend/service/web-scraper/web-scraper/grpc_server/proto"
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

// --- Redis key helpers ---
func jobDataKey(jobID string) string { return fmt.Sprintf("job:%s:data", jobID) }
func jobURLKey(url string) string    { return fmt.Sprintf("url:%s", url) }

// --- Store JobResult ---
func (r *RedisDB) SetJobResult(ctx context.Context, jobID, url string, result *pb.GetResultResponse, ttl time.Duration) error {
	data, err := proto.Marshal(result)
	if err != nil {
		return err
	}

	pipe := r.TxPipeline()
	pipe.Set(ctx, jobDataKey(jobID), data, ttl)
	// Store URL → jobID mapping
	pipe.Set(ctx, jobURLKey(url), jobID, ttl)

	_, err = pipe.Exec(ctx)
	return err
}

// Retrieve JobResult
func (r *RedisDB) GetJobResult(ctx context.Context, jobID string) (*pb.GetResultResponse, error) {
	bytes, err := r.Get(ctx, jobDataKey(jobID)).Bytes()
	if err != nil {
		return nil, err
	}
	var result pb.GetResultResponse
	if err := proto.Unmarshal(bytes, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *RedisDB) GetJobIDByURL(ctx context.Context, url string) (string, error) {
	return r.Get(ctx, jobURLKey(url)).Result()
}

// Query by URL: returns PageData & Status
// This allows quick access to the scraping result by URL
func (r *RedisDB) GetResultByURL(ctx context.Context, url string) (*pb.PageData, pb.Status, error) {
	jobID, err := r.GetJobIDByURL(ctx, url)
	if err == redis.Nil {
		// No mapping exists
		return nil, pb.Status_UNDEFINED, nil
	}
	if err != nil {
		return nil, pb.Status_UNDEFINED, err
	}

	result, err := r.GetJobResult(ctx, jobID)
	if err != nil {
		return nil, pb.Status_UNDEFINED, err
	}

	return result.Page, result.Status, nil
}
