package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Function to create the home page content
func showHomePage() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Home Page"),
		// Add more home-related widgets here
	)
}

func ShowDashboardPage(win fyne.Window) {
	win.Resize(fyne.NewSize(800, 550))
	// Initial content area
	contentArea := container.NewVBox(showHomePage()) // Start with home page

	// Create sidebar with navigation options
	sidebar := container.NewVBox(
		widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
			contentArea.RemoveAll()
			contentArea.Add(showHomePage())
			contentArea.Refresh()
		}),
		widget.NewButtonWithIcon("Profile", theme.AccountIcon(), func() {
			contentArea.RemoveAll()
			contentArea.Add(ShowProfilePage())
			contentArea.Refresh()
		}),
	)

	// Create a layout that combines the sidebar and content area
	mainLayout := container.NewHBox(sidebar, contentArea)

	// Set the content of the window
	win.SetContent(mainLayout)
}
