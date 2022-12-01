package exception

type WrongPasswordError struct {
	Error string
}

func NewWrongPasswordError(err error) WrongPasswordError {
	return WrongPasswordError{Error: string(err.Error())}
}
