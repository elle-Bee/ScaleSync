package app

import (
	"image/color"

	//"ScaleSync/app/sign_up"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func showMainPage(win fyne.Window) {

	SmallSpacer := canvas.NewText(" ", color.White)
	SmallSpacer.TextSize = 15

	AppName := canvas.NewText("ScaleSync", color.White)
	AppName.TextSize = 45
	AppName.TextStyle.Bold = true
	AppName.Alignment = fyne.TextAlignCenter

	Spacer := canvas.NewText(" ", color.White)
	Spacer.TextSize = 40

	IDEntry := widget.NewEntry()
	IDEntry.SetPlaceHolder("Enter userID")

	PasswordEntry := widget.NewPasswordEntry()
	PasswordEntry.SetPlaceHolder("Enter password")

	or := canvas.NewText("----------------- OR -----------------", color.White)
	or.TextStyle.Monospace = true
	or.Alignment = fyne.TextAlignCenter

	sign_up := widget.NewButton("Sign Up", func() {
		showSignUpPage(win)
	})

	content := container.NewVBox(SmallSpacer, AppName, Spacer, IDEntry, PasswordEntry, SmallSpacer, or, SmallSpacer, sign_up)

	win.SetContent(content)
}
