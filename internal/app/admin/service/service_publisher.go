package service

// Responsible for being decorator for service. Publish events to broker.

type Publisher interface {
	Publish(event string, data interface{}) error
}

type ServicePublisher struct {
	publisher Publisher
}
