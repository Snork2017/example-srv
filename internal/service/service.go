package app

type Service interface {
	storage domain.Storage
}

// create business logic