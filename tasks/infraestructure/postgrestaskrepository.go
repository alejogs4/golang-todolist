package infraestructure

import (
	"time"

	"alejandrogarcia.com/alejogs4/todolist/shared/infraestructure"
	"alejandrogarcia.com/alejogs4/todolist/tasks/domain"
	"alejandrogarcia.com/alejogs4/todolist/tasks/domain/taskstate"
)

// PostgresTaskRepository is the concrete implementation for postgres database of tasks repository
type PostgresTaskRepository struct{}

// CreateNewTask create a new task in a postgres database
func (repository *PostgresTaskRepository) CreateNewTask(task domain.Task) error {
	_, error := infraestructure.PostgresDB.Exec(
		"INSERT INTO tasks(id, title, description, due_date, state) VALUES($1, $2, $3, $4, $5)",
		task.ID, task.Title, task.Description, task.DueDate, task.State,
	)

	return error
}

// ChangeTaskState change task state in a postgres database
func (repository *PostgresTaskRepository) ChangeTaskState(task domain.Task) error {
	_, error := infraestructure.PostgresDB.Exec(
		"UPDATE tasks SET state=$1 WHERE id=$2",
		task.State, task.ID,
	)

	return error
}

// GetTasks get all tasks from postgres database avoiding DISCARTED tasks
func (repository *PostgresTaskRepository) GetTasks() ([]domain.Task, error) {
	rows, error := infraestructure.PostgresDB.Query(
		"SELECT id, title, description, due_date, state FROM tasks WHERE state != $1 ORDER BY state DESC",
		taskstate.DISCARTED,
	)

	if error != nil {
		return nil, error
	}

	var tasks []domain.Task
	for rows.Next() {
		var id string
		var title string
		var description string
		var dueDate time.Time
		var state string

		error = rows.Scan(&id, &title, &description, &dueDate, &state)
		if error != nil {
			return nil, error
		}

		tasks = append(tasks, domain.Task{
			ID:          id,
			Title:       title,
			Description: description,
			DueDate:     dueDate,
			State:       taskstate.TaskState{Value: state},
		})
	}

	return tasks, nil
}
