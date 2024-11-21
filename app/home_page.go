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

	Spacer := canvas.NewText(" ", color.White)
	Spacer.TextSize = 10

	smallSpacer := canvas.NewText(" ", color.White)
	smallSpacer.TextSize = 5

	text := canvas.NewText("   Warehouse(s) under your jurisdiction:", color.White)
	text.TextSize = 15
	text.TextStyle.Bold = true

	// Fetch details of warehouses by user name
	log.Printf("Fetching warehouses for user ID: %d", userLogin.ID)
	warehouses, err := warehouseRepo.GetWarehousesByAdminID(userLogin.ID)
	if err != nil {
		log.Println("Error fetching warehouses: ", err)
		return nil
	}
	log.Printf("Fetched warehouses: %+v\n", warehouses)

	// Convert warehouses to a bindable string list
	warehouseNames := []string{}
	for _, warehouse := range warehouses {
		if warehouse.Location != "" {
			warehouseNames = append(warehouseNames, fmt.Sprintf("   %s - Capacity: %d/%d", warehouse.Location, warehouse.CurrentCapacity, warehouse.TotalCapacity))
		}
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
	// Wrap the list in a container with fixed size
	scrollableList := container.NewScroll(list)
	scrollableList.SetMinSize(fyne.NewSize(600, 250))

	return container.NewVBox(
		Spacer,
		text,
		smallSpacer,
		scrollableList,
	)
}
