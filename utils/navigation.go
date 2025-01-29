package utils

import "fyne.io/fyne/v2"

func navigateTo(w fyne.Window, content fyne.CanvasObject) {
	w.SetContent(content)
}
