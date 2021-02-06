package domain

import (
	"time"

	"alejandrogarcia.com/alejogs4/todolist/tasks/domain/taskstate"
)

// Task represent the structure of a class
type Task struct {
	ID          string
	Title       string
	Description string
	DueDate     time.Time
	State       taskstate.TaskState
}
