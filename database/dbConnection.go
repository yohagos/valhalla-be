package database

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "postgres"
)

func Connect() {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected")
	}

	rows, err := db.Query(
		`SELECT amname, amstrategies, amsupport, amcanorder, amcanorderbyop, 
       amcanbackward, amcanunique, amcanmulticol, amoptionalkey, amsearcharray, 
       amsearchnulls, amstorage, amclusterable, ampredlocks, amkeytype, 
       aminsert, ambeginscan, amgettuple, amgetbitmap, amrescan, amendscan, 
       ammarkpos, amrestrpos, ambuild, ambuildempty, ambulkdelete, amvacuumcleanup, 
       amcanreturn, amcostestimate, amoptions
       FROM pg_am`)
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
	err1 := db.QueryRow(`SELECT amname FROM pg_am`).Scan(&amname)
	switch {
	case err1 == sql.ErrNoRows:
		log.Printf("No user with that ID.")
	case err1 != nil:
		log.Fatal(err)
	default:
		fmt.Printf("amname is %s\n", amname)
	}

	// Query
	rows1, err := db.Query(`SELECT amname FROM pg_am`)
	if err != nil {
		log.Fatal(err)
	}

	for rows1.Next() {
		var amname string
		if e := rows1.Scan(&amname); e != nil {
			fmt.Println(e)
		}
		fmt.Printf("amname is %s\n", amname)
	}
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
