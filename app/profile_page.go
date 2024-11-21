package app

import (
	"ScaleSync/pkg/models"
	"log"

	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ShowProfilePage(win fyne.Window, userLogin models.User_login) fyne.CanvasObject {

	// Display the user's name
	userName := canvas.NewText(" Hello "+userLogin.Name+" !", color.White)
	userName.TextSize = 30
	userName.TextStyle.Bold = true

	log.Printf("Fetching for admin ID: %d", userLogin.ID)

	// Display the user's email
	userEmail := canvas.NewText("   Email: "+userLogin.Email, color.White)
	userEmail.TextSize = 15

	LargeSpacer := canvas.NewText(" ", color.White)
	LargeSpacer.TextSize = 40

	logoutButton := widget.NewButton("Log Out", func() {
		userLogin = models.User_login{} // Clear user information
		logout(win)
	})

	// Layout the profile components
	return container.NewVBox(
		userName,
		userEmail,
		LargeSpacer,
		logoutButton,
	)
}

func logout(win fyne.Window) {
	dialog.ShowInformation("Logout", "You have been logged out successfully.", win)
	ShowSignInPage(win)
}
