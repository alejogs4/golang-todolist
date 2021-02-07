package taskstate

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
	if !IsValidState(newState) {
		return TaskState{}, InvalidState{NewState: newState}
	}

	if state.Value == DISCARTED {
		return TaskState{}, InvalidState{NewState: newState}
	}

	return TaskState{Value: newState}, nil
}

// IsValidState verify is passed state is a valid one
func IsValidState(state string) bool {
	return state == TODO || state == COMPLETED || state == DISCARTED
}
