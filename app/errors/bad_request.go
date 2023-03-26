package errors

type BadRequestError struct {
	Message string
}

func (e *BadRequestError) BadRequestError() string {
	return e.Message
}
