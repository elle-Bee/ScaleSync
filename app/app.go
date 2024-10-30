package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func App() {
	myApp := app.New()
	win := myApp.NewWindow("ScaleSync")
	win.Resize(fyne.NewSize(400, 550))

	// Show the main page initially
	ShowSignInPage(win)

	win.ShowAndRun() // Show and run the window
}
