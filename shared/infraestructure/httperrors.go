package shared

import (
	"net/http"

	shared "alejandrogarcia.com/alejogs4/todolist/shared/domain"
	"alejandrogarcia.com/alejogs4/todolist/tasks/domain/taskstate"
)

// HTTPError any valid http error
type HTTPError struct {
	Message string
	Status  int
}

// NewHTTPError is the mapper function to cast a domain error to a HTTPError
func NewHTTPError(domainError shared.DomainError) HTTPError {
	switch resultError := domainError.(type) {
	case taskstate.InvalidState:
		return HTTPError{Message: resultError.Message(), Status: http.StatusBadRequest}
	default:
		return HTTPError{}
	}
}
