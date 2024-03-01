package examplestore

type Storage interface {}

type StorageRedis struct {
	inner domain.Storage
	// ...
}

func NewStorageRedis(redisClient) {

}

// create redis methods which implement Storage interface

// redis caches entity by key
