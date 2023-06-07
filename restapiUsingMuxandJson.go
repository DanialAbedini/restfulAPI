package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initalazeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	err := http.ListenAndServe(":9000", r)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	intialMigration()
	initalazeRouter()

}
