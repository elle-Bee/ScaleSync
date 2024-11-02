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

	// Create a container for displaying warehouses
	warehouseCollectionContainer := container.NewVBox()
	warehouseCollectionContainer.MinSize()

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
		text.Hide()
		smallSpacer.Hide()

		// Clears previous content
		warehouseCollectionContainer.Objects = nil

		itemRepo := &repository.ItemRepository{DB: database.InitDB()}

		for _, checkedWarehouse := range checkedWarehouses {
			if checkedWarehouse.Location != "" {
				warehouseData := []string{}

				checkedItems, err := itemRepo.GetItemsByWarehouseID(checkedWarehouse.Warehouse_ID)
				if err != nil {
					log.Println("Error fetching items: ", err)
					return
				}

				for _, checkedItem := range checkedItems {
					warehouseData = append(warehouseData, fmt.Sprintf("%s", checkedItem))
				}

				data := binding.BindStringList(&warehouseData)

				//horizontalListContainer := container.NewHBox()
				var listOfHorizontalContainers []*container.Scroll

				for i := 0; i < data.Length(); i++ {
					horizontalListContainer := container.NewHBox()

					//Retrieve data
					itemData, err := data.GetValue(i)
					if err != nil {
						log.Println("Error getting item from data list:", err)
						continue
					}

					label := widget.NewLabel(itemData)
					label.Alignment = fyne.TextAlignCenter
					label.Resize(fyne.NewSize(150, 30)) // Set width and height for each label

					horizontalListContainer.Add(label)

					scrollableHorizontal := container.NewHScroll(horizontalListContainer)
					scrollableHorizontal.SetMinSize(fyne.NewSize(600, 50))
					listOfHorizontalContainers = append(listOfHorizontalContainers, scrollableHorizontal)
				}

				// A vertical container to hold all the scrollable horizontal containers
				verticalContainer := container.NewVBox()
				for _, hScrollContainer := range listOfHorizontalContainers {
					verticalContainer.Add(hScrollContainer)
				}

				warehouseName := canvas.NewText("  "+checkedWarehouse.Location, color.White)
				warehouseName.TextSize = 15
				warehouseName.TextStyle.Bold = true

				header := canvas.NewText("   ItemID      Item Name            Category        Quantity    Unit Price    Total Price", color.White)

				smallSpacer := canvas.NewText(" ", color.White)
				smallSpacer.TextSize = 5

				mediumSpacer := canvas.NewText(" ", color.White)
				mediumSpacer.TextSize = 10

				warehouseCollectionContainer.Add(warehouseName)
				warehouseCollectionContainer.Add(smallSpacer)
				warehouseCollectionContainer.Add(header)
				warehouseCollectionContainer.Add(verticalContainer)
				warehouseCollectionContainer.Add(mediumSpacer)
			}
		}
		warehouseCollectionContainer.MinSize()
		warehouseCollectionContainer.Refresh()
	})

	// Wrap the warehouse collection container in a vertical scroll container
	scrollableWarehouseCollection := container.NewVScroll(warehouseCollectionContainer)
	scrollableWarehouseCollection.SetMinSize(fyne.NewSize(600, 550))

	// Layout the profile components
	return container.NewVBox(
		Spacer,
		text,
		smallSpacer,
		scrollableWarehouseCollection, // Use the scrollable version here
		smallSpacer,
		fetch,
	)
}
