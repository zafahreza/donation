package exception

import "errors"

type NotFoundError struct {
	Error string
}

func NewNotFoundError(err error) NotFoundError {
	return NotFoundError{Error: string(err.Error())}
}

func PanicIfNotFound(userId int) {
	if userId == 0 {
		panic(NewNotFoundError(errors.New("user not found")))
	}
}
