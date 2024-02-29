package handlers

// Register user handlers

// it is used in case of http. Otherwise we create interceptors for grpc
// this is the decorator for UserHandlers
type AdminHandlersLogger struct { // implements handlers interface
	inner // UserHandlersInterface here
	logger
}