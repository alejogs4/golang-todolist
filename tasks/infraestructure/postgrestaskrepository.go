package infraestructure

import (
	"database/sql"
	"time"

	shared "alejandrogarcia.com/alejogs4/todolist/shared/infraestructure"
	"alejandrogarcia.com/alejogs4/todolist/tasks/domain"
	"alejandrogarcia.com/alejogs4/todolist/tasks/domain/taskstate"
)

// PostgresTaskRepository is the concrete implementation for postgres database of tasks repository
type PostgresTaskRepository struct{}

// CreateNewTask create a new task in a postgres database
func (repository PostgresTaskRepository) CreateNewTask(task domain.Task) error {
	_, error := shared.PostgresDB.Exec(
		"INSERT INTO tasks(id, title, description, due_date, state) VALUES($1, $2, $3, $4, $5)",
		task.ID, task.Title, task.Description, task.DueDate, task.State.Value,
	)

	return error
}

// ChangeTaskState change task state in a postgres database
func (repository PostgresTaskRepository) ChangeTaskState(taskID, newState string) error {
	_, error := shared.PostgresDB.Exec(
		"UPDATE tasks SET state=$1 WHERE id=$2",
		newState, taskID,
	)

	return error
}

// GetTask get single task by its ID, return error if task doesn't exist
func (repository PostgresTaskRepository) GetTask(taskID string) (domain.Task, error) {
	var id string
	var title string
	var description string
	var dueDate time.Time
	var state string

	row := shared.PostgresDB.QueryRow("SELECT id, title, description, due_date, state FROM tasks WHERE id=$1", taskID)
	err := row.Scan(&id, &title, &description, &dueDate, &state)

	if err != nil && err.Error() == sql.ErrNoRows.Error() {
		return domain.Task{}, domain.NotExistentTask{TaskID: taskID}
	}

	if err != nil {
		return domain.Task{}, err
	}

	return domain.Task{ID: id, Title: title, Description: description, DueDate: dueDate, State: taskstate.TaskState{Value: state}}, nil
}

// GetTasks get all tasks from postgres database avoiding DISCARTED tasks
func (repository PostgresTaskRepository) GetTasks() ([]domain.Task, error) {
	rows, error := shared.PostgresDB.Query(
		"SELECT id, title, description, due_date, state FROM tasks WHERE state != $1 ORDER BY state DESC",
		taskstate.DISCARTED,
	)

	if error != nil {
		return nil, error
	}

	var tasks []domain.Task = []domain.Task{}
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
