package api

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"ScaleSync/pkg/database"
	"ScaleSync/pkg/metrics"
	"ScaleSync/pkg/models"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

// CreateUser creates a new user and tracks Prometheus metrics
func CreateUser(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	metrics.ApiRequests.WithLabelValues("CreateUser").Inc()

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		metrics.ApiFailures.WithLabelValues("CreateUser").Inc()
		return
	}

	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM users WHERE name = $1)`
	err := database.Pool.QueryRow(context.Background(), checkQuery, user.Name).Scan(&exists)
	if err != nil || exists {
		http.Error(w, "Username is taken", http.StatusConflict)
		metrics.ApiFailures.WithLabelValues("CreateUser").Inc()
		return
	}

	hashPasswd := database.HashPassword(user.Password)
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	err = database.Pool.QueryRow(context.Background(), query, user.Name, user.Email, hashPasswd).Scan(&user.ID)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		log.Println("Create User Error:", err)
		metrics.ApiFailures.WithLabelValues("CreateUser").Inc()
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

	metrics.ApiSuccesses.WithLabelValues("CreateUser").Inc()
	metrics.ApiRequestDuration.WithLabelValues("CreateUser").Observe(time.Since(start).Seconds())
}

// GetUser retrieves a user by ID and tracks Prometheus metrics
func GetUser(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	metrics.ApiRequests.WithLabelValues("GetUser").Inc()

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		metrics.ApiFailures.WithLabelValues("GetUser").Inc()
		return
	}

	var user models.User
	query := `SELECT id, name, email FROM users WHERE id = $1`
	err = database.Pool.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		log.Printf("Error fetching user with ID %d: %v", id, err)
		http.Error(w, "User not found", http.StatusNotFound)
		metrics.ApiFailures.WithLabelValues("GetUser").Inc()
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

	metrics.ApiSuccesses.WithLabelValues("GetUser").Inc()
	metrics.ApiRequestDuration.WithLabelValues("GetUser").Observe(time.Since(start).Seconds())
}

// GetLoggedInUserHandler retrieves a logged-in user by name
func GetLoggedInUserHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	metrics.ApiRequests.WithLabelValues("GetLoggedInUser").Inc()

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Name parameter is required", http.StatusBadRequest)
		metrics.ApiFailures.WithLabelValues("GetLoggedInUser").Inc()
		return
	}

	user, err := GetLoggedInUser(name)
	if err != nil {
		log.Printf("Error fetching user with name %s: %v", name, err)
		http.Error(w, err.Error(), http.StatusNotFound)
		metrics.ApiFailures.WithLabelValues("GetLoggedInUser").Inc()
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

	metrics.ApiSuccesses.WithLabelValues("GetLoggedInUser").Inc()
	metrics.ApiRequestDuration.WithLabelValues("GetLoggedInUser").Observe(time.Since(start).Seconds())
}

// GetLoggedInUser retrieves a user by name
func GetLoggedInUser(name string) (models.User_login, error) {
	var user models.User_login
	query := `SELECT id, name, email FROM users WHERE name = $1`

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

// UpdateUser updates user information by ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	metrics.ApiRequests.WithLabelValues("UpdateUser").Inc()

	id := mux.Vars(r)["id"]
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		metrics.ApiFailures.WithLabelValues("UpdateUser").Inc()
		return
	}

	query := `UPDATE users SET name=$1, email=$2 WHERE id=$3`
	_, err := database.Pool.Exec(context.Background(), query, user.Name, user.Email, id)
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		log.Println("Update User Error:", err)
		metrics.ApiFailures.WithLabelValues("UpdateUser").Inc()
		return
	}

	json.NewEncoder(w).Encode(user)
	metrics.ApiSuccesses.WithLabelValues("UpdateUser").Inc()
	metrics.ApiRequestDuration.WithLabelValues("UpdateUser").Observe(time.Since(start).Seconds())
}

// DeleteUser deletes a user by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	metrics.ApiRequests.WithLabelValues("DeleteUser").Inc()

	id := mux.Vars(r)["id"]
	query := `DELETE FROM users WHERE id=$1`
	_, err := database.Pool.Exec(context.Background(), query, id)
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		metrics.ApiFailures.WithLabelValues("DeleteUser").Inc()
		return
	}

	w.WriteHeader(http.StatusNoContent)
	metrics.ApiSuccesses.WithLabelValues("DeleteUser").Inc()
	metrics.ApiRequestDuration.WithLabelValues("DeleteUser").Observe(time.Since(start).Seconds())
}

// GetAllUsers retrieves all users and tracks Prometheus metrics
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	metrics.ApiRequests.WithLabelValues("GetAllUsers").Inc()

	rows, err := database.Pool.Query(context.Background(), `SELECT id, name, email FROM users`)
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		log.Println("Get All Users Error:", err)
		metrics.ApiFailures.WithLabelValues("GetAllUsers").Inc()
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			log.Println("Error scanning user:", err)
			continue
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
	metrics.ApiSuccesses.WithLabelValues("GetAllUsers").Inc()
	metrics.ApiRequestDuration.WithLabelValues("GetAllUsers").Observe(time.Since(start).Seconds())
}
