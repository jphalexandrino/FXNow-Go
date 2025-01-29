package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"net/url"
)

// Utility function to load an icon from a file
func loadIconFromFile(path string) fyne.Resource {
	icon, err := fyne.LoadResourceFromPath(path)
	if err != nil {
		fmt.Println("Error loading icon:", err)
		return nil
	}
	return icon
}

// BuildMainLayout builds the main layout of the application
func BuildMainLayout(w fyne.Window) *fyne.Container {
	// Section: Buttons with navigation
	btnCoin := widget.NewButton("Currency quotes", func() {
		w.SetContent(BuildCurrencyQuotesScreen(w)) // Navigate to the Currency Quotes screen
	})
	btnCrypto := widget.NewButton("Cryptocurrency Quote", func() {
		w.SetContent(BuildCryptoQuotesScreen(w)) // Navigate to the Cryptocurrency Quotes screen
	})
	btnStock := widget.NewButton("Stock exchange", func() {
		w.SetContent(BuildStockExchangeScreen(w)) // Navigate to the Stock Exchange screen
	})

	// Wrap buttons in fixed-size containers
	btnCoinContainer := container.New(layout.NewGridWrapLayout(fyne.NewSize(200, 30)), btnCoin)
	btnCryptoContainer := container.New(layout.NewGridWrapLayout(fyne.NewSize(200, 30)), btnCrypto)
	btnStockContainer := container.New(layout.NewGridWrapLayout(fyne.NewSize(200, 30)), btnStock)

	// VBox layout for buttons, centered
	buttonsLayout := container.NewVBox(
		container.NewCenter(btnCoinContainer),
		container.NewCenter(btnCryptoContainer),
		container.NewCenter(btnStockContainer),
	)

	// Section: Footer Icons
	githubIcon := widget.NewButton("", func() {
		openURL("https://github.com/jphalexandrino")
	})
	githubIcon.SetIcon(loadIconFromFile("assets/icons/github.png"))

	globeIcon := widget.NewButton("", func() {
		openURL("https://jphalexandrino.vercel.app")
	})
	globeIcon.SetIcon(loadIconFromFile("assets/icons/Logo-site.png"))

	linkedinIcon := widget.NewButton("", func() {
		openURL("https://www.linkedin.com/in/joão-pedro-alexandrino/")
	})
	linkedinIcon.SetIcon(loadIconFromFile("assets/icons/LinkedIn-logo.png"))

	// Footer icon container (aligned bottom-left)
	iconContainer := container.NewHBox(
		githubIcon,
		layout.NewSpacer(),
		globeIcon,
		layout.NewSpacer(),
		linkedinIcon,
	)

	footer := container.NewBorder(nil, nil, iconContainer, nil)

	// Section: Header
	FXLogo := widget.NewLabelWithStyle("FX Now", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	headerContainer := container.NewCenter(FXLogo)

	// Section: Main Layout
	mainContent := container.NewBorder(
		headerContainer, // Header at the top
		footer,          // Footer at the bottom
		nil,             // Left empty
		nil,             // Right empty
		container.NewVBox(
			buttonsLayout, // Centered buttons
		),
	)

	return mainContent
}

// BuildCurrencyQuotesScreen creates the Currency Quotes screen
func BuildCurrencyQuotesScreen(w fyne.Window) *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("Currency Quotes Screen"),
		widget.NewButton("Back", func() {
			w.SetContent(BuildMainLayout(w)) // Go back to the main layout
		}),
	)
}

// BuildCryptoQuotesScreen creates the Cryptocurrency Quotes screen
func BuildCryptoQuotesScreen(w fyne.Window) *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("Cryptocurrency Quotes Screen"),
		widget.NewButton("Back", func() {
			w.SetContent(BuildMainLayout(w)) // Go back to the main layout
		}),
	)
}

// BuildStockExchangeScreen creates the Stock Exchange screen
func BuildStockExchangeScreen(w fyne.Window) *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("Stock Exchange Screen"),
		widget.NewButton("Back", func() {
			w.SetContent(BuildMainLayout(w)) // Go back to the main layout
		}),
	)
}

// Utility function to open a URL
func openURL(link string) {
	parsedURL, err := url.Parse(link)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	_ = fyne.CurrentApp().OpenURL(parsedURL)
}
