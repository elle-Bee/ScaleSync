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
)

func ShowHomePage(win fyne.Window, userLogin models.User_login, warehouseRepo *repository.WarehouseRepository) fyne.CanvasObject {
	// Create a spacer
	Spacer := canvas.NewText(" ", color.White)
	Spacer.TextSize = 10

	// Create a header text
	text := canvas.NewText("   Warehouses under your jurisdiction:", color.White)
	text.TextSize = 15

	// Fetch details of warehouses by user ID
	log.Printf("Fetching warehouses for user ID: %d", userLogin.ID) // Correct log message
	warehouses, err := warehouseRepo.GetWarehousesByAdminID(userLogin.ID)
	if err != nil {
		log.Println("Error fetching warehouses: ", err)
		return nil
	}
	log.Printf("Fetched warehouses: %+v\n", warehouses)

	// // Convert warehouses to a bindable string list
	// warehouseNames := []string{}
	// for _, warehouse := range warehouses {
	// 	if warehouse.Location != "" {
	// 		warehouseNames = append(warehouseNames, fmt.Sprintf("%s - Capacity: %d/%d", warehouse.Location, warehouse.CurrentCapacity, warehouse.TotalCapacity))
	// 	}
	// }

	// // Create a bindable string list for the warehouse names
	// data := binding.BindStringList(&warehouseNames)

	// // Create List widget with data binding
	// list := widget.NewListWithData(data,
	// 	func() fyne.CanvasObject {
	// 		return widget.NewLabel("template") // Template label
	// 	},
	// 	func(i binding.DataItem, o fyne.CanvasObject) {
	// 		o.(*widget.Label).Bind(i.(binding.String)) // Bind data item to label
	// 	})

	// // Wrap the list in a scrollable container
	// scrollableList := container.NewScroll(list)

	// // Optional: Set a minimum size for the scrollable list
	// list.MinSize() // Set to appropriate values for your design

	warehouseNames := ""
	for _, warehouse := range warehouses {
		if warehouse.Location != "" {
			warehouseNames += fmt.Sprintf("- %s: Capacity %d / %d\n", warehouse.Location, warehouse.CurrentCapacity, warehouse.TotalCapacity)
		}
	}
	log.Print(warehouseNames)
	list := canvas.NewText("   "+warehouseNames, color.White)
	list.TextSize = 15

	// Return a vertical box containing the spacer, header, and scrollable list
	return container.NewVBox(
		Spacer,
		text,
		list,
	)
}
