package exception

import "errors"

type NotFoundError struct {
	Error string
}

func newNotFoundError(err error) NotFoundError {
	return NotFoundError{Error: string(err.Error())}
}

func PanicIfNotFound(userId int) {
	if userId == 0 {
		panic(newNotFoundError(errors.New("user not found")))
	}
}
