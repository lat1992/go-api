package model

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Cache struct {
	redis *redis.Client
}

func newCache(uri string, database int) (*Cache, error) {
	options, err := redis.ParseURL(uri)
	if err != nil {
		return nil, err
	}
	options.Username = ""
	options.DB = database
	rd := redis.NewClient(options)
	_, err = rd.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &Cache{
		redis: rd,
	}, nil
}
