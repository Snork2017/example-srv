package domain

import (
	"context"

	"github.com/google/uuid"
)

type Storage interface {
	// GetByID returns the entity with the given ID
	GetByID(ctx context.Context, id uuid.UUID) (*ExampleData, error)
	// GetByField returns the entity with the given field
	GetByField(ctx context.Context, field string) (*ExampleData, error)
}
