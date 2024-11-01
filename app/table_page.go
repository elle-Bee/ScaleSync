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

	fetch := widget.NewButton("Fetch", func() {

		//text.Text = "" // Clear the text content
		text.Hide()
		smallSpacer.Hide()

		// Clears previous content
		warehouseContainer.Objects = warehouseContainer.Objects[:0]

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

		warehouseContainer.Add(list)
		warehouseContainer.Refresh()
	})

	// Layout the profile components
	return container.NewVBox(
		Spacer,
		text,
		smallSpacer,
		warehouseContainer,
		smallSpacer,
		fetch,
	)
}
