package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	pool = InitDB() // Initialize the database connection
	defer pool.Close()

	r := mux.NewRouter()

	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users", GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
