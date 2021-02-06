package taskstate

import "errors"

const (
	// TODO Task that is pending to be done
	TODO = "TODO"
	// COMPLETED Task that has been already done
	COMPLETED = "COMPLETED"
	// DISCARTED Task that has been discarted by any reason, it's supposed to be never done
	DISCARTED = "DISCARTED"
)

// TaskState value object which will represent task state and its state transitions
type TaskState struct {
	Value string
}

// NewTaskStateTransition is the finite state machine for task state
func (state TaskState) NewTaskStateTransition(newState string) (TaskState, error) {
	if newState != TODO && newState != COMPLETED && newState != DISCARTED {
		return TaskState{}, errors.New("Invalid state")
	}

	if state.Value == DISCARTED {
		return TaskState{}, errors.New("Invalid state")
	}

	return TaskState{Value: newState}, nil
}
