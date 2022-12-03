package exception

type WrongOtpError struct {
	Error string
}

func NewWrongOtpError(err error) WrongOtpError {
	return WrongOtpError{Error: string(err.Error())}
}
