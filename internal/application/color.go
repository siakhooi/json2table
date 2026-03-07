/*
Package application run the application
*/
package application

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
)

// SupportedColor represents text color in a column
type SupportedColor string

// SupportedColorArray represents an array of supported colors
type SupportedColorArray []SupportedColor

// DefaultTextColor is the default color (no color)
var DefaultTextColor = []SupportedColor{ColorDefault}

// UnmarshalJSON implements json.Unmarshaler for SupportedColorArray
func (s *SupportedColorArray) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as a string first
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		parsed := SupportedColor(str)
		if !isSupportedColor(parsed) {
			return fmt.Errorf("invalid color: %q", str)
		}
		*s = []SupportedColor{parsed}
		return nil
	}

	// Try to unmarshal as a string array
	var arr []string
	if err := json.Unmarshal(data, &arr); err != nil {
		return err
	}
	*s = make([]SupportedColor, len(arr))
	for i, v := range arr {
		parsed := SupportedColor(v)
		if !isSupportedColor(parsed) {
			return fmt.Errorf("invalid color at index %d: %q", i, v)
		}
		(*s)[i] = parsed
	}
	return nil
}

func isSupportedColor(c SupportedColor) bool {
	_, ok := supportedColorMeta[c]
	return ok
}

const (
	// ColorDefault represents default color (no color)
	ColorDefault SupportedColor = "default"

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

	// ColorHiRed represents high intensity red color
	ColorHiRed SupportedColor = "hiRed"
	// ColorHiGreen represents high intensity green color
	ColorHiGreen SupportedColor = "hiGreen"
	// ColorHiBlue represents high intensity blue color
	ColorHiBlue SupportedColor = "hiBlue"
	// ColorHiYellow represents high intensity yellow color
	ColorHiYellow SupportedColor = "hiYellow"
	// ColorHiMagenta represents high intensity magenta color
	ColorHiMagenta SupportedColor = "hiMagenta"
	// ColorHiCyan represents high intensity cyan color
	ColorHiCyan SupportedColor = "hiCyan"
	// ColorHiWhite represents high intensity white color
	ColorHiWhite SupportedColor = "hiWhite"
	// ColorHiBlack represents high intensity black color
	ColorHiBlack SupportedColor = "hiBlack"

	// ColorBgRed represents red background color
	ColorBgRed SupportedColor = "bgRed"
	// ColorBgGreen represents green background color
	ColorBgGreen SupportedColor = "bgGreen"
	// ColorBgBlue represents blue background color
	ColorBgBlue SupportedColor = "bgBlue"
	// ColorBgYellow represents yellow background color
	ColorBgYellow SupportedColor = "bgYellow"
	// ColorBgMagenta represents magenta background color
	ColorBgMagenta SupportedColor = "bgMagenta"
	// ColorBgCyan represents cyan background color
	ColorBgCyan SupportedColor = "bgCyan"
	// ColorBgWhite represents white background color
	ColorBgWhite SupportedColor = "bgWhite"
	// ColorBgBlack represents black background color
	ColorBgBlack SupportedColor = "bgBlack"

	// ColorHiBgRed represents high intensity red background color
	ColorHiBgRed SupportedColor = "hiBgRed"
	// ColorHiBgGreen represents high intensity green background color
	ColorHiBgGreen SupportedColor = "hiBgGreen"
	// ColorHiBgBlue represents high intensity blue background color
	ColorHiBgBlue SupportedColor = "hiBgBlue"
	// ColorHiBgYellow represents high intensity yellow background color
	ColorHiBgYellow SupportedColor = "hiBgYellow"
	// ColorHiBgMagenta represents high intensity magenta background color
	ColorHiBgMagenta SupportedColor = "hiBgMagenta"
	// ColorHiBgCyan represents high intensity cyan background color
	ColorHiBgCyan SupportedColor = "hiBgCyan"
	// ColorHiBgWhite represents high intensity white background color
	ColorHiBgWhite SupportedColor = "hiBgWhite"
	// ColorHiBgBlack represents high intensity black background color
	ColorHiBgBlack SupportedColor = "hiBgBlack"

	// ColorBold represents bold text
	ColorBold SupportedColor = "bold"
	// ColorFaint represents faint text
	ColorFaint SupportedColor = "faint"
	// ColorItalic represents italic text
	ColorItalic SupportedColor = "italic"
	// ColorUnderline represents underlined text
	ColorUnderline SupportedColor = "underline"
	// ColorBlinkSlow represents slow blinking text
	ColorBlinkSlow SupportedColor = "blinkSlow"
	// ColorBlinkRapid represents rapid blinking text
	ColorBlinkRapid SupportedColor = "blinkRapid"
	// ColorReverseVideo represents reverse video text
	ColorReverseVideo SupportedColor = "reverseVideo"
	// ColorConcealed represents concealed text
	ColorConcealed SupportedColor = "concealed"
	// ColorCrossedOut represents crossed out text
	ColorCrossedOut SupportedColor = "crossedOut"
)

