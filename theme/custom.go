package theme

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

// CustomTheme defines a custom theme
type CustomTheme struct{}

// Implementing the fyne.Theme interface

func (c *CustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{R: 30, G: 30, B: 30, A: 255} // Dark background color
	case theme.ColorNameButton:
		return color.RGBA{R: 0, G: 100, B: 0, A: 255} // Green buttons
	case theme.ColorNameDisabledButton:
		return color.RGBA{R: 50, G: 50, B: 50, A: 255} // Disabled button color
	case theme.ColorNameInputBackground:
		return color.RGBA{R: 50, G: 50, B: 50, A: 255} // Input field background
	case theme.ColorNamePlaceHolder:
		return color.RGBA{R: 150, G: 150, B: 150, A: 255} // Placeholder text
	default:
		// Fallback for any unhandled color names
		return theme.DefaultTheme().Color(name, variant)
	}
}

func (c *CustomTheme) Font(style fyne.TextStyle) fyne.Resource {
	// Return the default font (or provide custom fonts here)
	return theme.DefaultTheme().Font(style)
}

func (c *CustomTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	// Return default icons (or replace with custom icons here)
	return theme.DefaultTheme().Icon(name)
}

func (c *CustomTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameText:
		return 14 // Custom text size
	case theme.SizeNamePadding:
		return 8 // Custom padding
	default:
		// Fallback for any unhandled size names
		return theme.DefaultTheme().Size(name)
	}
}
