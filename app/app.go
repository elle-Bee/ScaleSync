package app

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func UI() {
	myApp := app.New()
	win := myApp.NewWindow("ScaleSync")
	defer win.ShowAndRun() // Show and run the window

	win.Resize(fyne.NewSize(400, 550))

	small_spacer := canvas.NewText(" ", color.White)
	small_spacer.TextSize = 15

	text := canvas.NewText("ScaleSync", color.White)
	text.TextSize = 45
	text.TextStyle.Bold = true
	text.Alignment = fyne.TextAlignCenter

	spacer := canvas.NewText(" ", color.White)
	spacer.TextSize = 40

	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("Enter userID")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Enter password")

	or := canvas.NewText("----------------- OR -----------------", color.White)
	or.TextStyle.Monospace = true
	or.Alignment = fyne.TextAlignCenter

	sign_up := widget.NewButton("Sign Up", func() {
		log.Println("tapped")
	})

	content := container.NewVBox(small_spacer, text, spacer, idEntry, passwordEntry, small_spacer, or, small_spacer, sign_up)

	win.SetContent(content)
}
