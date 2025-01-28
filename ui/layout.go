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

// BuildMainLayout Builds the main layout of the application
func BuildMainLayout() *fyne.Container {
	// Section: Buttons
	btnCoin := widget.NewButton("Currency quotes", func() {})
	btnCrypto := widget.NewButton("Cryptocurrency Quote", func() {})
	btnStock := widget.NewButton("Stock exchange", func() {})

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
		urlGitHub := "https://github.com/jphalexandrino"
		parsedURL, err := url.Parse(urlGitHub)
		if err != nil {
			fmt.Println("Error converting string to URL:", err)
			return
		}
		_ = fyne.CurrentApp().OpenURL(parsedURL)
	})
	githubIcon.SetIcon(loadIconFromFile("assets/icons/github.png"))

	globeIcon := widget.NewButton("", func() {
		urlPortfolio := "https://jphalexandrino.vercel.app"
		parsedURL, err := url.Parse(urlPortfolio)
		if err != nil {
			fmt.Println("Error converting string to URL:", err)
			return
		}
		_ = fyne.CurrentApp().OpenURL(parsedURL)
	})
	globeIcon.SetIcon(loadIconFromFile("assets/icons/Logo-site.png"))

	linkedinIcon := widget.NewButton("", func() {
		urlLinkedin := "https://www.linkedin.com/in/joão-pedro-alexandrino/"
		parsedURL, err := url.Parse(urlLinkedin)
		if err != nil {
			fmt.Println("Error converting string to URL:", err)
			return
		}
		_ = fyne.CurrentApp().OpenURL(parsedURL)
	})
	linkedinIcon.SetIcon(loadIconFromFile("assets/icons/LinkedIn-logo.png"))

	// Footer icon container
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
	headerContainer := container.New(
		layout.NewVBoxLayout(),
		container.NewPadded(FXLogo),
	)

	// Section: Main Layout
	mainContent := container.NewBorder(
		nil,
		footer,
		nil,
		nil,
		container.NewVBox(
			headerContainer,
			buttonsLayout,
		),
	)

	return mainContent
}
