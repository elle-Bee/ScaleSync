package app

import (
	"image/color"
	"log"

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
	emailEntry.SetPlaceHolder("Enter userID")

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter userID")

	IDEntry := widget.NewEntry()
	IDEntry.SetPlaceHolder("Enter userID")

	PasswordEntry := widget.NewPasswordEntry()
	PasswordEntry.SetPlaceHolder("Enter password")

	sign_up := widget.NewButton("Sign Up", func() {
		log.Println("Create User")
	})

	content := container.NewVBox(SmallSpacer, AppName, Spacer, emailEntry, nameEntry, IDEntry, PasswordEntry, SmallSpacer, sign_up)
	win.SetContent(content)

}
