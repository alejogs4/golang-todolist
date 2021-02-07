package shared

import (
	"database/sql"
	"fmt"
	"log"

	// Meanwhile I learn better patterns this is how I'll connect to database
	_ "github.com/lib/pq"
)

// TODO: Improve this information using enviroment variables
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "todolist"
)

// PostgresDB database connection to interact with the database
var PostgresDB *sql.DB

func init() {
	postgresConnection := fmt.Sprint("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var error error
	PostgresDB, error = sql.Open("postgres", postgresConnection)

	if error != nil {
		log.Fatal(error)
	}
}
