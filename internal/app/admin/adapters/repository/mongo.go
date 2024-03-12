package examplestore

import "go.mongodb.org/mongo-driver/mongo"

type StorageMongo struct {
	// ...
}

func NewStorageMongo(mongoClient *mongo.Client) {
	// ...
}

// create DTO models for mongo with bson tags

// create converters(from, to) for entity and dto

// create mongo methods which implement Storage interface
