package service

import "github.com/Snork2017/example-srv/internal/domain"

type Service interface{}

type ServiceBasic struct {
	storage domain.Storage
}

// create business logic
