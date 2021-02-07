package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	shared "alejandrogarcia.com/alejogs4/todolist/shared/infraestructure"
	tasksRouting "alejandrogarcia.com/alejogs4/todolist/tasks/infraestructure"
	"github.com/gorilla/mux"
)

func main() {
	defer shared.PostgresDB.Close()

	router := mux.NewRouter()
	tasksRouting.HandleTasksPetitions(router)

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Runnig app in 8080")
	log.Fatal(server.ListenAndServe())
}
