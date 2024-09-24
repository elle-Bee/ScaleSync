package app

import (
	"ScaleSync/pkg/api"
	"ScaleSync/pkg/database"
	"ScaleSync/pkg/models"
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"log"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/gorilla/mux"
)

// Main UI for the app's home page
func showMainPage(win fyne.Window) {
	// Define components
	smallSpacer := canvas.NewText(" ", color.White)
	smallSpacer.TextSize = 15

	appName := canvas.NewText("ScaleSync", color.White)
	appName.TextSize = 45
	appName.TextStyle.Bold = true
	appName.Alignment = fyne.TextAlignCenter

	largeSpacer := canvas.NewText(" ", color.White)
	largeSpacer.TextSize = 40

	// Create input fields for ID and Password
	IDEntry := widget.NewEntry()
	IDEntry.SetPlaceHolder("Enter userID")

	PasswordEntry := widget.NewPasswordEntry()
	PasswordEntry.SetPlaceHolder("Enter password")

	// Sign-in button functionality
	signIn := widget.NewButton("Sign In", func() {
		loginUser(IDEntry.Text, PasswordEntry.Text, win)
	})

	// Create "OR" divider
	or := canvas.NewText("----------------- OR -----------------", color.White)
	or.TextStyle.Monospace = true
	or.Alignment = fyne.TextAlignCenter

	// Sign-up button
	signUp := widget.NewButton("Sign Up", func() {
		showSignUpPage(win) // Navigate to the sign-up page
	})

	// Form layout
	form := container.NewVBox(
		smallSpacer,
		appName,
		largeSpacer,
		IDEntry,
		PasswordEntry,
		smallSpacer,
		signIn,
		smallSpacer,
		or,
		smallSpacer,
		signUp,
	)

	// Set form as the content of the window
	win.SetContent(form)
}

// Initialize HTTP server for user management and login
func Login() {
	database.Pool = database.InitDB() // Initialize the database connection
	defer database.Pool.Close()

	r := mux.NewRouter()

	// User routes
	r.HandleFunc("/users", api.CreateUser).Methods("POST")
	r.HandleFunc("/users", api.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", api.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", api.UpdateUser).Methods("PATCH")
	r.HandleFunc("/users/{id}", api.DeleteUser).Methods("DELETE")

	// Login route
	r.HandleFunc("/login", api.LoginUser).Methods("POST")

	// Start the HTTP server
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Fetch and display all users from the server
func fetchUsers(win fyne.Window) {
	resp, err := http.Get("http://localhost:8080/users")
	if err != nil {
		dialog.ShowError(fmt.Errorf("Error fetching users: %v", err), win)
		return
	}
	defer resp.Body.Close()

	// If the response status is OK, process the data
	if resp.StatusCode == http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		var users []models.User
		json.Unmarshal(body, &users)

		// Create a string representation of the users
		userList := ""
		for _, user := range users {
			userList += fmt.Sprintf("Name: %s, Email: %s\n", user.Name, user.Email)
		}

		// Display user information
		dialog.ShowInformation("Users", userList, win)
	} else {
		dialog.ShowError(fmt.Errorf("Could not fetch users: %v", resp.StatusCode), win)
	}
}

// Login user function to validate credentials
func loginUser(userID, password string, win fyne.Window) {
	if userID == "" || password == "" {
		dialog.ShowError(fmt.Errorf("Please enter both User ID and Password"), win)
		return
	}

	// You can implement your login logic here, for example:
	resp, err := http.PostForm("http://localhost:8080/login", map[string][]string{
		"userID":   {userID},
		"password": {password},
	})

	if err != nil {
		dialog.ShowError(fmt.Errorf("Error logging in: %v", err), win)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// Login successful, proceed to next page or show success
		dialog.ShowInformation("Login Success", "You have successfully logged in.", win)
	} else {
		dialog.ShowError(fmt.Errorf("Invalid login credentials"), win)
	}
}
