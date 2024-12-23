package app

import (
	"ScaleSync/pkg/models"
	"bytes"
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
)

// Function to create a new user
func createUser(name, email, password string, win fyne.Window) {
	// Prepare user data to send in POST request
	userData := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	// Serialize user data into JSON format
	jsonData, err := json.Marshal(userData)
	if err != nil {
		dialog.ShowError(fmt.Errorf("error serializing user data: %v", err), win)
		return
	}

	// Send POST request to create a user
	resp, err := http.Post("http://localhost:8080/users", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		dialog.ShowError(fmt.Errorf("network error: %v", err), win)
		return
	}
	defer resp.Body.Close() // Ensure response body is closed to prevent resource leaks

	// Check response status code
	switch resp.StatusCode {
	case http.StatusOK:
		dialog.ShowInformation("Success", "User created successfully you are now being logged in", win)

		// Read the response body
		Body, err := io.ReadAll(resp.Body)
		if err != nil {
			dialog.ShowError(fmt.Errorf("error reading response body: %v", err), win)
			return
		}

		var userLogin models.User_login
		fmt.Printf("Logged in user: %+v\n", userLogin)
		// Unmarshal the JSON response into userLogin struct
		err = json.Unmarshal(Body, &userLogin)
		if err != nil {
			dialog.ShowError(fmt.Errorf("error decoding response: %v", err), win)
			return
		}

		// Proceed to next page or dashboard
		ShowDashboardPage(win, userLogin)
	case http.StatusBadRequest:
		dialog.ShowError(fmt.Errorf("bad request: invalid user data"), win)
	case http.StatusConflict:
		dialog.ShowError(fmt.Errorf("user already exists"), win)
	case http.StatusInternalServerError:
		dialog.ShowError(fmt.Errorf("server error: please try again later"), win)
	default:
		dialog.ShowInformation("Success", "User created successfully you are now being logged in", win)
	}
}

// ShowSignUpPage displays the sign-up interface
func showSignUpPage(win fyne.Window) {
	log.Println("Sign Up Tapped!")

	// UI Components
	smallSpacer := canvas.NewText(" ", color.White)
	smallSpacer.TextSize = 15

	appName := canvas.NewText("ScaleSync", color.White)
	appName.TextSize = 45
	appName.TextStyle.Bold = true
	appName.Alignment = fyne.TextAlignCenter

	largeSpacer := canvas.NewText(" ", color.White)
	largeSpacer.TextSize = 40

	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Enter Email")

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter Full Name")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Enter Password")

	signUpButton := widget.NewButton("Sign Up", func() {
		log.Println("Create User")
		createUser(nameEntry.Text, emailEntry.Text, passwordEntry.Text, win)
	})

	or := canvas.NewText("----------------- OR -----------------", color.White)
	or.TextStyle.Monospace = true
	or.Alignment = fyne.TextAlignCenter

	signInButton := widget.NewButton("Sign In", func() {
		ShowSignInPage(win) // Navigate to the sign-in page
	})

	// Set content layout
	content := container.NewVBox(
		smallSpacer,
		appName,
		largeSpacer,
		emailEntry,
		nameEntry,
		passwordEntry,
		smallSpacer,
		signUpButton,
		smallSpacer,
		or,
		smallSpacer,
		signInButton,
	)

	// Set content as the content of the window
	win.SetContent(content)
}
