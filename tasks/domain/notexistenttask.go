package domain

import "fmt"

// NotExistentTask modeled an error when a task doesn't exist
type NotExistentTask struct {
	TaskID string
}

func (nt NotExistentTask) Error() string {
	return nt.Message()
}

// Message of the task that doesn't exist
func (nt NotExistentTask) Message() string {
	return fmt.Sprintf("Task %s doesn't exist", nt.TaskID)
}
