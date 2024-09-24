// package main

// import (
// 	"ScaleSync/pkg/models"
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"sync"

// 	"ScaleSync/pkg/api"
// 	"ScaleSync/pkg/database"

// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/dialog"
// 	"fyne.io/fyne/v2/widget"
// 	"github.com/gorilla/mux"
// 	"github.com/joho/godotenv"
// )

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

// func Login() {
// 	database.Pool = database.InitDB() // Initialize the database connection
// 	defer database.Pool.Close()

// 	r := mux.NewRouter()

// 	r.HandleFunc("/users", api.CreateUser).Methods("POST")
// 	r.HandleFunc("/users", api.GetAllUsers).Methods("GET")

// 	r.HandleFunc("/users/{id}", api.GetUser).Methods("GET")
// 	r.HandleFunc("/users/{id}", api.UpdateUser).Methods("PATCH")
// 	r.HandleFunc("/users/{id}", api.DeleteUser).Methods("DELETE")

// 	r.HandleFunc("/login", api.LoginUser).Methods("POST")

// 	log.Println("Server running at http://localhost:8080")
// 	log.Fatal(http.ListenAndServe(":8080", r))
// }

// func createUser(name, email, password string) {
// 	userData := models.User{
// 		Name:     name,
// 		Email:    email,
// 		Password: password,
// 	}
// 	jsonData, _ := json.Marshal(userData)

// 	resp, err := http.Post("http://localhost:8080/users", "application/json", bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		dialog.ShowError(err, nil)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode == http.StatusOK {
// 		dialog.ShowInformation("Success", "User created successfully", nil)
// 	} else {
// 		dialog.ShowError(fmt.Errorf("Failed to create user"), nil)
// 	}
// }

// func fetchUsers() {
// 	resp, err := http.Get("http://localhost:8080/users")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode == http.StatusOK {
// 		body, _ := io.ReadAll(resp.Body)
// 		var users []models.User
// 		json.Unmarshal(body, &users)
// 		userList := ""
// 		for _, user := range users {
// 			userList += fmt.Sprintf("Name: %s, Email: %s\n", user.Name, user.Email)
// 		}
// 		dialog.ShowInformation("Users", userList, nil)
// 	} else {
// 		dialog.ShowError(fmt.Errorf("Could not fetch users"), nil)
// 	}
// }

// func startGUI() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("User Management")

// 	nameEntry := widget.NewEntry()
// 	nameEntry.SetPlaceHolder("Enter name")
// 	emailEntry := widget.NewEntry()
// 	emailEntry.SetPlaceHolder("Enter email")
// 	passwordEntry := widget.NewPasswordEntry()
// 	passwordEntry.SetPlaceHolder("Enter password")

// 	createButton := widget.NewButton("Create User", func() {
// 		createUser(nameEntry.Text, emailEntry.Text, passwordEntry.Text)
// 	})
// 	fetchButton := widget.NewButton("Fetch Users", fetchUsers)

// 	form := container.NewVBox(
// 		widget.NewLabel("Name:"),
// 		nameEntry,
// 		widget.NewLabel("Email:"),
// 		emailEntry,
// 		widget.NewLabel("Password:"),
// 		passwordEntry,
// 		createButton,
// 		fetchButton,
// 	)

// 	myWindow.SetContent(form)
// 	myWindow.ShowAndRun()
// }

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)

// 	// Start the HTTP server in a goroutine
// 	go func() {
// 		defer wg.Done()
// 		Login() // Start the server
// 	}()

// 	// Start the GUI in the main goroutine
// 	startGUI()

// 	wg.Wait()
// }

package main

import (
	"ScaleSync/app"
)

func main() {
	app.UI()
}
