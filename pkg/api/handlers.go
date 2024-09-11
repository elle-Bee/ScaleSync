package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	db "ScaleSync/pkg/database"
	"ScaleSync/pkg/models"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	err := db.Pool.QueryRow(context.Background(), query, user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		log.Println("Create User Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var user models.User
	query := `SELECT id, name, email FROM users WHERE id = $1`
	err := db.Pool.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query := `UPDATE users SET name=$1, email=$2 WHERE id=$3`
	_, err := db.Pool.Exec(context.Background(), query, user.Name, user.Email, id)
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		log.Println("Update User Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	query := `DELETE FROM users WHERE id=$1`

	_, err := db.Pool.Exec(context.Background(), query, id)
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Pool.Query(context.Background(), `SELECT id, name, email FROM users`)
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		log.Println("Get All Users Error:", err)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Println("Error scanning user:", err)
			continue
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
