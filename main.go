package main

import (
	"fmt"

	"alejandrogarcia.com/alejogs4/todolist/shared/infraestructure"
)

func main() {
	defer infraestructure.PostgresDB.Close()
	fmt.Println("Runnig app")
}
