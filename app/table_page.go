package app

import (
	"ScaleSync/pkg/models"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowTablePage(win fyne.Window, userLogin models.User_login) fyne.CanvasObject {

	text := canvas.NewText("Select the warehouses within your jurisdiction for which you would like to view the data of :", color.White)
	// Display the user's email
	userEmail := canvas.NewText("Email: "+userLogin.Email, color.White)
	userEmail.TextSize = 15

	logoutButton := widget.NewButton("Log Out", func() {
		userLogin = models.User_login{} // Clear user information
		logout(win)
	})

	// Layout the profile components
	return container.NewVBox(
		title,
		userName,
		userEmail,
		logoutButton,
	)
}
