package app

import (
	"ScaleSync/pkg/models"
	"bytes"
	"encoding/json"
	"fmt"
	"image/color"
	"log"
	"net/http"

	"fyne.io/fyne/dialog"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func showSignUpPage(win fyne.Window) {
	log.Println("Sign Up Tapped!")

	SmallSpacer := canvas.NewText(" ", color.White)
	SmallSpacer.TextSize = 15

	AppName := canvas.NewText("ScaleSync", color.White)
	AppName.TextSize = 45
	AppName.TextStyle.Bold = true
	AppName.Alignment = fyne.TextAlignCenter

	Spacer := canvas.NewText(" ", color.White)
	Spacer.TextSize = 40

	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Enter Email")

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter Full Name")

	PasswordEntry := widget.NewPasswordEntry()
	PasswordEntry.SetPlaceHolder("Enter password")

	sign_up := widget.NewButton("Sign Up", func() {
		log.Println("Create User")
		createUser(nameEntry.Text, emailEntry.Text, PasswordEntry.Text)
	})

	content := container.NewVBox(SmallSpacer, AppName, Spacer, emailEntry, nameEntry, PasswordEntry, SmallSpacer, sign_up)
	win.SetContent(content)

}

func createUser(name, email, password string) {
	userData := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	jsonData, _ := json.Marshal(userData)

	resp, err := http.Post("http://localhost:8080/users", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		dialog.ShowError(err, nil)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		dialog.ShowInformation("Success", "User created successfully", nil)
	} else {
		dialog.ShowError(fmt.Errorf("Failed to create user"), nil)
	}
}
