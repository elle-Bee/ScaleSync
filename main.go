package main

import (
	"ScaleSync/app"
	"ScaleSync/pkg/api"
	"ScaleSync/pkg/database"
	"ScaleSync/pkg/metrics"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	database := database.InitDB()
	defer database.Close() // Ensure the connection pool closes when the program exits

	go startServer()

	os.Setenv("FYNE_THEME", "dark")
	app.App()
}

func startServer() {
	r := mux.NewRouter()

	r.Handle("/metrics", promhttp.Handler())
	metrics.RegisterMetrics()
	http.Handle("/metrics", promhttp.Handler())

	// User routes
	r.HandleFunc("/users", api.CreateUser).Methods("POST")
	r.HandleFunc("/users", api.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", api.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", api.UpdateUser).Methods("PATCH")
	r.HandleFunc("/users/{id}", api.DeleteUser).Methods("DELETE")

	// Login route
	r.HandleFunc("/login", api.LoginUser).Methods("POST")
	r.HandleFunc("/add", api.InsertUser).Methods("GET")

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
