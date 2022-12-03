package exception

type UnauthorizedError struct {
	Error string
}

func NewUnauthorizedError(err error) UnauthorizedError {
	return UnauthorizedError{Error: string(err.Error())}
}
