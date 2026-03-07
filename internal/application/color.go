/*
Package application run the application
*/
package application

import "github.com/fatih/color"

// SupportedColor represents text color in a column
type SupportedColor string

const (
	// ColorRed represents red color
	ColorRed SupportedColor = "red"
	// ColorGreen represents green color
	ColorGreen SupportedColor = "green"
	// ColorBlue represents blue color
	ColorBlue SupportedColor = "blue"
	// ColorYellow represents yellow color
	ColorYellow SupportedColor = "yellow"
	// ColorMagenta represents magenta color
	ColorMagenta SupportedColor = "magenta"
	// ColorCyan represents cyan color
	ColorCyan SupportedColor = "cyan"
	// ColorWhite represents white color
	ColorWhite SupportedColor = "white"
	// ColorBlack represents black color
	ColorBlack SupportedColor = "black"
	// ColorDefault represents default color (no color)
	ColorDefault SupportedColor = "default"
)

type colorMeta struct {
	color color.Attribute
}

var supportedColorMeta = map[SupportedColor]colorMeta{
	ColorRed:     {color: color.FgRed},
	ColorGreen:   {color: color.FgGreen},
	ColorBlue:    {color: color.FgBlue},
	ColorYellow:  {color: color.FgYellow},
	ColorMagenta: {color: color.FgMagenta},
	ColorCyan:    {color: color.FgCyan},
	ColorWhite:   {color: color.FgWhite},
	ColorBlack:   {color: color.FgBlack},
	ColorDefault: {color: color.Reset},
}

// GetColored returns the printValue wrapped in color codes based on the color string
func GetColored(printValue string, s SupportedColor) any {
	return color.New(supportedColorMeta[s].color).SprintFunc()(printValue)
}
