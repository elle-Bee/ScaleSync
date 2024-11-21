package api

import (
	"ScaleSync/pkg/database"
	"ScaleSync/pkg/metrics"
	"ScaleSync/pkg/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	timer := prometheus.NewTimer(metrics.ApiRequestDuration.WithLabelValues("login_user"))
	defer timer.ObserveDuration()

	metrics.ApiRequests.WithLabelValues("login_user").Inc()

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		metrics.ApiFailures.WithLabelValues("login_user").Inc()
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	name := user.Email
	password := user.Password

	query := `Select id, name, email, password FROM users WHERE email = ` + `'` + name + `'`
	fmt.Println(query)
	err := database.Pool.QueryRow(context.Background(), query).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		http.Error(w, "Email or Password is wrong", http.StatusNotFound)
		return
	}

	if database.CheckHash(password, user.Password) != nil {
		metrics.ApiFailures.WithLabelValues("login_user").Inc()
		http.Error(w, "Email or Password is wrong", http.StatusNotFound)
		return
	}
	var User_log models.User_login
	User_log.ID = user.ID
	User_log.Name = user.Name
	User_log.Email = user.Email
	User_log.Session = true

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(User_log); err != nil {
		metrics.ApiFailures.WithLabelValues("login_user").Inc()
		http.Error(w, "Failed to encode user login data", http.StatusInternalServerError)
	}
}
