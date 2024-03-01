package service

type Service interface {
}

type ServiceBasic struct {
	storage domain.Storage
}

// create business logic