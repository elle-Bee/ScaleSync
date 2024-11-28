package app

import (
	"ScaleSync/pkg/models"
	"ScaleSync/pkg/repository"
	"bytes"
	"fmt"
	"image/color"
	"log"
	"os/exec"

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

	genai_text := canvas.NewText("   Some AI insights", color.White)
	genai_text.TextSize = 15
	genai_text.TextStyle.Bold = true

	aiResponseText, err := RunPythonScript()
	if err != nil {
		log.Println("Error running Python script:", err)
		aiResponseText = "Error generating AI insights."
	}

	aiInsights := canvas.NewText(aiResponseText, color.White)
	aiInsights.TextSize = 10

	text := canvas.NewText("   Warehouse(s) under your jurisdiction:", color.White)
	text.TextSize = 15
	text.TextStyle.Bold = true

	log.Printf("Fetching warehouses for user ID: %d", userLogin.ID)
	warehouses, err := warehouseRepo.GetWarehousesByAdminID(userLogin.ID)
	if err != nil {
		log.Println("Error fetching warehouses: ", err)
		return container.NewVBox(
			Spacer,
			text,
			canvas.NewText("Error fetching warehouse data. Please try again later.", color.White),
		)
	}
	log.Printf("Fetched warehouses: %+v\n", warehouses)

	warehouseNames := []string{}
	if len(warehouses) == 0 {
		warehouseNames = append(warehouseNames, "No warehouses found for your jurisdiction.")
	} else {
		for _, warehouse := range warehouses {
			if warehouse.Location != "" {
				warehouseNames = append(warehouseNames, fmt.Sprintf("   %s - Capacity: %d/%d", warehouse.Location, warehouse.CurrentCapacity, warehouse.TotalCapacity))
			}
		}
	}

	data := binding.BindStringList(&warehouseNames)

	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template") // Template label
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String)) // Bind data item to label
		})

	scrollableList := container.NewScroll(list)
	scrollableList.SetMinSize(fyne.NewSize(600, 250))

	return container.NewVBox(
		Spacer,
		genai_text,
		smallSpacer,
		aiInsights,
		Spacer,
		text,
		smallSpacer,
		scrollableList,
	)
}

func RunPythonScript() (string, error) {
	cmd := exec.Command("python", "app/gemini_api.py")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
