package exception

import "errors"

type EmailUsedError struct {
	Error string
}

func newEmailUsedError(err error) EmailUsedError {
	return EmailUsedError{
		Error: string(err.Error()),
	}
}

func PanicIfEmailUsed(requestedEmail string, usedEmail string) {
	if usedEmail == requestedEmail {
		panic(newEmailUsedError(errors.New("email is already used")))
	}
}
