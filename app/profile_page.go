package app

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowProfilePage() {
	return container.NewVBox(
		widget.NewLabel("Profile Page"),
		// Add more profile-related widgets here
	)
}
