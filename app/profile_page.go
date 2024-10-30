package app

import (
	"ScaleSync/pkg/models"

	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowProfilePage(userLogin models.User_login) fyne.CanvasObject {
	title := widget.NewLabel("Profile Page")
	title.Alignment = fyne.TextAlignCenter

	// Display the user's name
	userName := canvas.NewText(userLogin.Name, color.White)
	userName.TextSize = 30
	userName.TextStyle.Bold = true
	userName.Alignment = fyne.TextAlignCenter

	// Display the user's email
	userEmail := canvas.NewText(userLogin.Email, color.White)
	userEmail.TextSize = 20
	userEmail.Alignment = fyne.TextAlignCenter

	// Layout the profile components
	return container.NewVBox(
		title,
		userName,
		userEmail,
	)
}
