package application

import "alejandrogarcia.com/alejogs4/todolist/tasks/domain"

// TaskQueries structure
type TaskQueries struct {
	TaskRepository domain.TaskRespository
}

// GetUndiscartedTasks query to get all those tasks which are either in todo or completed state
func (t TaskQueries) GetUndiscartedTasks() ([]domain.Task, error) {
	return t.TaskRepository.GetTasks()
}
