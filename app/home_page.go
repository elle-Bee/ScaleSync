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
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func ShowHomePage(win fyne.Window, userLogin models.User_login, warehouseRepo *repository.WarehouseRepository) fyne.CanvasObject {

	text := canvas.NewText("   Warehouses under your jurisdiction: : "+userLogin.Email, color.White)
	text.TextSize = 15

	warehouses, err := warehouseRepo.GetWarehousesByAdminID(userLogin.ID)
	if err != nil {
		log.Println("Error fetching warehouses: ", err)
		return nil
	}

	// Convert warehouse locations to a bindable string list
	warehouseNames := []string{}
	for _, warehouse := range warehouses {
		warehouseNames = append(warehouseNames, fmt.Sprintf("%s - Capacity: %d/%d", warehouse.Location, warehouse.CurrentCapacity, warehouse.TotalCapacity))
	}
	data := binding.BindStringList(&warehouseNames)

	// Create List widget with data binding
	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template") // Template label
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String)) // Bind data item to label
		})

	return container.NewVBox(
		text,
		list,
	)
}
