package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
/* 	DB_HOST = "postgres"
	DB_PORT = 5432 */
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "default_database"
)

func Connect() *sql.DB {
	/* dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME) */
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	CheckError(err)

	return db
	/* if err != nil {
			log.Println("connect open")
			log.Fatal(err)
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			log.Println("ping")
			panic(err)
		} else {
			fmt.Println("Connected")
		}

		rows, err := db.Query(
			`SELECT *
	       	 FROM new`)
		if err != nil {
			log.Fatal(err)
		}

		// Columns
		cs, err := rows.Columns()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", cs)

		// Query Row
		var amname string
		err1 := db.QueryRow(`SELECT id FROM new`).Scan(&amname)
		switch {
		case err1 == sql.ErrNoRows:
			log.Printf("No user with that ID.")
		case err1 != nil:
			log.Fatal(err)
		default:
			fmt.Printf("amname is %s\n", amname)
		}

		// Query
		rows1, err := db.Query(`SELECT * FROM new`)
		if err != nil {
			log.Fatal(err)
		}

		for rows1.Next() {
			var amname string
			if e := rows1.Scan(&amname); e != nil {
				fmt.Println(e)
			}
			fmt.Printf("amname is %s\n", amname)
		} */
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
