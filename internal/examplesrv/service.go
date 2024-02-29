package examplesrv

type Service interface {
	storage examplestore.Storage
}

// create business logic