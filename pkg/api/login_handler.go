package api

import (
	"ScaleSync/pkg/database"
	"ScaleSync/pkg/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	name := user.Name
	password := user.Password

	query := `Select name, email, password FROM users WHERE name = ` + `'` + name + `'`
	fmt.Println(query)
	err := database.Pool.QueryRow(context.Background(), query).Scan(&user.Name, &user.Email, &user.Password)

	if err != nil {
		http.Error(w, "Username or Password is wrong", http.StatusNotFound)
		return
	}

	if database.CheckHash(password, user.Password) != nil {
		fmt.Println("YOU HAVE ENETERD WRONG PASSWD")
		http.Error(w, "Username or Password is wrong", http.StatusNotFound)
		return
	}
	var User_log models.User_login
	User_log.ID = user.ID
	User_log.Name = user.Name
	User_log.Email = user.Email
	User_log.Session = true
	if user.Name == "Arnav" {
		User_log.Privelages = true
	}

	json.NewEncoder(w).Encode(User_log)
}
