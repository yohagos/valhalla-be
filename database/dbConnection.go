package database

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host = ""
	port = ""
	user = ""
	pwd = ""
	dbname = ""
)

func Connect() {
	psqlconn := fmt.Sprintf("host = %s port = %d user = %s dbname = %s sslmode=disable", host, port, user, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	insertStmt := `insert into "Employee("Name", "EmpID") values ('Yosef', 12)`

	result, err := db.Exec(insertStmt)
	if err != nil {
		log.Println(err)
	}

	println(result)

	insertDynStmt := `insert into "Employee("Name", "EmpID") values ($1, $2)`

	result2, err := db.Exec(insertDynStmt, "Binyam", 160)
	if err != nil {
		log.Println(err)
	}

	println(result2)
}