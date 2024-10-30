package app

import (
	"ScaleSync/pkg/models"
	"ScaleSync/pkg/repository"
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowTablePage(win fyne.Window, userLogin models.User_login, warehouseRepo *repository.WarehouseRepository) fyne.CanvasObject {

	text := canvas.NewText("Select the warehouses within your jurisdiction for which you would like to view the data of :", color.White)

	warehouses, err := warehouseRepo.GetWarehousesByAdminID(userLogin.ID)
	if err != nil {
		log.Println("Error fetching warehouses: ", err)
		return nil
	}

	// Create container for displaying warehouses
	warehouseContainer := container.NewVBox()
	for _, warehouse := range warehouses {
		// Create a checkbox for each warehouse
		warehouseCheck := widget.NewCheck(warehouse.Location, func(checked bool) {
			if checked {
				log.Printf("Warehouse %d (%s) selected", warehouse.Warehouse_ID, warehouse.Location)
			} else {
				log.Printf("Warehouse %d (%s) deselected", warehouse.Warehouse_ID, warehouse.Location)
			}
		})
		// Customize checkbox label with more warehouse info if needed
		warehouseCheck.SetText(warehouse.Location + " - Capacity: " + fmt.Sprintf("%d/%d", warehouse.CurrentCapacity, warehouse.TotalCapacity))
		warehouseContainer.Add(warehouseCheck)
	}

	// Layout the profile components
	return container.NewVBox(
		text,
		warehouseContainer,
	)
}
