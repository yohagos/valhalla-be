package main

import (
	"fmt"
	"valhalla-be/database"
)

func main() {
	db := database.Connect()

	rows, err := db.Query("SELECT * FROM new")
	CheckError(err)

    fmt.Println(rows)

	/* router := router.NewRouter()

	log.Println("Router starting...")
	if err := http.ListenAndServe(":8888", router); err != nil {
		log.Panic(err)
	} */
}

func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}