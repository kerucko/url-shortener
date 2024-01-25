package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	db *redis.Client
}

func New(dbPath string, timeout time.Duration) (*RedisStorage, error) {
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

	return &RedisStorage{client}, nil
}

func (s *RedisStorage) SaveUrl(saveUrl string, alias string, expiration time.Duration, timeout time.Duration) error {
	op := "redis.SaveUrl"

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err := s.db.Get(ctx, alias).Result()
	if errors.Is(err, redis.Nil) {
		return fmt.Errorf("%s: alias already exist", op)
	}
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err = s.db.Set(ctx, alias, saveUrl, expiration).Err()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}


func (s *RedisStorage) GetUrl(alias string, timeout time.Duration) (string, error) {
	op := "redis.GetUrl"

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	result, err := s.db.Get(ctx, alias).Result()
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return result, nil
} 

func (s *RedisStorage) DeleteUrl(alias string, timeout time.Duration) error {
	op := "redis.DeleteUrl"

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err := s.db.Del(ctx, alias).Err()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}