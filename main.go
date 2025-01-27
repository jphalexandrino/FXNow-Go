package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	fmt.Println("Starting Fyne")

	a := app.New()
	w := a.NewWindow("My new title")

	w.Resize(fyne.NewSize(800, 500)) // Resizing the window

	img := canvas.NewImageFromFile("images/my-logo.png")
	w.SetContent(img)

	w.ShowAndRun()

}
