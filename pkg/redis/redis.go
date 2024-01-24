package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Storage struct {
	db *redis.Client
}

func New(dbPath string, timeout time.Duration) (*Storage, error) {
	op := "redis.New"
	client := redis.NewClient(&redis.Options{
		Addr:     dbPath,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{client}, nil
}
