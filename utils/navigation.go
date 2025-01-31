package utils

import "fyne.io/fyne/v2"

var (
	MainLayout           func(w fyne.Window) *fyne.Container
	StockExchangeScreen  func(w fyne.Window) *fyne.Container
	CurrencyQuotesScreen func(w fyne.Window) *fyne.Container
	CryptoQuotesScreen   func(w fyne.Window) *fyne.Container
)
