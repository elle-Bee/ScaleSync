package main

import (
	"log"
	"net/http"

	"ScaleSync/pkg/api"
	"ScaleSync/pkg/database"

	"github.com/gorilla/mux"
)

func Login() {
	database.Pool = database.InitDB() // Initialize the database connection
	defer database.Pool.Close()

	r := mux.NewRouter()

	r.HandleFunc("/users", api.CreateUser).Methods("POST")
	r.HandleFunc("/users", api.GetAllUsers).Methods("GET")

	r.HandleFunc("/users/{id}", api.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", api.UpdateUser).Methods("PATCH")
	r.HandleFunc("/users/{id}", api.DeleteUser).Methods("DELETE")

	r.HandleFunc("/login", api.LoginUser).Methods("POST")

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	Login()
}
