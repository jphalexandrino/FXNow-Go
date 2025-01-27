package cmd

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BuildCoinScreen() *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("Tela de Moedas"),
		widget.NewButton("Voltar", func() {
			// Continuar...
		}),
	)
}
