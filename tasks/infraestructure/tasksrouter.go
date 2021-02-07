package infraestructure

import "github.com/gorilla/mux"

func HandleTasksPetitions(router *mux.Router) {
	router.HandleFunc("/api/v1/tasks", getAllTasksController).Methods("GET")
	router.HandleFunc("/api/v1/task", createNewTask).Methods("POST")
	router.HandleFunc("/api/v1/task/state", changeTaskState).Methods("PUT")
}
