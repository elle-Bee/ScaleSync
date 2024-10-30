package app

import (
	"ScaleSync/pkg/models"
	"ScaleSync/pkg/repository"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowTablePage(win fyne.Window, userLogin models.User_login, warehouseRepo *repository.WarehouseRepository) fyne.CanvasObject {
	warehouses, err := repository.WarehouseRepository.GetWarehousesByAdminID(userLogin.ID)
	if err != nil {
		log.Println("Error fetching warehouses:", err)
		return widget.NewLabel("Error fetching warehouses") // Return an error message
	}

	var checkboxes []fyne.CanvasObject
	for _, warehouse := range warehouses {
		// Creates a checkbox for each warehouse
		check := widget.NewCheck(warehouse.Location, func(value bool) {
			log.Printf("Checkbox for %s set to %v\n", warehouse.Location, value)
		})
		checkboxes = append(checkboxes, check)
	}

	// Return a VBox containing the widgets
	return container.NewVBox(checkboxes...)
}
