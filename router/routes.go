package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	log.Println("load routes..")
	router.HandleFunc("/", getHandleFunc).Methods("GET")
	router.HandleFunc("/post", postHandleFunc).Methods("GET")
	return router
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func getHandleFunc(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("get")
}

func postHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("posts")
}

