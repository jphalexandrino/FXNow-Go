package cmd

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/jphalexandrino/FXNow-Go/utils"
)

// BuildCryptoQuotesScreen creates the Cryptocurrency Quotes screen
func BuildCryptoQuotesScreen(w fyne.Window) *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("Cryptocurrency Quotes Screen"),
		widget.NewButton("Back", func() {
			w.SetContent(utils.MainLayout(w)) // Go back to the main layout
		}),
	)
}
