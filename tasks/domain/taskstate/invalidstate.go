package taskstate

import "errors"

type InvalidState struct {
	NewState string
}

func (state InvalidState) Error() error {
	return errors.New("Invalid state")
}
