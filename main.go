package main

import (
	"fmt"

	shared "alejandrogarcia.com/alejogs4/todolist/shared/infraestructure"
)

func main() {
	defer shared.PostgresDB.Close()
	fmt.Println("Runnig app")
}
