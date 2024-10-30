package app

import (
	"ScaleSync/pkg/models"
	"ScaleSync/pkg/repository"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowTablePage(win fyne.Window, userLogin models.User_login, warehouseRepo *repository.WarehouseRepository) fyne.CanvasObject {

	Spacer := canvas.NewText(" ", color.White)
	Spacer.TextSize = 10

	smallSpacer := canvas.NewText(" ", color.White)
	smallSpacer.TextSize = 5

	text := canvas.NewText("   Select the warehouse(s) within your jurisdiction for which you would like to view the data of :", color.White)
	text.TextSize = 15
	text.TextStyle.Bold = true

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
		warehouseCheck.SetText(warehouse.Location)
		warehouseContainer.Add(warehouseCheck)
	}

	// Layout the profile components
	return container.NewVBox(
		Spacer,
		text,
		smallSpacer,
		warehouseContainer,
	)
}
