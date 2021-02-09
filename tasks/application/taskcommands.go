package application

import (
	"time"

	"alejandrogarcia.com/alejogs4/todolist/tasks/domain"
	"alejandrogarcia.com/alejogs4/todolist/tasks/domain/taskstate"
	"github.com/google/uuid"
)

// TaskCommands contains the set of operations that are thought to change information about tasks
type TaskCommands struct {
	TaskRepository domain.TaskRespository
}

// CreateNewTask is the application layer operation to add a new task, here bussines rules are checked
func (t TaskCommands) CreateNewTask(title, description string, dueDate time.Time, state string) error {
	newTaskState, error := taskstate.CreateTasktState(state)
	if error != nil {
		return error
	}

	taskID := uuid.New()
	newTask := domain.Task{
		ID:          taskID.String(),
		Description: description,
		Title:       title,
		DueDate:     dueDate,
		State:       newTaskState,
	}

	return t.TaskRepository.CreateNewTask(newTask)
}

// ChangeTaskState will change a task state if accomplish with bussines rules
func (t TaskCommands) ChangeTaskState(taskID, newState string) error {
	task, error := t.TaskRepository.GetTask(taskID)
	if error != nil {
		return error
	}

	newTaskState, error := task.State.NewTaskStateTransition(newState)
	if error != nil {
		return error
	}

	return t.TaskRepository.ChangeTaskState(taskID, newTaskState.Value)
}
