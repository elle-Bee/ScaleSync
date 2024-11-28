package app

import (
	"ScaleSync/pkg/database"
	"ScaleSync/pkg/models"
	"ScaleSync/pkg/repository"

	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func ShowDashboardPage(win fyne.Window, userLogin models.User_login) {
	win.Resize(fyne.NewSize(850, 580))

	warehouseRepo := &repository.WarehouseRepository{DB: database.InitDB()}

	homePage := ShowHomePage(win, userLogin, warehouseRepo)
	if homePage == nil {
		log.Println("Error: HomePage returned nil")
		return
	}
	contentArea := container.NewVBox(homePage)

	sidebar := container.NewVBox(
		widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
			contentArea.RemoveAll()
			homePage := ShowHomePage(win, userLogin, warehouseRepo)
			if homePage != nil {
				contentArea.Add(homePage)
				contentArea.Refresh()
			} else {
				log.Println("Error: HomePage returned nil on button click")
			}
		}),
		widget.NewButtonWithIcon("Dashboard", theme.ComputerIcon(), func() {
			contentArea.RemoveAll()
			tablePage := ShowTablePage(win, userLogin, warehouseRepo)
			if tablePage != nil {
				contentArea.Add(tablePage)
				contentArea.Refresh()
			} else {
				log.Println("Error: TablePage returned nil on button click")
			}
		}),
		widget.NewButtonWithIcon("Profile", theme.AccountIcon(), func() {
			contentArea.RemoveAll()
			profilePage := ShowProfilePage(win, userLogin)
			if profilePage != nil {
				contentArea.Add(profilePage)
				contentArea.Refresh()
			} else {
				log.Println("Error: ProfilePage returned nil on button click")
			}
		}),
	)

	mainLayout := container.NewHBox(sidebar, contentArea)
	win.SetContent(mainLayout)
}
