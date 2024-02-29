package handlers

// it is used in case of http. Otherwise we create interceptors for grpc
// this is the decorator for AdminHandlers
type AdminHandlersLogger struct { // implements handlers interface
	inner // AdminHandlersInterface here
	logger
}