package app

import (
	"ScaleSync/pkg/models"
	"ScaleSync/pkg/repository"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func ShowDashboardPage(win fyne.Window, userLogin models.User_login) {
	win.Resize(fyne.NewSize(800, 550))

	warehouseRepo := &repository.WarehouseRepository{DB: database}

	// Initial content area
	contentArea := container.NewVBox(ShowHomePage(win, userLogin, warehouseRepo)) // Starts at home page

	// Sidebar with navigation options
	sidebar := container.NewVBox(
		widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
			contentArea.RemoveAll()
			contentArea.Add(ShowHomePage(win, userLogin, warehouseRepo))
			contentArea.Refresh()
		}),
		widget.NewButtonWithIcon("Dashboard", theme.ComputerIcon(), func() {
			contentArea.RemoveAll()
			contentArea.Add(ShowTablePage(win, userLogin, warehouseRepo))
			contentArea.Refresh()
		}),
		widget.NewButtonWithIcon("Profile", theme.AccountIcon(), func() {
			contentArea.RemoveAll()
			contentArea.Add(ShowProfilePage(win, userLogin))
			contentArea.Refresh()
		}),
	)

	// Layout that combines the sidebar and content area
	mainLayout := container.NewHBox(sidebar, contentArea)
	win.SetContent(mainLayout)
}
