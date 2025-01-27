package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/jphalexandrino/FXNow-Go/cmd"
)

func BuildMainLayout() *fyne.Container {
	// Menu buttons
	btnCoin := widget.NewButton("Currency quotes", func() {
		// Abrir tela de Moedas
		cmd.BuildCoinScreen()
	})
	btnCrypto := widget.NewButton("Cryptocurrency Quote", func() {
		// Abrir tela de crypto
	})
	btnStock := widget.NewButton("stock exchange", func() {
		// Abrir tela de Ações
	})

	// Layout principal
	return container.NewVBox(
		widget.NewLabelWithStyle("FX Now", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		btnCoin,
		btnCrypto,
		btnStock,
	)
}
