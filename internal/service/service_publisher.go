package service

// Responsible for being decorator for service. Publish events to broker.

type ServicePublisher struct {
	publisher domain.Publisher
}