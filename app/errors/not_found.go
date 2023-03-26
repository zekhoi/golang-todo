package errors

import "fmt"

type NotFoundError struct {
	EntityName string
	EntityID   string
}

func (e *NotFoundError) NotFoundError() string {
	return fmt.Sprintf("%s with ID %s not found", e.EntityName, e.EntityID)
}
