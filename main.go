package main

import (
	"ScaleSync/app"
	"ScaleSync/pkg/api"
	"ScaleSync/pkg/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database connection
	database.Pool = database.InitDB()
	defer database.Pool.Close() // Ensure the connection pool closes when the program exits

	// Start the HTTP server
	go startServer()

	// Start the app (e.g., GUI or other application logic)
	app.App()
}

func startServer() {
	r := mux.NewRouter()

	// User routes
	r.HandleFunc("/users", api.CreateUser).Methods("POST")
	r.HandleFunc("/users", api.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", api.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", api.UpdateUser).Methods("PATCH")
	r.HandleFunc("/users/{id}", api.DeleteUser).Methods("DELETE")

	// Login route
	r.HandleFunc("/login", api.LoginUser).Methods("POST")

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
