package infraestructure

import (
	"database/sql"
	"fmt"
	"log"

	// Meanwhile I learn better patterns this is how I'll connect to database
	_ "github.com/lib/pq"
)

// TODO: Improve this information using enviroment variables
const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "todolist"
)

// PostgresDB database connection to interact with the database
var PostgresDB *sql.DB

func init() {
	postgresConnection := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable", host, port, user, dbname)

	var error error
	PostgresDB, error = sql.Open("postgres", postgresConnection)

	if error != nil {
		log.Fatal(error)
	}

	err := PostgresDB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
