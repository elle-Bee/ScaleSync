package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"ScaleSync/pkg/database"
	"ScaleSync/pkg/models"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	var exists bool
	checkQuery := `Select exists(SELECT 1 FROM users WHERE name = '` + user.Name + `')`
	database.Pool.QueryRow(context.Background(), checkQuery).Scan(&exists)
	fmt.Println(user.Name)
	fmt.Println(exists)
	if exists {
		http.Error(w, "Username is taken", http.StatusInternalServerError)
		return
	}
	hash_passwd := database.HashPassword(user.Password)

	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	err2 := database.Pool.QueryRow(context.Background(), query, user.Name, user.Email, hash_passwd).Scan(&user.ID)
	if err2 != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		log.Println("Create User Error:", err2)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]

	// Convert the ID from string to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	query := `SELECT id, name, email FROM users WHERE id = $1`

	// Execute the query and scan the result into the user struct
	err = database.Pool.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		// Log the error for debugging
		log.Printf("Error fetching user with ID %d: %v", id, err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Set content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Return the user as a JSON response
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("Error encoding user to JSON: %v", err)
		http.Error(w, "Failed to encode user data", http.StatusInternalServerError)
	}
}

func GetLoggedInUserHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Name parameter is required", http.StatusBadRequest)
		return
	}

	user, err := GetLoggedInUser(name)
	if err != nil {
		// Log the error for debugging
		log.Printf("Error fetching user with name %s: %v", name, err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Set content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Return the user as a JSON response
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("Error encoding user to JSON: %v", err)
		http.Error(w, "Failed to encode user data", http.StatusInternalServerError)
	}
}

// GetLoggedInUser retrieves the user based on the provided name
func GetLoggedInUser(name string) (models.User_login, error) {
	var user models.User_login
	query := `SELECT id, name, email FROM users WHERE name = $1`

	// Use QueryRow to get a single row
	err := database.Pool.QueryRow(context.Background(), query, name).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return user, errors.New("user not found")
		}
		return user, err
	}

	user.Session = true // Assuming session is active
	return user, nil
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	query := `UPDATE users SET name=$1, email=$2 WHERE id=$3`
	_, err := database.Pool.Exec(context.Background(), query, user.Name, user.Email, id)
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		log.Println("Update User Error:", err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	query := `DELETE FROM users WHERE id=$1`

	_, err := database.Pool.Exec(context.Background(), query, id)
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := database.Pool.Query(context.Background(), `SELECT id, name, email FROM users`)
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

	json.NewEncoder(w).Encode(users)
}