type colorMeta struct {
	color color.Attribute
}

var supportedColorMeta = map[SupportedColor]colorMeta{
	ColorDefault:      {color: color.Reset},
	ColorRed:          {color: color.FgRed},
	ColorGreen:        {color: color.FgGreen},
	ColorBlue:         {color: color.FgBlue},
	ColorYellow:       {color: color.FgYellow},
	ColorMagenta:      {color: color.FgMagenta},
	ColorCyan:         {color: color.FgCyan},
	ColorWhite:        {color: color.FgWhite},
	ColorBlack:        {color: color.FgBlack},
	ColorHiRed:        {color: color.FgHiRed},
	ColorHiGreen:      {color: color.FgHiGreen},
	ColorHiBlue:       {color: color.FgHiBlue},
	ColorHiYellow:     {color: color.FgHiYellow},
	ColorHiMagenta:    {color: color.FgHiMagenta},
	ColorHiCyan:       {color: color.FgHiCyan},
	ColorHiWhite:      {color: color.FgHiWhite},
	ColorHiBlack:      {color: color.FgHiBlack},
	ColorBgRed:        {color: color.BgRed},
	ColorBgGreen:      {color: color.BgGreen},
	ColorBgBlue:       {color: color.BgBlue},
	ColorBgYellow:     {color: color.BgYellow},
	ColorBgMagenta:    {color: color.BgMagenta},
	ColorBgCyan:       {color: color.BgCyan},
	ColorBgWhite:      {color: color.BgWhite},
	ColorBgBlack:      {color: color.BgBlack},
	ColorHiBgRed:      {color: color.BgHiRed},
	ColorHiBgGreen:    {color: color.BgHiGreen},
	ColorHiBgBlue:     {color: color.BgHiBlue},
	ColorHiBgYellow:   {color: color.BgHiYellow},
	ColorHiBgMagenta:  {color: color.BgHiMagenta},
	ColorHiBgCyan:     {color: color.BgHiCyan},
	ColorHiBgWhite:    {color: color.BgHiWhite},
	ColorHiBgBlack:    {color: color.BgHiBlack},
	ColorBold:         {color: color.Bold},
	ColorFaint:        {color: color.Faint},
	ColorItalic:       {color: color.Italic},
	ColorUnderline:    {color: color.Underline},
	ColorBlinkSlow:    {color: color.BlinkSlow},
	ColorBlinkRapid:   {color: color.BlinkRapid},
	ColorReverseVideo: {color: color.ReverseVideo},
	ColorConcealed:    {color: color.Concealed},
	ColorCrossedOut:   {color: color.CrossedOut},
}

// GetColored returns the printValue wrapped in color codes based on the color string
func GetColored(printValue string, s []SupportedColor) any {
	colors := make([]color.Attribute, 0, len(s))
	for _, c := range s {
		meta, ok := supportedColorMeta[c]
		if !ok {
			continue
		}
		colors = append(colors, meta.color)
	}
	if len(colors) == 0 {
		return printValue
	}
	return color.New(colors...).SprintFunc()(printValue)
}
