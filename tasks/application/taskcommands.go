package application

import (
	"strings"
	"time"

	"alejandrogarcia.com/alejogs4/todolist/tasks/domain"
	"alejandrogarcia.com/alejogs4/todolist/tasks/domain/taskstate"
	"github.com/google/uuid"
)

// TaskCommands contains the set of operations that are thought to change information about tasks
type TaskCommands struct {
	taskRepository domain.TaskRespository
}

// CreateNewTask is the application layer operation to add a new task, here bussines rules are checked
func (t TaskCommands) CreateNewTask(title, description string, dueDate time.Time, state string) error {
	normalizedState := strings.ToUpper(strings.Join(strings.Fields(strings.TrimSpace(state)), ""))

	if !taskstate.IsValidState(normalizedState) {
		return taskstate.InvalidState{NewState: normalizedState}
	}

	taskID := uuid.New()
	newTask := domain.Task{
		ID:      taskID.String(),
		Title:   title,
		DueDate: dueDate,
		State:   taskstate.TaskState{Value: normalizedState},
	}

	return t.taskRepository.CreateNewTask(newTask)
}