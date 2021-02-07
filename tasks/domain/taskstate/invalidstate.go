package taskstate

import (
	"fmt"
)

// InvalidState model the domain error that trigger the try of change one state to an invalid one
type InvalidState struct {
	NewState string
}

func (state InvalidState) Error() string {
	return state.Message()
}

// Message of the invalid state
func (state InvalidState) Message() string {
	return fmt.Sprintf("%s is an invalid state", state.NewState)
}
