/*
Package application run the application
*/
package application

import (
	"fmt"

	"github.com/fatih/color"
)

type supportedColor string

// TextColor represents text color specification
type TextColor struct {
	Type  string           `json:"type"`
	Color []supportedColor `json:"color"`
}

// DefaultTextColor is the default color (no color)
var DefaultTextColor = TextColor{
	Type: "fixed", Color: []supportedColor{ColorDefault},
}

// UnmarshalJSON implements json.Unmarshaler for TextColor
func (s *TextColor) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as a string first
	str, err := UnmarshalAsString(data)
	if err == nil {
		parsed := supportedColor(str)
		if !isSupportedColor(parsed) {
			return fmt.Errorf("invalid color: %q", str)
		}
		*s = TextColor{
			Type:  "fixed",
			Color: []supportedColor{parsed},
		}
		return nil
	}

	// Try to unmarshal as a string array
	arr, err := UnmarshalAsStringArray(data)
	if err != nil {
		return err
	}
	*s = TextColor{
		Type:  "fixed",
		Color: make([]supportedColor, len(arr)),
	}
	for i, v := range arr {
		parsed := supportedColor(v)
		if !isSupportedColor(parsed) {
			return fmt.Errorf("invalid color at index %d: %q", i, v)
		}
		(*s).Color[i] = parsed
	}
	return nil
}

func isSupportedColor(c supportedColor) bool {
	_, ok := supportedColorMeta[c]
	return ok
}

const (
	// ColorDefault represents default color (no color)
	ColorDefault supportedColor = "default"

	// ColorRed represents red color
	ColorRed supportedColor = "red"
	// ColorGreen represents green color
	ColorGreen supportedColor = "green"
	// ColorBlue represents blue color
	ColorBlue supportedColor = "blue"
	// ColorYellow represents yellow color
	ColorYellow supportedColor = "yellow"
	// ColorMagenta represents magenta color
	ColorMagenta supportedColor = "magenta"
	// ColorCyan represents cyan color
	ColorCyan supportedColor = "cyan"
	// ColorWhite represents white color
	ColorWhite supportedColor = "white"
	// ColorBlack represents black color
	ColorBlack supportedColor = "black"

	// ColorHiRed represents high intensity red color
	ColorHiRed supportedColor = "hiRed"
	// ColorHiGreen represents high intensity green color
	ColorHiGreen supportedColor = "hiGreen"
	// ColorHiBlue represents high intensity blue color
	ColorHiBlue supportedColor = "hiBlue"
	// ColorHiYellow represents high intensity yellow color
	ColorHiYellow supportedColor = "hiYellow"
	// ColorHiMagenta represents high intensity magenta color
	ColorHiMagenta supportedColor = "hiMagenta"
	// ColorHiCyan represents high intensity cyan color
	ColorHiCyan supportedColor = "hiCyan"
	// ColorHiWhite represents high intensity white color
	ColorHiWhite supportedColor = "hiWhite"
	// ColorHiBlack represents high intensity black color
	ColorHiBlack supportedColor = "hiBlack"

	// ColorBgRed represents red background color
	ColorBgRed supportedColor = "bgRed"
	// ColorBgGreen represents green background color
	ColorBgGreen supportedColor = "bgGreen"
	// ColorBgBlue represents blue background color
	ColorBgBlue supportedColor = "bgBlue"
	// ColorBgYellow represents yellow background color
	ColorBgYellow supportedColor = "bgYellow"
	// ColorBgMagenta represents magenta background color
	ColorBgMagenta supportedColor = "bgMagenta"
	// ColorBgCyan represents cyan background color
	ColorBgCyan supportedColor = "bgCyan"
	// ColorBgWhite represents white background color
	ColorBgWhite supportedColor = "bgWhite"
	// ColorBgBlack represents black background color
	ColorBgBlack supportedColor = "bgBlack"

	// ColorHiBgRed represents high intensity red background color
	ColorHiBgRed supportedColor = "hiBgRed"
	// ColorHiBgGreen represents high intensity green background color
	ColorHiBgGreen supportedColor = "hiBgGreen"
	// ColorHiBgBlue represents high intensity blue background color
	ColorHiBgBlue supportedColor = "hiBgBlue"
	// ColorHiBgYellow represents high intensity yellow background color
	ColorHiBgYellow supportedColor = "hiBgYellow"
	// ColorHiBgMagenta represents high intensity magenta background color
	ColorHiBgMagenta supportedColor = "hiBgMagenta"
	// ColorHiBgCyan represents high intensity cyan background color
	ColorHiBgCyan supportedColor = "hiBgCyan"
	// ColorHiBgWhite represents high intensity white background color
	ColorHiBgWhite supportedColor = "hiBgWhite"
	// ColorHiBgBlack represents high intensity black background color
	ColorHiBgBlack supportedColor = "hiBgBlack"

	// ColorBold represents bold text
	ColorBold supportedColor = "bold"
	// ColorFaint represents faint text
	ColorFaint supportedColor = "faint"
	// ColorItalic represents italic text
	ColorItalic supportedColor = "italic"
	// ColorUnderline represents underlined text
	ColorUnderline supportedColor = "underline"
	// ColorBlinkSlow represents slow blinking text
	ColorBlinkSlow supportedColor = "blinkSlow"
	// ColorBlinkRapid represents rapid blinking text
	ColorBlinkRapid supportedColor = "blinkRapid"
	// ColorReverseVideo represents reverse video text
	ColorReverseVideo supportedColor = "reverseVideo"
	// ColorConcealed represents concealed text
	ColorConcealed supportedColor = "concealed"
	// ColorCrossedOut represents crossed out text
	ColorCrossedOut supportedColor = "crossedOut"
)

type colorMeta struct {
	color color.Attribute
}

var supportedColorMeta = map[supportedColor]colorMeta{
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
func GetColored(printValue string, textColor TextColor) any {
	s := textColor.Color
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
