package app

import (
	"ScaleSync/pkg/database"
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

	checkedWarehouses := []models.Warehouse{}

	// Create container for displaying warehouses
	warehouseCollectionContainer := container.NewVBox()
	for _, warehouse := range warehouses {
		// Create a checkbox for each warehouse
		warehouseCheck := widget.NewCheck(warehouse.Location, func(checked bool) {
			if checked {
				log.Printf("Warehouse %d (%s) selected", warehouse.Warehouse_ID, warehouse.Location)
				checkedWarehouses = append(checkedWarehouses, warehouse)
			} else {
				log.Printf("Warehouse %d (%s) deselected", warehouse.Warehouse_ID, warehouse.Location)
			}
		})
		// Customize checkbox label with more warehouse info if needed
		warehouseCheck.SetText(warehouse.Location)
		warehouseCollectionContainer.Add(warehouseCheck)
	}

	fetch := widget.NewButton("Fetch", func() {

		//text.Text = "" // Clear the text content
		text.Hide()
		smallSpacer.Hide()

		// Clears previous content
		warehouseCollectionContainer.Objects = warehouseCollectionContainer.Objects[:0]
		warehouseCollectionContainer.Resize(fyne.NewSize(600, 250))

		itemRepo := &repository.ItemRepository{DB: database.InitDB()}

		for _, checkedWarehouse := range checkedWarehouses {
			if checkedWarehouse.Location != "" {
				warehouseData := []string{}

				checkedItems, err := itemRepo.GetItemsByWarehouseID(checkedWarehouse.Warehouse_ID)
				if err != nil {
					log.Println("Error fetching warehouses: ", err)
					return
				}

				for _, checkedItem := range checkedItems {
					warehouseData = append(warehouseData, fmt.Sprintf("%s", checkedItem))
				}

				data := binding.BindStringList(&warehouseData)

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

				warehouseName := canvas.NewText(checkedWarehouse.Location, color.White)
				warehouseName.TextSize = 15

				header := canvas.NewText("   ItemID   Item Name   Category   Quantity   Unit Price   Total Price   Description", color.White)

				warehouseContainer := container.NewVBox(warehouseName, header, list)
				warehouseContainer.Resize(fyne.NewSize(600, 250))

				warehouseCollectionContainer.Add(warehouseContainer)
				warehouseCollectionContainer.Refresh()

			}
		}

		warehouseCollectionContainer.Refresh()
	})

	// Layout the profile components
	return container.NewVBox(
		Spacer,
		text,
		smallSpacer,
		warehouseCollectionContainer,
		smallSpacer,
		fetch,
	)
}
