package interfaces

type ServiceError struct {
	Error      error
	StatusCode int
}
