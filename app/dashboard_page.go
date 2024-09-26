package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func ShowDashboardPage(win fyne.Window) {
	win.Resize(fyne.NewSize(800, 550))

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")),
		container.NewTabItemWithIcon("Profile", theme.AccountIcon(), widget.NewLabel("Profile")),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	win.SetContent(tabs)

	win.SetContent(tabs)
}
