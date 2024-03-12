package examplestore

import (
	"github.com/Snork2017/example-srv/internal/domain"
	"github.com/redis/go-redis/v9"
)

type Storage interface{}

type StorageRedis struct {
	inner       domain.Storage
	redisClient *redis.Client
}

func NewStorageRedis(redisClient *redis.Client, inner domain.Storage) *StorageRedis {
	return &StorageRedis{
		inner:       inner,
		redisClient: redisClient,
	}
}

// create redis methods which implement Storage interface

// redis caches entity by key
