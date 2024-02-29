package examplestore

// storage interface uses entities 
// methods receive entites and convert them to models(redis or mongo -> depends on usage)
type Storage interface {

}

// if we use facade then we create facade here.

type StorageCache struct {
	cache Storage
	db Storage
}
// or
type StorageFacade struct {
	cache Storage
	db Storage
}