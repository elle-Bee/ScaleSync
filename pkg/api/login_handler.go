package api

import (
	"log"
	"net/http"

	"ScaleSync/pkg/database"
	"ScaleSync/pkg/service"

	"github.com/gorilla/mux"
)

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users", h.CreateUser).Methods("POST")
	r.HandleFunc("/users", h.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", h.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", h.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", h.DeleteUser).Methods("DELETE")
}

func StartServer() {
	dbPool := database.InitDB()
	defer dbPool.Close()

	r := mux.NewRouter()

	userService := service.NewUserService(dbPool) // Assuming NewUserService initializes the service with the database connection
	handler := NewUserHandler(userService)
	handler.RegisterRoutes(r)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
