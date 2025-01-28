package theme

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

// CustomTheme defines a custom theme for the application
type CustomTheme struct{}

// Color overrides theme.Color to define custom colors
func (c *CustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{R: 34, G: 34, B: 34, A: 255} // Dark gray background
	case theme.ColorNameButton:
		return color.RGBA{R: 39, G: 118, B: 245, A: 255} // Green button by default
	case theme.ColorNameDisabledButton:
		return color.RGBA{R: 60, G: 60, B: 60, A: 255} // Dark gray for disabled buttons
	case theme.ColorNameDisabled:
		return color.RGBA{R: 160, G: 160, B: 160, A: 255} // Gray for disabled text
	case theme.ColorNameHover:
		return color.RGBA{R: 83, G: 101, B: 209, A: 255} // Dark green for hover effect
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

// Font overrides theme.Font to provide custom fonts (if needed)
func (c *CustomTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

// Icon overrides theme.Icon to provide custom icons
func (c *CustomTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

// Size overrides theme.Size to define custom dimensions
func (c *CustomTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameText:
		return 16 // Custom text size
	case theme.SizeNamePadding:
		return 20 // Custom padding size
	case theme.SizeNameInputBorder:
		return 100 // Custom button height
	case theme.SizeNameSubHeadingText:
		return 20 // Custom button width (optional)
	default:
		return theme.DefaultTheme().Size(name)
	}
}
