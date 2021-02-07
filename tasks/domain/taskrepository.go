package domain

// TaskRespository represents how interact with any datasource
type TaskRespository interface {
	// Queries
	GetTasks() ([]Task, error)
	GetTask(taskID string) (Task, error)
	// Commands
	CreateNewTask(task Task) error
	ChangeTaskState(taskID, newState string) error
}
