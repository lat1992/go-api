package store

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"golang.org/x/exp/slog"
)

type Store struct {
	redis *redis.Client
}

func NewStore(uri string, database int) (*Store, error) {
	options, err := redis.ParseURL(uri)
	if err != nil {
		return nil, fmt.Errorf("newCache: %w", err)
	}
	options.Username = ""
	options.DB = database
	rd := redis.NewClient(options)
	_, err = rd.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("newCache: %w", err)
	}
	return &Store{
		redis: rd,
	}, nil
}

func (s *Store) Close() {
	if err := s.redis.Close(); err != nil {
		slog.Error("Cache closing error", err)
	}
}
