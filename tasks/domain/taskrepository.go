package domain

// TaskRespository represents how interact with any datasource
type TaskRespository interface {
	// Queries
	GetTasks() ([]Task, error)
	// Commands
	CreateNewTask(task Task) error
	ChangeTaskState(task Task) error
}
