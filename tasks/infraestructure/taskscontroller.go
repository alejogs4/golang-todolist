package infraestructure

import (
	"encoding/json"
	"net/http"
	"time"

	shared "alejandrogarcia.com/alejogs4/todolist/shared/domain"
	"alejandrogarcia.com/alejogs4/todolist/shared/infraestructure"
	"alejandrogarcia.com/alejogs4/todolist/tasks/application"
)

var taskQueries = application.TaskQueries{TaskRepository: PostgresTaskRepository{}}
var taskCommands = application.TaskCommands{TaskRepository: PostgresTaskRepository{}}

func getAllTasksController(response http.ResponseWriter, request *http.Request) {
	tasks, error := taskQueries.GetUndiscartedTasks()

	response.Header().Set("Content-Type", "application/json")

	if error != nil {
		infraestructure.DispatchNewHttpError(response, "Something wen wrong", http.StatusInternalServerError)
		return
	}

	responseContent, _ := json.Marshal(infraestructure.WrapAPIResponse(tasks))

	response.WriteHeader(http.StatusOK)
	response.Write(responseContent)
}

func createNewTask(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var task struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		DueDate     string `json:"dueDate"`
		State       string `json:"state"`
	}
	error := json.NewDecoder(request.Body).Decode(&task)

	if error != nil {
		infraestructure.DispatchNewHttpError(response, "Something went wrong", http.StatusInternalServerError)
		return
	}

	dueDate, dateError := time.Parse("2006-01-02T15:04:05.000Z", task.DueDate)
	if dateError != nil {
		infraestructure.DispatchNewHttpError(response, "Incorrect date format", http.StatusBadRequest)
		return
	}

	error = taskCommands.CreateNewTask(task.Title, task.Description, dueDate, task.State)
	if error == nil {
		response.WriteHeader(http.StatusCreated)
		return
	}

	domainError, ok := error.(shared.DomainError)
	if ok {
		httpError := infraestructure.NewHTTPError(domainError)
		infraestructure.DispatchNewHttpError(response, httpError.Message, httpError.Status)
		return
	}

	infraestructure.DispatchNewHttpError(response, "Something went wrong", http.StatusInternalServerError)
}

func changeTaskState(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var taskInformation struct {
		TaskID   string `json:"taskID"`
		NewState string `json:"newState"`
	}

	error := json.NewDecoder(request.Body).Decode(&taskInformation)
	if error != nil {
		infraestructure.DispatchNewHttpError(response, "Something went wrong", http.StatusInternalServerError)
		return
	}

	error = taskCommands.ChangeTaskState(taskInformation.TaskID, taskInformation.NewState)
	if error == nil {
		response.WriteHeader(http.StatusAccepted)
		return
	}

	domainError, ok := error.(shared.DomainError)
	if ok {
		httpError := infraestructure.NewHTTPError(domainError)
		infraestructure.DispatchNewHttpError(response, httpError.Message, httpError.Status)
		return
	}

	infraestructure.DispatchNewHttpError(response, "Something went wrong", http.StatusInternalServerError)
}
