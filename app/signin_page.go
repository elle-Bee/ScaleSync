package app

import (
	"ScaleSync/pkg/models"
	"bytes"
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// Main UI for the app's home page
func ShowSignInPage(win fyne.Window) {
	win.Resize(fyne.NewSize(400, 550))
	// Define components
	smallSpacer := canvas.NewText(" ", color.White)
	smallSpacer.TextSize = 15

	appName := canvas.NewText("ScaleSync", color.White)
	appName.TextSize = 45
	appName.TextStyle.Bold = true
	appName.Alignment = fyne.TextAlignCenter

	largeSpacer := canvas.NewText(" ", color.White)
	largeSpacer.TextSize = 40

	// Create input fields for username and Password
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter username")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Enter password")

	// Sign-in button functionality
	signIn := widget.NewButton("Sign In", func() {
		loginUser(nameEntry.Text, passwordEntry.Text, win)
	})

	// Create "OR" divider
	or := canvas.NewText("----------------- OR -----------------", color.White)
	or.TextStyle.Monospace = true
	or.Alignment = fyne.TextAlignCenter

	// Sign-up button
	signUp := widget.NewButton("Sign Up", func() {
		showSignUpPage(win) // Navigate to the sign-up page
	})

	// content layout
	content := container.NewVBox(
		smallSpacer,
		appName,
		largeSpacer,
		nameEntry,
		passwordEntry,
		smallSpacer,
		signIn,
		smallSpacer,
		or,
		smallSpacer,
		signUp,
	)

	// Set content as the content of the window
	win.SetContent(content)
}

// Login user function to validate credentials with the server
func loginUser(username, password string, win fyne.Window) {
	if username == "" || password == "" {
		dialog.ShowError(fmt.Errorf("please enter both username and password"), win)
		return
	}

	// Prepare request body
	user := models.User{
		Name:     username,
		Password: password,
	}
	jsonData, err := json.Marshal(user)
	if err != nil {
		dialog.ShowError(fmt.Errorf("failed to encode user data: %v", err), win)
		return
	}

	// Make a POST request to the login endpoint
	resp, err := http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		dialog.ShowError(fmt.Errorf("error logging in: %v", err), win)
		return
	}
	defer resp.Body.Close()

	// Handle response from the server
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		dialog.ShowError(fmt.Errorf("error reading response: %v", err), win)
		return
	}

	// If the response status is OK, proceed to the next step
	if resp.StatusCode == http.StatusOK {
		var userLogin models.User_login
		err = json.Unmarshal(body, &userLogin)
		if err != nil {
			dialog.ShowError(fmt.Errorf("error decoding response: %v", err), win)
			return
		}
		// Login successful
		dialog.ShowInformation("Login Success", "You have successfully logged in.", win)
		fmt.Printf("Logged in user: %+v\n", userLogin)

		// Proceed to next page or dashboard
		ShowDashboardPage(win, userLogin)

	} else {
		dialog.ShowError(fmt.Errorf("invalid login credentials"), win)
	}
}
