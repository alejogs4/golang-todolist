package application

import "alejandrogarcia.com/alejogs4/todolist/tasks/domain"

type TaskQueries struct {
	TaskRepository domain.TaskRespository
}

func (t TaskQueries) GetUndiscartedTasks() ([]domain.Task, error) {
	return t.TaskRepository.GetTasks()
}
