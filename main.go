package main

import (
	"log"
	"net/http"
	"valhalla-be/database"
	"valhalla-be/router"
)

func main() {
	database.Connect()

	router := router.NewRouter()

	log.Println("Router starting...")
	if err := http.ListenAndServe(":8888", router); err != nil {
		log.Panic(err)
	}
}
