package examplestore

type StorageMongo struct {
	// ...
}


func NewStorageMongo(mongoClient *mongodriver.Client) {
	// ...
}

// create DTO models for mongo with bson tags

// create converters(from, to) for entity and dto 

// create mongo methods which implement Storage interface