package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/jphalexandrino/FXNow-Go/theme"
	"github.com/jphalexandrino/FXNow-Go/ui"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get working directory:", err)
		return
	}
	fmt.Println("Current working directory:", wd)

	// Create the application instance
	a := app.New()
	w := a.NewWindow("Fx Now - By: João Pedro Hack Alexandrino")

	// Load the application icon
	icon, err := fyne.LoadResourceFromPath("assets/icons/Jp-logo.png")
	if err != nil {
		fmt.Println("Failed to load icon:", err)
		return
	}
	a.SetIcon(icon)
	w.SetIcon(icon)

	a.Settings().SetTheme(&theme.CustomTheme{})
	w.Resize(fyne.NewSize(800, 500))
	w.SetContent(ui.BuildMainLayout())

	w.ShowAndRun()
}
