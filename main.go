package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/jphalexandrino/FXNow-Go/theme"
	"github.com/jphalexandrino/FXNow-Go/ui"
)

func main() {
	fmt.Println("Starting Fyne")

	a := app.New()
	w := a.NewWindow("Fx Now - By: João Pedro Hack Alexandrino")

	// SetTheme
	a.Settings().SetTheme(&theme.CustomTheme{})

	w.Resize(fyne.NewSize(800, 500)) // Resizing the window
	w.SetContent(ui.BuildMainLayout())

	w.ShowAndRun()

}
