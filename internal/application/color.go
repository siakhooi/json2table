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
	// ColorDefault represents default color (no color)
	ColorDefault SupportedColor = "default"
)

type colorMeta struct {
	color color.Attribute
}

// SupportedColorMeta is a list of supported colors
var SupportedColorMeta = map[SupportedColor]colorMeta{
	ColorRed:   {color: color.FgRed},
	ColorGreen: {color: color.FgGreen},
	ColorBlue:  {color: color.FgBlue},
}

// GetColored returns the printValue wrapped in color codes based on the color string
func GetColored(printValue string, s SupportedColor) any {
	if s == ColorDefault {
		return printValue
	}
	return color.New(SupportedColorMeta[s].color).SprintFunc()(printValue)
}
