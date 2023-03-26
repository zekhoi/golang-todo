package errors

type Unauthorized struct {
	Message string
}

func (e *Unauthorized) Unauthorized() string {
	return e.Message
}
