package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func New(dbPath string, timeout time.Duration) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     dbPath,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := client.Ping(ctx); err != nil {
		return nil, err.Err()
	}

	return client, nil
}
