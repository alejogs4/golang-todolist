package taskstate

import "strings"

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

// CreateTasktState Factory function to create a valid task state
func CreateTasktState(taskState string) (TaskState, error) {
	normalizedTaskState := NormalizeTaskState(taskState)

	if !IsValidState(normalizedTaskState) {
		return TaskState{}, InvalidState{NewState: normalizedTaskState}
	}

	return TaskState{Value: normalizedTaskState}, nil
}

// NewTaskStateTransition is the finite state machine for task state
func (state TaskState) NewTaskStateTransition(newState string) (TaskState, error) {
	normalizedTaskState := NormalizeTaskState(newState)

	if !IsValidState(normalizedTaskState) {
		return TaskState{}, InvalidState{NewState: normalizedTaskState}
	}

	if state.Value == DISCARTED {
		return TaskState{}, InvalidState{NewState: normalizedTaskState}
	}

	return TaskState{Value: normalizedTaskState}, nil
}

// IsValidState verify is passed state is a valid one
func IsValidState(state string) bool {
	return state == TODO || state == COMPLETED || state == DISCARTED
}

// NormalizeTaskState utility function to put task state in a proper shape to be used in comparisons and bussines logic
func NormalizeTaskState(taskState string) string {
	return strings.ToUpper(strings.Join(strings.Fields(strings.TrimSpace(taskState)), ""))
}
