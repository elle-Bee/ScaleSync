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
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if database.CheckHash(password, user.Password) != nil {
		fmt.Println("YOU HAVE ENETERD WRONG PASSWD")
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	var user_log models.User_login
	user_log.ID = user.ID
	user_log.Name = user.Name
	user_log.Email = user.Email

	json.NewEncoder(w).Encode(user_log)
}
