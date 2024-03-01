package service

type Storage interface {}

type Service interface {
}

type ServiceBasic struct {
	storage Storage
}

// create business logic